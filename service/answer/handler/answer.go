package handler

import (
	mongoDB "boxin/service/answer/mongoDB"
	pb "boxin/service/answer/proto/answer"
	repo "boxin/service/answer/repository"

	// "boxin/service/homework/proto/homework"

	// "boxin/service/answer/proto/answer"
	"context"
	"log"
	"time"
	// "golang.org/x/crypto/openpgp/errors"
	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
)

type AnswerHandler struct {
	AnswerRepository repo.AnswerRepository
	AnswerMongo      mongoDB.AnswerMongo
}

func (h *AnswerHandler) CreateAnswer(ctx context.Context, req *pb.CreateAnswerParam, resp *pb.CreateAnswerResponse) error {
	ctime := time.Unix(req.CommitTime, 0)
	answer := repo.Answer{
		CommitTime: ctime,
	}
	var resp_answer repo.Answer
	var err error
	resp_answer, err = h.AnswerRepository.AddAnswer(ctx, answer)
	resp.Status = 0
	resp.Msg = "Success"
	resp.AnswerID = resp_answer.AnswerID

	mongo_answer := mongoDB.Answer{
		AnswerID: resp_answer.AnswerID,
		Content:  req.Content,
		Note:     req.Note,
	}
	err = h.AnswerMongo.AddAnswer(ctx, mongo_answer)
	return err
}

