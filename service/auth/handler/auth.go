package handler

import (
	pb "boxin/service/auth/proto/auth"
	repo "boxin/service/auth/repository"
	"context"
	"log"
)

type AuthHandler struct {
	AuthRepository repo.AuthRepository
}

func (a *AuthHandler) Login(ctx context.Context, req *pb.LoginParam, resp *pb.LoginResponse) error {
	user, err := a.AuthRepository.Login(ctx, req.UserName, req.Password)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("AuthHandler Login error: ", err)
		return err
	}
	*resp = pb.LoginResponse{
		Status: 0,
		Msg: "Success",
		Data: &pb.UserData{
			UserID: user.UserID,
			UserType: user.UserType,
			UserName: user.UserName,
			School: user.School,
			Id: user.ID,
			Phone: user.Phone,
			Email: user.Email,
		},
	}
	return nil
}