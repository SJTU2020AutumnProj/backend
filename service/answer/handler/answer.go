package handler

import (
	mongoDB "boxin/service/answer/mongoDB"
	pb "boxin/service/answer/proto/answer"
	repo "boxin/service/answer/repository"
	"boxin/service/homework/proto/homework"

	// "boxin/service/answer/proto/answer"
	"context"
	"log"
	"time"

	// "golang.org/x/crypto/openpgp/errors"

	"github.com/micro/go-micro/v2/client"
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
		AnswerID: resp_answer.AnswerID,
		Content:  req.Content,
		Note:     req.Note,
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
		AnswerID: req.AnswerID,
		Content:  req.Content,
		Note:     req.Note,
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
			CommitTime: answer.CommitTime.Unix(),
			Content:    mongo_answer.Content,
			Note:       mongo_answer.Note,
		},
	}
	return nil
}

func (h *AnswerHandler) SearchAnswerByUserID(ctx context.Context, req *pb.UserID, resp *pb.SearchAnswerByUserIDResponse) error {
	answers, err := h.AnswerRepository.SearchAnswerByUserID(ctx, req.UserID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchAnswerByUserID error: ", err)
		return err
	}

	var ans []*pb.AnswerInfo
	for i := range answers {
		mongo_answer, err := h.AnswerMongo.SearchAnswer(ctx, answers[i].AnswerID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchAnswer error:", err)
			return err
		}
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
		mongo_answer, err := h.AnswerMongo.SearchAnswer(ctx, answers[i].AnswerID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchAnswer error:", err)
			return err
		}
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
	return nil
}

func (h *AnswerHandler) SearchAnswerByUserIDAndHomeworkID(ctx context.Context, req *pb.UserIDAndHomeworkID, resp *pb.SearchAnswerByUserIDAndHomeworkIDResponse) error {
	answers, err := h.AnswerRepository.SearchAnswerByHomeworkID(ctx, req.HomeworkID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchAnswerByUserIDAndHomeworkID error: ", err)
		return err
	}

	answers2, err := h.AnswerRepository.SearchAnswerByUserID(ctx, req.UserID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchAnswerByUserIDAndHomeworkID error: ", err)
		return err
	}

	var ans *pb.AnswerInfo
	var flag bool = false
	for i := range answers {
		for j := range answers2 {
			if answers[i].AnswerID == answers2[j].AnswerID {
				mongo_answer, err := h.AnswerMongo.SearchAnswer(ctx, answers[i].AnswerID)
				if nil != err {
					resp.Status = -1
					resp.Msg = "Error"
					log.Println("Handler SearchAnswerByUserAndHomeworkID error:", err)
					return err
				}
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

	if flag == false {
		var err error
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchAnswerByUserIDAndHomeworkID error: Cannot find the answer.", err)
		return err
	}
	*resp = pb.SearchAnswerByUserIDAndHomeworkIDResponse{
		Status: 0,
		Msg:    "Success",
		Answer: ans,
	}
	return nil
}

func (h *AnswerHandler) PostAnswerByStudent(ctx context.Context, req *pb.PostAnswerParam, resp *pb.PostAnswerResponse) error {
	//要比较commit_time和homework表中的end_time来填 user_homework表中的state 2为按时交 3为迟交
	homeworkService := homework.NewHomeworkService("go.micro.service.homework", client.DefaultClient)
	hw, err1 := homeworkService.SearchHomework(ctx, &homework.HomeworkID{HomeworkID: req.HomeworkID})

	if nil != err1 {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchTakeByCourse error: ", err1)
		return err1
	}

	ctime := req.CommitTime
	etime := hw.Homework.EndTime

	var state int32

	if ctime <= etime {
		state = 2
	} else {
		state = 3
	}

	ctime_time := time.Unix(req.CommitTime, 0)
	answer := repo.Answer{
		CommitTime: ctime_time,
	}
	var resp_answer repo.Answer
	var err error
	if resp_answer, err = h.AnswerRepository.PostAnswerByStudent(ctx, req.UserID, req.HomeworkID, state, answer); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler PostAnswerByStudent error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	resp.AnswerID = resp_answer.AnswerID

	mongo_answer := mongoDB.Answer{
		AnswerID: resp_answer.AnswerID,
		Content:  req.Content,
		Note:     req.Note,
	}
	if err = h.AnswerMongo.UpdateAnswer(ctx, mongo_answer); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler CreateAnswer error: ", err)
		return err
	}
	return nil
	// answer,err:=h.AnswerRepository.PostAnswer(ctx,req.UserID,req.HomeworkID,)
}

func (h *AnswerHandler) PostAnswerByTeacher(ctx context.Context, req *pb.PostAnswerParam, resp *pb.PostAnswerResponse) error {

	ctime_time := time.Unix(req.CommitTime, 0)
	answer := repo.Answer{
		CommitTime: ctime_time,
	}
	var resp_answer repo.Answer
	var err error
	if resp_answer, err = h.AnswerRepository.PostAnswerByTeacher(ctx, req.UserID, req.HomeworkID, answer); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler PostAnswerByTeacher error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	resp.AnswerID = resp_answer.AnswerID

	mongo_answer := mongoDB.Answer{
		AnswerID: resp_answer.AnswerID,
		Content:  req.Content,
		Note:     req.Note,
	}
	if err = h.AnswerMongo.UpdateAnswer(ctx, mongo_answer); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AnswerHandler CreateAnswer error: ", err)
		return err
	}
	return nil
	// answer,err:=h.AnswerRepository.PostAnswer(ctx,req.UserID,req.HomeworkID,)
}
