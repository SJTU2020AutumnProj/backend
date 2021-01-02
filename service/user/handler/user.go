package handler

import (
	pb "boxin/service/user/proto/user"
	repo "boxin/service/user/repository"
	"context"
	"errors"
	"log"
)

/*
UserHandler : handle requests regarding user
*/
type UserHandler struct {
	UserRepository repo.UserRepository
}

/*
RegisterAdmin : register administrator user
*/
func (u *UserHandler) RegisterAdmin(ctx context.Context, in *pb.RegisterUserParam, out *pb.RegisterUserResponse) error {
	_, erro := u.UserRepository.SearchUserByUserName(ctx, in.UserName)
	if nil == erro {
		log.Println("RegisterAdmin error: duplicate username")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate username",
		}
		return errors.New("Duplicate username")
	}
	_, erro = u.UserRepository.SearchUserByPhone(ctx, in.Phone)
	if nil == erro {
		log.Println("RegisterAdmin error: duplicate phone")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate phone",
		}
		return errors.New("Duplicate phone")
	}
	_, erro = u.UserRepository.SearchUserByEmail(ctx, in.Email)
	if nil == erro {
		log.Println("RegisterAdmin error: duplicate email")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate email",
		}
		return errors.New("Duplicate email")
	}
	user := repo.User{
		UserType: 0,
		UserName: in.UserName,
		Password: in.Password,
		School:   in.School,
		ID:       in.ID,
		Phone:    in.Phone,
		Email:    in.Email,
		Name:     in.Name,
	}
	res, err := u.UserRepository.AddUser(ctx, user)
	// if nil != err {
	// 	log.Println("RegisterAdmin error: ", err)
	// 	*out = pb.RegisterUserResponse{
	// 		Status: -1,
	// 		Msg:    "Error",
	// 		UserID: nil,
	// 	}
	// 	return err
	// }
	*out = pb.RegisterUserResponse{
		Status: 0,
		Msg:    "Success",
		UserID: &pb.UserID{UserID: res.UserID},
	}
	return err
}

/*
RegisterTeacher : register teacher user
*/
func (u *UserHandler) RegisterTeacher(ctx context.Context, in *pb.RegisterUserParam, out *pb.RegisterUserResponse) error {
	_, erro := u.UserRepository.SearchUserByUserName(ctx, in.UserName)
	if nil == erro {
		log.Println("RegisterTeacher error: duplicate username")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate username",
		}
		return errors.New("Duplicate username")
	}
	_, erro = u.UserRepository.SearchUserByPhone(ctx, in.Phone)
	if nil == erro {
		log.Println("RegisterTeacher error: duplicate phone")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate phone",
		}
		return errors.New("Duplicate phone")
	}
	_, erro = u.UserRepository.SearchUserByEmail(ctx, in.Email)
	if nil == erro {
		log.Println("RegisterTeacher error: duplicate email")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate email",
		}
		return errors.New("Duplicate email")
	}
	user := repo.User{
		UserType: 1,
		UserName: in.UserName,
		Password: in.Password,
		School:   in.School,
		ID:       in.ID,
		Phone:    in.Phone,
		Email:    in.Email,
		Name:     in.Name,
	}
	res, err := u.UserRepository.AddUser(ctx, user)
	// if nil != err {
	// 	log.Println("RegisterTeacher error: ", err)
	// 	*out = pb.RegisterUserResponse{
	// 		Status: -1,
	// 		Msg:    "Error",
	// 		UserID: nil,
	// 	}
	// 	return err
	// }
	*out = pb.RegisterUserResponse{
		Status: 0,
		Msg:    "Success",
		UserID: &pb.UserID{UserID: res.UserID},
	}
	return err
}

