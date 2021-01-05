package handler

import (
	// "boxin/service/answer/proto/answer"
	mongoDB "boxin/service/homework/mongoDB"
	pb "boxin/service/homework/proto/homework"
	repo "boxin/service/homework/repository"

	// "boxin/service/courseclass/proto/courseclass"

	// "boxin/service/homework/proto/homework"
	"context"
	// "log"
	"time"

	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
)

// HomeworkHandler struct
type HomeworkHandler struct {
	HomeworkRepository repo.HomeworkRepository
	HomeworkMongo      mongoDB.HomeworkMongo
	// 由go-micro封装，用于发送消息的接口，老版本叫micro.Publisher
	// HomeworkAssignedPubEvent micro.Event
}

const (
	// HomeworkAssignedTopic topic of AssignHomework message
	HomeworkAssignedTopic = "assigned"
)

// AssignHomework assign homework
func (h *HomeworkHandler) AssignHomework(ctx context.Context, req *pb.AssignHomeworkParam, resp *pb.AssignHomeworkResponse) error {
	stime := time.Unix(req.StartTime, 0)
	etime := time.Unix(req.EndTime, 0)
	homework := repo.Homework{
		CourseID:  req.CourseID,
		UserID:    req.UserID,
		StartTime: stime,
		EndTime:   etime,
		Title:     req.Title,
		State:     req.State,
		AnswerID:  req.AnswerID,
	}
	var resp_homework repo.Homework
	var err error
	resp_homework, err = h.HomeworkRepository.AddHomework(ctx, homework)
	resp.Status = 0
	resp.Msg = "Success"
	resp.HomeworkID = resp_homework.HomeworkID

	mongo_homework := mongoDB.Homework{
		HomeworkID:  resp_homework.HomeworkID,
		Description: req.Description,
		Content:     req.Content,
		Note:        req.Note,
	}
	err = h.HomeworkMongo.AddHomework(ctx, mongo_homework)
	return err
}

func (h *HomeworkHandler) DeleteHomework(ctx context.Context, req *pb.HomeworkID, resp *pb.DeleteHomeworkResponse) error {
	err := h.HomeworkRepository.DeleteHomework(ctx, req.HomeworkID)
	resp.Status = 0
	resp.Msg = "Success"

	err = h.HomeworkMongo.DeleteHomework(ctx, req.HomeworkID)
		
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (h *HomeworkHandler) UpdateHomework(ctx context.Context, req *pb.HomeworkInfo, resp *pb.UpdateHomeworkResponse) error {
	stime := time.Unix(req.StartTime, 0)
	etime := time.Unix(req.EndTime, 0)
	homework := repo.Homework{
		HomeworkID: req.HomeworkID,
		CourseID:   req.CourseID,
		UserID:     req.UserID,
		StartTime:  stime,
		EndTime:    etime,
		Title:      req.Title,
		State:      req.State,
		AnswerID:   req.AnswerID,
	}

	err := h.HomeworkRepository.UpdateHomework(ctx, homework)
	resp.Status = 0
	resp.Msg = "Success"

	mongo_homework := mongoDB.Homework{
		HomeworkID:  req.HomeworkID,
		Description: req.Description,
		Content:     req.Content,
		Note:        req.Note,
	}

	err = h.HomeworkMongo.UpdateHomework(ctx, mongo_homework)
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (h *HomeworkHandler) SearchHomework(ctx context.Context, req *pb.HomeworkID, resp *pb.SearchHomeworkResponse) error {

	homework, err := h.HomeworkRepository.SearchHomework(ctx, req.HomeworkID)

	mongo_homework, err := h.HomeworkMongo.SearchHomework(ctx, req.HomeworkID)

	*resp = pb.SearchHomeworkResponse{
		Status: 0,
		Homework: &pb.HomeworkInfo{
			HomeworkID:  homework.HomeworkID,
			CourseID:    homework.CourseID,
			UserID:      homework.UserID,
			StartTime:   homework.StartTime.Unix(),
			EndTime:     homework.EndTime.Unix(),
			Title:       homework.Title,
			State:       homework.State,
			AnswerID:    homework.AnswerID,
			Description: mongo_homework.Description,
			Content:     mongo_homework.Content,
			Note:        mongo_homework.Note,
		},
	}
	return err
}

func (h *HomeworkHandler) SearchHomeworkByUserID(ctx context.Context, req *pb.UserID, resp *pb.SearchHomeworkByUserIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByUserID(ctx, req.UserID)

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		var homework_json mongoDB.Homework
		homework_json,err = h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
		
		ans = append(ans, &pb.HomeworkInfo{
			HomeworkID:  homeworks[i].HomeworkID,
			CourseID:    homeworks[i].CourseID,
			UserID:      homeworks[i].UserID,
			StartTime:   homeworks[i].StartTime.Unix(),
			EndTime:     homeworks[i].EndTime.Unix(),
			Title:       homeworks[i].Title,
			State:       homeworks[i].State,
			AnswerID:    homeworks[i].AnswerID,
			Description: homework_json.Description,
			Content:     homework_json.Content,
			Note:        homework_json.Note,
		})
	}

	*resp = pb.SearchHomeworkByUserIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return err
}

func (h *HomeworkHandler) SearchHomeworkByCourseID(ctx context.Context, req *pb.CourseID, resp *pb.SearchHomeworkByCourseIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByCourseID(ctx, req.CourseID)

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		var homework_json mongoDB.Homework
		homework_json, err = h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
	
		ans = append(ans, &pb.HomeworkInfo{
			HomeworkID:  homeworks[i].HomeworkID,
			CourseID:    homeworks[i].CourseID,
			UserID:      homeworks[i].UserID,
			StartTime:   homeworks[i].StartTime.Unix(),
			EndTime:     homeworks[i].EndTime.Unix(),
			Title:       homeworks[i].Title,
			State:       homeworks[i].State,
			Description: homework_json.Description,
			Content:     homework_json.Content,
		})
	}

	*resp = pb.SearchHomeworkByCourseIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return err
}

