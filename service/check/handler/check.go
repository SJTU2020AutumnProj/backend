package handler

import (
	mongoDB "boxin/service/check/mongoDB"
	pb "boxin/service/check/proto/check"
	repo "boxin/service/check/repository"

	"context"
	"log"
	"time"
)

// CheckHandler struct
type CheckHandler struct {
	CheckRepository repo.CheckRepository
	CheckMongo      mongoDB.CheckMongo
}

// CreateCheck create a check
func (h *CheckHandler) CreateCheck(ctx context.Context, in *pb.CreateCheckParam, out *pb.CreateCheckResponse) error {
	check := repo.Check{
		CheckTime: time.Unix(in.CheckTime, 0),
		Score: in.Score,
	}
	respCheck, err := h.CheckRepository.AddCheck(ctx, check);
	log.Println(err)

	mongoCheck := mongoDB.Check{
		CheckID:   respCheck.CheckID,
		Description: in.Description,
		Comment: in.Comment,
	}

	err = h.CheckMongo.AddCheck(ctx, mongoCheck)
	log.Println(mongoCheck)
	*out = pb.CreateCheckResponse {
		Status: 0,
		Msg: "CreateCheck success",
		CheckID: respCheck.CheckID,
	}
	return nil
}

// DeleteCheck delete a check by its ID
func (h *CheckHandler) DeleteCheck(ctx context.Context, in *pb.CheckID, out *pb.DeleteCheckResponse) error {
	err := h.CheckRepository.DeleteCheck(ctx, in.CheckID)
	log.Println(err)
	err = h.CheckMongo.DeleteCheck(ctx, in.CheckID)
	log.Println(err)

	out.Status = 0
	out.Msg = "Success"
	return nil
}

// UpdateCheck update a check
func (h *CheckHandler) UpdateCheck(ctx context.Context, in *pb.UpdateCheckParam, out *pb.UpdateCheckResponse) error {
	check := repo.Check{
		CheckID:   in.CheckID,
		CheckTime: time.Unix(in.CheckTime, 0),
		Score: in.Score,
	}
	err := h.CheckRepository.UpdateCheck(ctx, check)
	log.Println(err)

	mongoCheck := mongoDB.Check{
		CheckID:   in.CheckID,
		Description: in.Description,
		Comment: in.Comment,
	}
	err = h.CheckMongo.UpdateCheck(ctx, mongoCheck)
	log.Println(err)

	out.Status = 0
	out.Msg = "Success"
	return nil
}

// SearchCheckByID search a check by its ID
func (h *CheckHandler) SearchCheckByID(ctx context.Context, in *pb.CheckID, out *pb.SearchCheckByIDResponse) error {
	// stime := time.Unix(in.StartTime, 0)
	// etime := time.Unix(in.EndTime, 0)
	check, err := h.CheckRepository.SearchCheckByID(ctx, in.CheckID)
	log.Println(err)
	mongoCheck, err := h.CheckMongo.SearchCheckByID(ctx, in.CheckID)
	log.Println(err)

	*out = pb.SearchCheckByIDResponse{
		Status: 0,
		Msg: "SearchCheckByID success",
		Check: &pb.CheckInfo{
			CheckID:   check.CheckID,
			CheckTime: check.CheckTime.Unix(),
			Score: check.Score,
			Description: mongoCheck.Description,
			Comment: mongoCheck.Comment,
		},
	}
	return nil
}

