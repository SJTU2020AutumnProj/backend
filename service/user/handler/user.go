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
		School:   req.School,
		ID:       req.Id,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := u.UserRepository.AddUser(ctx, user); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("UserHandler AddUser error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (u *UserHandler) RegisterAdmin(ctx context.Context, req *pb.UserInfo, resp *pb.RegisterResponse) error {
	user := repo.User{
		UserType: 0,
		UserName: req.UserName,
		Password: req.Password,
		School:   req.School,
		ID:       req.Id,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := u.UserRepository.AddUser(ctx, user); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("UserHandler RegisterAdmin error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (u *UserHandler) RegisterTeacher(ctx context.Context, req *pb.UserInfo, resp *pb.RegisterResponse) error {
	user := repo.User{
		UserType: 1,
		UserName: req.UserName,
		Password: req.Password,
		School:   req.School,
		ID:       req.Id,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := u.UserRepository.AddUser(ctx, user); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("UserHandler RegisterTeacher error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (u *UserHandler) RegisterStudent(ctx context.Context, req *pb.UserInfo, resp *pb.RegisterResponse) error {
	user := repo.User{
		UserType: 2,
		UserName: req.UserName,
		Password: req.Password,
		School:   req.School,
		ID:       req.Id,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := u.UserRepository.AddUser(ctx, user); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("UserHandler RegisterStudent error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}



func (u *UserHandler) DeleteUser(ctx context.Context, req *pb.UserID, resp *pb.EditResponse) error {
	if err := u.UserRepository.DeleteUser(ctx, req.UserID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("UserHandler DeleteUser error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (u *UserHandler) UpdateUser(ctx context.Context, req *pb.User, resp *pb.EditResponse) error {
	user := repo.User{
		UserID:   req.UserID,
		UserType: req.UserType,
		UserName: req.UserName,
		Password: req.Password,
		School:   req.School,
		ID:       req.Id,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := u.UserRepository.UpdateUser(ctx, user); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("UserHandler UpdateUser error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (u *UserHandler) SearchUser(ctx context.Context, req *pb.UserID, resp *pb.SearchResponse) error {
	user, err := u.UserRepository.SearchUser(ctx, req.UserID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchUser error: ", err)
		return err
	}

	*resp = pb.SearchResponse{
		Status: 0,
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
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("UserHandler SearchUsers error: ", err)
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
		Status: 0,
		Msg:    "Success",
		Users:  users,
	}
	return nil
}