/*
RegisterStudent : register student user
*/
func (u *UserHandler) RegisterStudent(ctx context.Context, in *pb.RegisterUserParam, out *pb.RegisterUserResponse) error {
	_, erro := u.UserRepository.SearchUserByUserName(ctx, in.UserName)
	if nil == erro {
		log.Println("RegisterStudent error: duplicate username")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate username",
		}
		return errors.New("Duplicate username")
	}
	_, erro = u.UserRepository.SearchUserByPhone(ctx, in.Phone)
	if nil == erro {
		log.Println("RegisterStudent error: duplicate phone")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate phone",
		}
		return errors.New("Duplicate phone")
	}
	_, erro = u.UserRepository.SearchUserByEmail(ctx, in.Email)
	if nil == erro {
		log.Println("RegisterStudent error: duplicate email")
		*out = pb.RegisterUserResponse{
			Status: -1,
			Msg:    "Duplicate email",
		}
		return errors.New("Duplicate email")
	}
	user := repo.User{
		UserType: 2,
		UserName: in.UserName,
		Password: in.Password,
		School:   in.School,
		ID:       in.ID,
		Phone:    in.Phone,
		Email:    in.Email,
		Name:     in.Name,
	}
	res, err := u.UserRepository.AddUser(ctx, user)
	// if nil != err {
	// 	log.Println("RegisterStudent error: ", err)
	// 	*out = pb.RegisterUserResponse{
	// 		Status: -1,
	// 		Msg:    "Error",
	// 		UserID: nil,
	// 	}
	// 	return err
	// }
	*out = pb.RegisterUserResponse{
		Status: 0,
		Msg:    "Success",
		UserID: &pb.UserID{UserID: res.UserID},
	}
	return err
}

/*
UpdateUser : update user information(except for UserID, UserType and Password)
*/
func (u *UserHandler) UpdateUser(ctx context.Context, in *pb.UpdateUserParam, out *pb.UpdateUserResponse) error {
	user := repo.User{
		UserID:   in.UserID,
		UserType: in.UserType,
		UserName: in.UserName,
		Password: in.Password,
		School:   in.School,
		ID:       in.ID,
		Phone:    in.Phone,
		Email:    in.Email,
		Name:     in.Name,
	}
	err := u.UserRepository.UpdateUser(ctx, user)
	if nil != err {
		log.Println("UpdateUser error: ", err)
		*out = pb.UpdateUserResponse{
			Status: -1,
			Msg:    "Error",
		}
		return err
	}
	*out = pb.UpdateUserResponse{
		Status: 0,
		Msg:    "Success",
	}
	return nil
}

/*
SearchUser : search one user by the given UserID
*/
func (u *UserHandler) SearchUser(ctx context.Context, in *pb.UserID, out *pb.SearchUserResponse) error {
	user, err := u.UserRepository.SearchUser(ctx, in.UserID)
	if nil != err {
		log.Println("SearchUser error: ", err)
		*out = pb.SearchUserResponse{
			Status: -1,
			Msg:    "Error",
			User:   &pb.UserInfo{},
		}
		return err
	}
	*out = pb.SearchUserResponse{
		Status: 0,
		Msg:    "Success",
		User: &pb.UserInfo{
			UserID:   user.UserID,
			UserType: user.UserType,
			UserName: user.UserName,
			School:   user.School,
			ID:       user.ID,
			Phone:    user.Phone,
			Email:    user.Email,
			Name:     user.Name,
		},
	}
	return nil
}

/*
SearchUsers : search multiple users by the given UserIDs(UserIDArray)
*/
func (u *UserHandler) SearchUsers(ctx context.Context, in *pb.UserIDArray, out *pb.SearchUsersResponse) error {
	var users []*pb.UserInfo
	for i := range in.UserIDArray {
		user, err := u.UserRepository.SearchUser(ctx, in.UserIDArray[i])
		if nil != err {
			log.Println("SearchUsers error: ", err)
			*out = pb.SearchUsersResponse{
				Status: -1,
				Msg:    "Error",
				Users:  []*pb.UserInfo{},
			}
		}
		users = append(users, &pb.UserInfo{
			UserID:   user.UserID,
			UserType: user.UserType,
			UserName: user.UserName,
			School:   user.School,
			ID:       user.ID,
			Phone:    user.Phone,
			Email:    user.Email,
			Name:     user.Name,
		})
	}
	*out = pb.SearchUsersResponse{
		Status: 0,
		Msg:    "Success",
		Users:  users,
	}
	return nil
}
