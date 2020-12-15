package handler

import (
	mongoDB "boxin/service/homework/mongoDB"
	pb "boxin/service/homework/proto/homework"
	repo "boxin/service/homework/repository"

	// "boxin/service/homework/proto/homework"
	"context"
	"log"
	"time"
	// "golang.org/x/crypto/openpgp/errors"
	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
)

type HomeworkHandler struct {
	HomeworkRepository repo.HomeworkRepository
	HomeworkMongo      mongoDB.HomeworkMongo
}

func (h *HomeworkHandler) AssignHomework(ctx context.Context, req *pb.AssignHomeworkParam, resp *pb.AssignHomeworkResponse) error {
	stime := time.Unix(req.StartTime, 0)
	etime := time.Unix(req.EndTime, 0)
	homework := repo.Homework{
		CourseID:  req.CourseID,
		TeacherID: req.TeacherID,
		StartTime: stime,
		EndTime:   etime,
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
		HomeworkID:   resp_homework.HomeworkID,
		HomeworkJson: req.HomeworkJson,
	}
	if err = h.HomeworkMongo.AddHomework(ctx, mongo_homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler AssignHomework error: ", err)
		return err
	}
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
		TeacherID:  req.TeacherID,
		StartTime:  stime,
		EndTime:    etime,
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
		HomeworkID:   req.HomeworkID,
		HomeworkJson: req.HomeworkJson,
	}

	if err := h.HomeworkMongo.UpdateHomework(ctx,mongo_homework);nil!=err {
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

	mongo_homework,err:=h.HomeworkMongo.SearchHomework(ctx,req.HomeworkID)

	*resp = pb.SearchHomeworkResponse{
		Status: 0,
		Homework: &pb.HomeworkInfo{
			HomeworkID: homework.HomeworkID,
			CourseID:   homework.CourseID,
			TeacherID:  homework.TeacherID,
			StartTime:  homework.StartTime.Unix(),
			EndTime:    homework.EndTime.Unix(),
			HomeworkJson: mongo_homework.HomeworkJson,
		},
	}
	return nil
}

func (h *HomeworkHandler) SearchHomeworkByTeacherID(ctx context.Context, req *pb.TeacherID, resp *pb.SearchHomeworkByTeacherIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByTeacherID(ctx, req.TeacherID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomeByTeacherID error: ", err)
		return err
	}

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		homework_json,err := h.HomeworkMongo.SearchHomework(ctx,homeworks[i].HomeworkID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchHomework error:", err)
			return err
		}
		ans = append(ans, &pb.HomeworkInfo{
			HomeworkID: homeworks[i].HomeworkID,
			CourseID:   homeworks[i].CourseID,
			TeacherID:  homeworks[i].TeacherID,
			StartTime:  homeworks[i].StartTime.Unix(),
			EndTime:    homeworks[i].EndTime.Unix(),
			HomeworkJson: homework_json.HomeworkJson,
		})
	}

	*resp = pb.SearchHomeworkByTeacherIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return nil
}
