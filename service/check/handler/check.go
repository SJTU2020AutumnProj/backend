package handler

import (
	mongoDB "boxin/service/check/mongoDB"
	pb "boxin/service/check/proto/check"
	repo "boxin/service/check/repository"

	// "boxin/service/check/proto/check"
	"context"
	"log"
	"time"
	// "golang.org/x/crypto/openpgp/errors"
	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
)

type CheckHandler struct {
	CheckRepository repo.CheckRepository
	CheckMongo      mongoDB.CheckMongo
}

func (h *CheckHandler) CreateCheck(ctx context.Context, req *pb.CreateCheckParam, resp *pb.CreateCheckResponse) error {
	ctime := time.Unix(req.CheckTime, 0)
	check := repo.Check{
		HomeworkID: req.HomeworkID,
		TeacherID:  req.TeacherID,
		CheckTime: ctime,
	}
	var resp_check repo.Check
	var err error
	if resp_check, err = h.CheckRepository.AddCheck(ctx, check); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CheckHandler CreateCheck error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	resp.CheckID = resp_check.CheckID

	mongo_check := mongoDB.Check{
		CheckID:   resp_check.CheckID,
		CheckJson: req.CheckJson,
	}
	if err = h.CheckMongo.AddCheck(ctx, mongo_check); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CheckHandler CreateCheck error: ", err)
		return err
	}
	return nil
}

func (h *CheckHandler) DeleteCheck(ctx context.Context, req *pb.CheckID, resp *pb.DeleteCheckResponse) error {
	if err := h.CheckRepository.DeleteCheck(ctx, req.CheckID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CheckHandler DeleteCheck error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"

	if err := h.CheckMongo.DeleteCheck(ctx, req.CheckID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CheckHandler DeleteCheck error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (h *CheckHandler) UpdateCheck(ctx context.Context, req *pb.CheckInfo, resp *pb.UpdateCheckResponse) error {
	ctime := time.Unix(req.CheckTime, 0)
	check := repo.Check{
		CheckID:   req.CheckID,
		HomeworkID: req.HomeworkID,
		TeacherID:  req.TeacherID,
		CheckTime: ctime,
	}

	if err := h.CheckRepository.UpdateCheck(ctx, check); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CheckHandler UpdateCheck error:", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"

	mongo_check := mongoDB.Check{
		CheckID:   req.CheckID,
		CheckJson: req.CheckJson,
	}

	if err := h.CheckMongo.UpdateCheck(ctx, mongo_check); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CheckHandler DeleteCheck error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (h *CheckHandler) SearchCheck(ctx context.Context, req *pb.CheckID, resp *pb.SearchCheckResponse) error {
	// stime := time.Unix(req.StartTime, 0)
	// etime := time.Unix(req.EndTime, 0)
	check, err := h.CheckRepository.SearchCheck(ctx, req.CheckID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchCheck error:", err)
		return err
	}

	mongo_check, err := h.CheckMongo.SearchCheck(ctx, req.CheckID)

	*resp = pb.SearchCheckResponse{
		Status: 0,
		Check: &pb.CheckInfo{
			CheckID:   check.CheckID,
			HomeworkID: check.HomeworkID,
			TeacherID:  check.TeacherID,
			CheckTime: check.CheckTime.Unix(),
			CheckJson: mongo_check.CheckJson,
		},
	}
	return nil
}

func (h *CheckHandler) SearchCheckByTeacherID(ctx context.Context, req *pb.TeacherID, resp *pb.SearchCheckByTeacherIDResponse) error {
	checks, err := h.CheckRepository.SearchCheckByTeacherID(ctx, req.TeacherID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchCheckByTeacherID error: ", err)
		return err
	}

	var ans []*pb.CheckInfo
	for i := range checks {
		check_json, err := h.CheckMongo.SearchCheck(ctx, checks[i].CheckID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchCheck error:", err)
			return err
		}
		ans = append(ans, &pb.CheckInfo{
			CheckID:   checks[i].CheckID,
			HomeworkID: checks[i].HomeworkID,
			TeacherID:  checks[i].TeacherID,
			CheckTime: checks[i].CheckTime.Unix(),
			CheckJson: check_json.CheckJson,
		})
	}

	*resp = pb.SearchCheckByTeacherIDResponse{
		Status:  0,
		Msg:     "Success",
		Checks: ans,
	}
	return nil
}

func (h *CheckHandler) SearchCheckByHomeworkID(ctx context.Context, req *pb.HomeworkID, resp *pb.SearchCheckByHomeworkIDResponse) error {
	checks, err := h.CheckRepository.SearchCheckByHomeworkID(ctx, req.HomeworkID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchCheckByHomeworkID error: ", err)
		return err
	}

	var ans []*pb.CheckInfo
	for i := range checks {
		check_json, err := h.CheckMongo.SearchCheck(ctx, checks[i].CheckID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchCheck error:", err)
			return err
		}
		ans = append(ans, &pb.CheckInfo{
			CheckID:   checks[i].CheckID,
			HomeworkID: checks[i].HomeworkID,
			TeacherID:  checks[i].TeacherID,
			CheckTime: checks[i].CheckTime.Unix(),
			CheckJson: check_json.CheckJson,
		})
	}

	*resp = pb.SearchCheckByHomeworkIDResponse{
		Status:  0,
		Msg:     "Success",
		Checks: ans,
	}
	return nil
}
