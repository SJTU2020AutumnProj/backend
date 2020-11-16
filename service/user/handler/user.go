package handler

import (
	"context"
	"github.com/micro/go-micro/v2"
	pb "boxin/service/user/proto/user"
	repo "boxin/service/user/repository"
)

type UserHandler struct {
	UserRepository repo.UserRepository
	TaskFinishedPubEvent micro.Event
}

func (u *UserHandler) AddUser(ctx context.Context, req *pb.User, resp *pb.EditResponse) error {
	user := u.UserRepository.GenerateUser (
		req.UserID,
		req.UserType,
		req.UserName,
		req.Password,
		req.School,
		req.Id,
		req.Phone,
		req.Email)
	if err := u.UserRepository.AddUser(ctx, user); err != nil {
		resp.Status = -1
		resp.Msg = "Error"
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (u *UserHandler) DeleteUser(ctx context.Context, req *pb.UserID, resp *pb.EditResponse) error {
	if err := u.UserRepository.DeleteUser(ctx, req.UserID); err != nil {
		resp.Status = -1
		resp.Msg = "Error"
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (u *UserHandler) UpdateUser(ctx context.Context, req *pb.User, resp *pb.EditResponse) error {
	user := u.UserRepository.GenerateUser (
		req.UserID,
		req.UserType,
		req.UserName,
		req.Password,
		req.School,
		req.Id,
		req.Phone,
		req.Email)
	if err := u.UserRepository.UpdateUser(ctx, user); err != nil {
		resp.Status = -1
		resp.Msg = "Error"
		return err
	}
	resp.Status = 0
	resp.Msg = "SUccess"
	return nil
}

func (u *UserHandler) SearchUser(ctx context.Context, req *pb.UserID, resp *pb.SearchResponse) error {
	user, err := u.UserRepository.SearchUser(ctx, req.GetUserID())
	if err != nil {
		resp.Status = -1
		return err
	}
	resp.Status = 1
	resp.User.UserID = user.UserID
	resp.User.UserType = user.UserType
	resp.User.UserName = user.UserName
	resp.User.Password = user.Password
	resp.User.School = user.School
	resp.User.Id = user.ID
	resp.User.Phone = user.Phone
	resp.User.Email = user.Email
	return nil
}
