package handler

import (
	mongoDB "boxin/service/answer/mongoDB"
	pb "boxin/service/answer/proto/answer"
	repo "boxin/service/answer/repository"

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
		HomeworkID: req.HomeworkID,
		StudentID:  req.StudentID,
		Status:     req.Status,
		CommitTime: ctime,
	}
	var resp_answer repo.Answer
	var err error
	if resp_answer, err = h.AnswerRepository.AddAnswer(ctx, answer); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler CreateAnswer error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	resp.AnswerID = resp_answer.AnswerID

	mongo_answer := mongoDB.Answer{
		AnswerID:   resp_answer.AnswerID,
		AnswerJson: req.AnswerJson,
	}
	if err = h.AnswerMongo.AddAnswer(ctx, mongo_answer); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler CreateAnswer error: ", err)
		return err
	}
	return nil
}

func (h *AnswerHandler) DeleteAnswer(ctx context.Context, req *pb.AnswerID, resp *pb.DeleteAnswerResponse) error {
	if err := h.AnswerRepository.DeleteAnswer(ctx, req.AnswerID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler DeleteAnswer error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"

	if err := h.AnswerMongo.DeleteAnswer(ctx, req.AnswerID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler DeleteAnswer error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (h *AnswerHandler) UpdateAnswer(ctx context.Context, req *pb.AnswerInfo, resp *pb.UpdateAnswerResponse) error {
	ctime := time.Unix(req.CommitTime, 0)
	answer := repo.Answer{
		AnswerID:   req.AnswerID,
		HomeworkID: req.HomeworkID,
		StudentID:  req.StudentID,
		Status:     req.Status,
		CommitTime: ctime,
	}

	if err := h.AnswerRepository.UpdateAnswer(ctx, answer); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler UpdateAnswer error:", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"

	mongo_answer := mongoDB.Answer{
		AnswerID:   req.AnswerID,
		AnswerJson: req.AnswerJson,
	}

	if err := h.AnswerMongo.UpdateAnswer(ctx, mongo_answer); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler DeleteAnswer error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (h *AnswerHandler) SearchAnswer(ctx context.Context, req *pb.AnswerID, resp *pb.SearchAnswerResponse) error {
	// stime := time.Unix(req.StartTime, 0)
	// etime := time.Unix(req.EndTime, 0)
	answer, err := h.AnswerRepository.SearchAnswer(ctx, req.AnswerID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchAnswer error:", err)
		return err
	}

	mongo_answer, err := h.AnswerMongo.SearchAnswer(ctx, req.AnswerID)

	*resp = pb.SearchAnswerResponse{
		Status: 0,
		Answer: &pb.AnswerInfo{
			AnswerID:   answer.AnswerID,
			HomeworkID: answer.HomeworkID,
			StudentID:  answer.StudentID,
			Status:     answer.Status,
			CommitTime: answer.CommitTime.Unix(),
			AnswerJson: mongo_answer.AnswerJson,
		},
	}
	return nil
}

func (h *AnswerHandler) SearchAnswerByStudentID(ctx context.Context, req *pb.StudentID, resp *pb.SearchAnswerByStudentIDResponse) error {
	answers, err := h.AnswerRepository.SearchAnswerByStudentID(ctx, req.StudentID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchAnswerByStudentID error: ", err)
		return err
	}

	var ans []*pb.AnswerInfo
	for i := range answers {
		answer_json, err := h.AnswerMongo.SearchAnswer(ctx, answers[i].AnswerID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchAnswer error:", err)
			return err
		}
		ans = append(ans, &pb.AnswerInfo{
			AnswerID:   answers[i].AnswerID,
			HomeworkID: answers[i].HomeworkID,
			StudentID:  answers[i].StudentID,
			Status:     answers[i].Status,
			CommitTime: answers[i].CommitTime.Unix(),
			AnswerJson: answer_json.AnswerJson,
		})
	}

	*resp = pb.SearchAnswerByStudentIDResponse{
		Status:  0,
		Msg:     "Success",
		Answers: ans,
	}
	return nil
}

func (h *AnswerHandler) SearchAnswerByHomeworkID(ctx context.Context, req *pb.HomeworkID, resp *pb.SearchAnswerByHomeworkIDResponse) error {
	answers, err := h.AnswerRepository.SearchAnswerByHomeworkID(ctx, req.HomeworkID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchAnswerByHomeworkID error: ", err)
		return err
	}

	var ans []*pb.AnswerInfo
	for i := range answers {
		answer_json, err := h.AnswerMongo.SearchAnswer(ctx, answers[i].AnswerID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchAnswer error:", err)
			return err
		}
		ans = append(ans, &pb.AnswerInfo{
			AnswerID:   answers[i].AnswerID,
			HomeworkID: answers[i].HomeworkID,
			StudentID:  answers[i].StudentID,
			Status:     answers[i].Status,
			CommitTime: answers[i].CommitTime.Unix(),
			AnswerJson: answer_json.AnswerJson,
		})
	}

	*resp = pb.SearchAnswerByHomeworkIDResponse{
		Status:  0,
		Msg:     "Success",
		Answers: ans,
	}
	return nil
}
