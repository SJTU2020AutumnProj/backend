package handler

import (
	pb "boxin/service/auth/proto/auth"
	redis "boxin/service/auth/redis"
	repo "boxin/service/auth/repository"
	"boxin/utils"

	// "strconv"
	"context"
	"log"
	// "time"
)

/*
AuthHandler : handler of auth service
*/
type AuthHandler struct {
	AuthRepository repo.AuthRepository
	AuthRedis      redis.AuthRedis
}

/*
Login : user can login by username and password
*/
func (a *AuthHandler) Login(ctx context.Context, in *pb.LoginParam, out *pb.LoginResponse) error {
	user, err := a.AuthRepository.Login(ctx, in.UserName, in.Password)
	token, JWTerr := utils.JWTSign(user.UserID, user.UserName, user.Password, user.UserType)
	log.Println(JWTerr)
	*out = pb.LoginResponse{
		Status: 0,
		Msg:    "Success",
		Data: &pb.UserData{
			UserID:   user.UserID,
			UserType: user.UserType,
			UserName: user.UserName,
			Password: user.Password,
			School:   user.School,
			Id:       user.ID,
			Phone:    user.Phone,
			Email:    user.Email,
		},
		Token: token,
	}
	return err
}