func (h *AnswerHandler) DeleteAnswer(ctx context.Context, req *pb.AnswerID, resp *pb.DeleteAnswerResponse) error {
	err := h.AnswerRepository.DeleteAnswer(ctx, req.AnswerID)
	resp.Status = 0
	resp.Msg = "Success"

	err = h.AnswerMongo.DeleteAnswer(ctx, req.AnswerID)
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (h *AnswerHandler) UpdateAnswer(ctx context.Context, req *pb.AnswerInfo, resp *pb.UpdateAnswerResponse) error {
	ctime := time.Unix(req.CommitTime, 0)
	answer := repo.Answer{
		AnswerID:   req.AnswerID,
		CommitTime: ctime,
	}

	err := h.AnswerRepository.UpdateAnswer(ctx, answer)
	resp.Status = 0
	resp.Msg = "Success"

	mongo_answer := mongoDB.Answer{
		AnswerID: req.AnswerID,
		Content:  req.Content,
		Note:     req.Note,
	}

	err = h.AnswerMongo.UpdateAnswer(ctx, mongo_answer)
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (h *AnswerHandler) SearchAnswer(ctx context.Context, req *pb.AnswerID, resp *pb.SearchAnswerResponse) error {
	// stime := time.Unix(req.StartTime, 0)
	// etime := time.Unix(req.EndTime, 0)
	answer, err := h.AnswerRepository.SearchAnswer(ctx, req.AnswerID)
	mongo_answer, err := h.AnswerMongo.SearchAnswer(ctx, req.AnswerID)

	*resp = pb.SearchAnswerResponse{
		Status: 0,
		Answer: &pb.AnswerInfo{
			AnswerID:   answer.AnswerID,
			CommitTime: answer.CommitTime.Unix(),
			Content:    mongo_answer.Content,
			Note:       mongo_answer.Note,
		},
	}
	return err
}

func (h *AnswerHandler) SearchAnswerByUserID(ctx context.Context, req *pb.UserID, resp *pb.SearchAnswerByUserIDResponse) error {
	answers, err := h.AnswerRepository.SearchAnswerByUserID(ctx, req.UserID)

	var ans []*pb.AnswerInfo
	for i := range answers {
		mongo_answer, err := h.AnswerMongo.SearchAnswer(ctx, answers[i].AnswerID)
		log.Println(err)
		ans = append(ans, &pb.AnswerInfo{
			AnswerID:   answers[i].AnswerID,
			CommitTime: answers[i].CommitTime.Unix(),
			Content:    mongo_answer.Content,
			Note:       mongo_answer.Note,
		})
	}

	*resp = pb.SearchAnswerByUserIDResponse{
		Status:  0,
		Msg:     "Success",
		Answers: ans,
	}
	return err
}

func (h *AnswerHandler) SearchAnswerByHomeworkID(ctx context.Context, req *pb.HomeworkID, resp *pb.SearchAnswerByHomeworkIDResponse) error {
	answers, err := h.AnswerRepository.SearchAnswerByHomeworkID(ctx, req.HomeworkID)

	var ans []*pb.AnswerInfo
	for i := range answers {
		mongo_answer, err := h.AnswerMongo.SearchAnswer(ctx, answers[i].AnswerID)
		log.Println(err)
		ans = append(ans, &pb.AnswerInfo{
			AnswerID:   answers[i].AnswerID,
			CommitTime: answers[i].CommitTime.Unix(),
			Content:    mongo_answer.Content,
			Note:       mongo_answer.Note,
		})
	}

	*resp = pb.SearchAnswerByHomeworkIDResponse{
		Status:  0,
		Msg:     "Success",
		Answers: ans,
	}
	return err
}

func (h *AnswerHandler) SearchAnswerByUserIDAndHomeworkID(ctx context.Context, req *pb.UserIDAndHomeworkID, resp *pb.SearchAnswerByUserIDAndHomeworkIDResponse) error {
	answers, err := h.AnswerRepository.SearchAnswerByHomeworkID(ctx, req.HomeworkID)
	answers2, err := h.AnswerRepository.SearchAnswerByUserID(ctx, req.UserID)
	var ans *pb.AnswerInfo
	var flag bool = false
	for i := range answers {
		for j := range answers2 {
			if answers[i].AnswerID == answers2[j].AnswerID {
				mongo_answer, err := h.AnswerMongo.SearchAnswer(ctx, answers[i].AnswerID)
				log.Println(err, flag)
				ans = &pb.AnswerInfo{
					AnswerID:   answers[i].AnswerID,
					CommitTime: answers[i].CommitTime.Unix(),
					Content:    mongo_answer.Content,
					Note:       mongo_answer.Note,
				}
				flag = true
				break
			}
		}
	}

	*resp = pb.SearchAnswerByUserIDAndHomeworkIDResponse{
		Status: 0,
		Msg:    "Success",
		Answer: ans,
	}
	return err
}

const (
	ServiceName = "go.micro.client.homework"
	EtcdAddr    = "localhost:2379"
)

func (h *AnswerHandler) PostAnswerByStudent(ctx context.Context, req *pb.PostAnswerParam, resp *pb.PostAnswerResponse) error {

	// ctime := req.CommitTime
	// etime := time.Now().Unix()

	var state int32

	state = 2

	ctime_time := time.Unix(req.CommitTime, 0)
	answer := repo.Answer{
		CommitTime: ctime_time,
	}
	var resp_answer repo.Answer
	var err error
	resp_answer, err = h.AnswerRepository.PostAnswerByStudent(ctx, req.UserID, req.HomeworkID, state, answer)
	resp.Status = 0
	resp.Msg = "Success"
	resp.AnswerID = resp_answer.AnswerID

	mongo_answer := mongoDB.Answer{
		AnswerID: resp_answer.AnswerID,
		Content:  req.Content,
		Note:     req.Note,
	}
	err = h.AnswerMongo.AddAnswer(ctx, mongo_answer)
	return err
	// answer,err:=h.AnswerRepository.PostAnswer(ctx,req.UserID,req.HomeworkID,)
}

func (h *AnswerHandler) PostAnswerByTeacher(ctx context.Context, req *pb.PostAnswerParam, resp *pb.PostAnswerResponse) error {

	ctime_time := time.Unix(req.CommitTime, 0)
	answer := repo.Answer{
		CommitTime: ctime_time,
	}
	var resp_answer repo.Answer
	var err error
	resp_answer, err = h.AnswerRepository.PostAnswerByTeacher(ctx, req.UserID, req.HomeworkID, answer)
	resp.Status = 0
	resp.Msg = "Success"
	resp.AnswerID = resp_answer.AnswerID

	mongo_answer := mongoDB.Answer{
		AnswerID: resp_answer.AnswerID,
		Content:  req.Content,
		Note:     req.Note,
	}
	err = h.AnswerMongo.AddAnswer(ctx, mongo_answer)
	return err
	// answer,err:=h.AnswerRepository.PostAnswer(ctx,req.UserID,req.HomeworkID,)
}
