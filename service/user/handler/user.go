package handler

import (
	pb "boxin/service/user/proto/user"
	repo "boxin/service/user/repository"
	"context"
	"log"
)

type UserHandler struct {
	UserRepository repo.UserRepository
}

func (u *UserHandler) AddUser(ctx context.Context, req *pb.User, resp *pb.EditResponse) error {
	user := repo.User{
		UserType: req.UserType,
		UserName: req.UserName,
		Password: req.Password,
		School: req.School,
		ID: req.Id,
		Phone: req.Phone,
		Email: req.Email,
	}
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
	user := u.UserRepository.GenerateUser(
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
	resp.Msg = "Success"
	return nil
}

func (u *UserHandler) SearchUser(ctx context.Context, req *pb.UserID, resp *pb.SearchResponse) error {
	user, err := u.UserRepository.SearchUser(ctx, req.GetUserID())
	if err != nil {
		resp.Status = -1
		log.Println("Handler SearchUser", err)
		return err
	}

	*resp = pb.SearchResponse{
		Status: 1,
		User: &pb.User{
			UserID:   user.UserID,
			UserType: user.UserType,
			UserName: user.UserName,
			Password: user.Password,
			School:   user.School,
			Id:       user.ID,
			Phone:    user.Phone,
			Email:    user.Email,
		},
	}
	return nil
}

func (u *UserHandler) SearchUsers(ctx context.Context, req *pb.UserIDArray, resp *pb.SearchUsersResponse) error {
	var users []*pb.User
	for i := range req.IDArray {
		user, err := u.UserRepository.SearchUser(ctx, req.IDArray[i])
		if err != nil {
			resp.Status = -1
			log.Fatal("UserHandler SearchUsers", err)
			return err
		}
		users = append(users, &pb.User{
			UserID:   user.UserID,
			UserType: user.UserType,
			UserName: user.UserName,
			Password: user.Password,
			School:   user.School,
			Id:       user.ID,
			Phone:    user.Phone,
			Email:    user.Email,
		})
	}
	*resp = pb.SearchUsersResponse{
		Status: 1,
		Users: users, 
	}
	return nil
}
