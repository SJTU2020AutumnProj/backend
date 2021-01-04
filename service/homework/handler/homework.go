package handler

import (
	"boxin/service/answer/proto/answer"
	mongoDB "boxin/service/homework/mongoDB"
	pb "boxin/service/homework/proto/homework"
	repo "boxin/service/homework/repository"
	

	// "boxin/service/homework/proto/homework"
	"context"
	"log"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// HomeworkHandler struct
type HomeworkHandler struct {
	HomeworkRepository repo.HomeworkRepository
	HomeworkMongo      mongoDB.HomeworkMongo
	// 由go-micro封装，用于发送消息的接口，老版本叫micro.Publisher
	HomeworkAssignedPubEvent micro.Event
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
		AnswerID: req.AnswerID,
	}
	var resp_homework repo.Homework
	var err error
	if resp_homework, err = h.HomeworkRepository.AddHomework(ctx, homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler AssignHomework error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	resp.HomeworkID = resp_homework.HomeworkID

	mongo_homework := mongoDB.Homework{
		HomeworkID:  resp_homework.HomeworkID,
		Description: req.Description,
		Content:     req.Content,
		Note: 		req.Note,
	}
	if err = h.HomeworkMongo.AddHomework(ctx, mongo_homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler AssignHomework error: ", err)
		return err
	}
	// assignedHomework := &pb.AssignedHomework{
	// 	HomeworkID:  resp_homework.HomeworkID,
	// 	CourseID:    resp_homework.CourseID,
	// 	UserID:      resp_homework.UserID,
	// 	StartTime:   resp_homework.StartTime.Unix(),
	// 	EndTime:     resp_homework.EndTime.Unix(),
	// 	Title:       resp_homework.Title,
	// 	State:       resp_homework.State,
	// 	Description: req.Description,
	// 	Content: req.Content,  
	// 	Note:req.Note,
	// }
	// if err = h.HomeworkAssignedPubEvent.Publish(ctx, assignedHomework); err != nil {
	// 	log.Println("HomeworkHandler AssignHomework error when sending message: ", err)
	// }
	return nil
}

func (h *HomeworkHandler) DeleteHomework(ctx context.Context, req *pb.HomeworkID, resp *pb.DeleteHomeworkResponse) error {
	if err := h.HomeworkRepository.DeleteHomework(ctx, req.HomeworkID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler DeleteHomework error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"

	if err := h.HomeworkMongo.DeleteHomework(ctx, req.HomeworkID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler DeleteHomework error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
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
		AnswerID: req.AnswerID,
	}

	if err := h.HomeworkRepository.UpdateHomework(ctx, homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler UpdateHomework error:", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"

	mongo_homework := mongoDB.Homework{
		HomeworkID:  req.HomeworkID,
		Description: req.Description,
		Content:     req.Content,
		Note:	 req.Note,
	}

	if err := h.HomeworkMongo.UpdateHomework(ctx, mongo_homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler DeleteHomework error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (h *HomeworkHandler) SearchHomework(ctx context.Context, req *pb.HomeworkID, resp *pb.SearchHomeworkResponse) error {
	// stime := time.Unix(req.StartTime, 0)
	// etime := time.Unix(req.EndTime, 0)
	homework, err := h.HomeworkRepository.SearchHomework(ctx, req.HomeworkID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomework error:", err)
		return err
	}

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
			AnswerID: homework.AnswerID,
			Description: mongo_homework.Description,
			Content:     mongo_homework.Content,
			Note: 		mongo_homework.Note,
		},
	}
	return nil
}

func (h *HomeworkHandler) SearchHomeworkByUserID(ctx context.Context, req *pb.UserID, resp *pb.SearchHomeworkByUserIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByUserID(ctx, req.UserID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomeByUserID error: ", err)
		return err
	}

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		homework_json, err := h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchHomework error:", err)
			return err
		}
		ans = append(ans, &pb.HomeworkInfo{
			HomeworkID:  homeworks[i].HomeworkID,
			CourseID:    homeworks[i].CourseID,
			UserID:      homeworks[i].UserID,
			StartTime:   homeworks[i].StartTime.Unix(),
			EndTime:     homeworks[i].EndTime.Unix(),
			Title:       homeworks[i].Title,
			State:       homeworks[i].State,
			AnswerID: homeworks[i].AnswerID,
			Description: homework_json.Description,
			Content:     homework_json.Content,
			Note: homework_json.Note,
		})
	}

	*resp = pb.SearchHomeworkByUserIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return nil
}

//发布作业答案
// func (h*HomeworkHandler) PostHomeworkAnswer(ctx context.Context, req *pb.HomeworkAnswer,resp *pb.PostHomeworkAnswerResponse) error{
// 	homework,err:=h.HomeworkRepository.SearchHomework(ctx,req.HomeworkID)
// 	if nil != err {
// 		resp.Status = -1
// 		resp.Msg = "Error"
// 		log.Println("Handler PostHomeworkAnswer error: ", err)
// 		return err
// 	}
// 	homework.AnswerID = req.AnswerID

// 	if err := h.HomeworkRepository.UpdateHomework(ctx, homework); nil != err {
// 		resp.Status = -1
// 		resp.Msg = "Error"
// 		log.Println("HomeworkHandler UpdateHomework error:", err)
// 		return err
// 	}
// 	resp.Status = 0
// 	resp.Msg = "Success"
// 	return nil
// }


func (h *HomeworkHandler) SearchHomeworkByCourseID(ctx context.Context, req *pb.CourseID, resp *pb.SearchHomeworkByCourseIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByCourseID(ctx, req.CourseID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomeByCourseID error: ", err)
		return err
	}

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		homework_json, err := h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchHomework error:", err)
			return err
		}
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
	return nil
}

const (
	ServiceName = "go.micro.client.answer"
	EtcdAddr    = "localhost:2379"
)
//老师上传作业答案
func (h*HomeworkHandler) PostHomeworkAnswer(ctx context.Context, req *pb.PostParam,resp *pb.PostHomeworkAnswerResponse) error{
	server := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)
	server.Init()
	answerService := answer.NewAnswerService("go.micro.service.answer", server.Client())

	aw,err := answerService.PostAnswerByTeacher(context.Background(),&answer.PostAnswerParam{HomeworkID: req.HomeworkID , UserID:req.UserID,CommitTime:req.CommitTime,Content:req.Content,Note:req.Note})

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler PostHomeworkAnswer error: ", err)
		return err
	}

	resp.AnswerID = aw.AnswerID

	err1:= h.HomeworkRepository.PostHomeworkAnswer(ctx,req.HomeworkID,aw.AnswerID)
	if nil != err1 {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler PostHomeworkAnswer error: ", err1)
		return err1
	}
	return nil
}

//老师公布作业答案
func (h*HomeworkHandler) ReleaseHomeworkAnswer(ctx context.Context, req *pb.RealeaseParam, resp *pb.ReleaseHomeworkAnswerResponse)error{
	if err := h.HomeworkRepository.ReleaseHomeworkAnswer(ctx, req.HomeworkID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler ReleaseHomeworkAnswer error:", err)
		return err
	}
	*resp = pb.ReleaseHomeworkAnswerResponse{
		Status:    0,
		Msg:       "Success",
	}
	return nil
}