//老师上传作业答案
func (h *HomeworkHandler) PostHomeworkAnswer(ctx context.Context, req *pb.PostParam, resp *pb.PostHomeworkAnswerResponse) error {

	err1 := h.HomeworkRepository.PostHomeworkAnswer(ctx, req.HomeworkID, 1)

	return err1
}

//老师公布作业答案
func (h *HomeworkHandler) ReleaseHomeworkAnswer(ctx context.Context, req *pb.ReleaseParam, resp *pb.ReleaseHomeworkAnswerResponse) error {
	err := h.HomeworkRepository.ReleaseHomeworkAnswer(ctx, req.HomeworkID)
	*resp = pb.ReleaseHomeworkAnswerResponse{
		Status: 0,
		Msg:    "Success",
	}
	return err
}

func (h *HomeworkHandler) StudentSearchHomework(ctx context.Context, req *pb.StudentSearchHomeworkParam, resp *pb.StudentSearchHomeworkResponse) error {
	homework, err := h.HomeworkRepository.SearchHomework(ctx, req.HomeworkID)
	
	mongo_homework, err := h.HomeworkMongo.SearchHomework(ctx, req.HomeworkID)

	*resp = pb.StudentSearchHomeworkResponse{
		Status: 0,
		Homework: &pb.HomeworkInfo{
			HomeworkID:  homework.HomeworkID,
			CourseID:    homework.CourseID,
			UserID:      homework.UserID,
			StartTime:   homework.StartTime.Unix(),
			EndTime:     homework.EndTime.Unix(),
			Title:       homework.Title,
			State:       homework.State,
			AnswerID:    homework.AnswerID,
			Description: mongo_homework.Description,
			Content:     mongo_homework.Content,
			Note:        mongo_homework.Note,
		},
	}

	err = h.HomeworkRepository.UpdateUserHomeworkState(ctx, req.UserID, req.HomeworkID, 1)
	
	return err
}

//老师公布批改情况,即修改user_homework表中的state为4
func (h *HomeworkHandler) ReleaseCheck(ctx context.Context, req *pb.ReleaseCheckParam, resp *pb.ReleaseCheckResponse) error {
	err := h.HomeworkRepository.UpdateUserHomeworkState(ctx, req.UserID, req.HomeworkID, 4)
	*resp = pb.ReleaseCheckResponse{
		Status: 0,
		Msg:    "Success",
	}
	return err
}
