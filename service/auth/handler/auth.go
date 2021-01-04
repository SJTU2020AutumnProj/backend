package handler

import (
	pb "boxin/service/auth/proto/auth"
	repo "boxin/service/auth/repository"
	redis "boxin/service/auth/redis"
	"boxin/utils"
	"strconv"
	"context"
	"log"
	"time"
)

/*
AuthHandler : handler of auth service
*/
type AuthHandler struct {
	AuthRepository repo.AuthRepository
	AuthRedis redis.AuthRedis
}

/*
Login : user can login by username and password
*/
func (a *AuthHandler) Login(ctx context.Context, in *pb.LoginParam, out *pb.LoginResponse) error {
	user, err := a.AuthRepository.Login(ctx, in.UserName, in.Password)
	if nil != err {
		*out = pb.LoginResponse{
			Status: -1,
			Msg:    "Error",
			Data:   nil,
		}
		log.Println("AuthHandler Login error: ", err)
		return err
	}
	token, JWTerr := utils.JWTSign(user.UserID, user.UserName, user.Password, user.UserType)
	if nil != JWTerr {
		*out = pb.LoginResponse{
			Status: -1,
			Msg:    "Error",
			Data:   nil,
		}
		log.Println("AuthHandler Login error: ", JWTerr)
		return JWTerr
	}
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
	return nil
}

/*
CheckAuth : check user's permission
*/
func (a *AuthHandler) CheckAuth(ctx context.Context, in *pb.CheckAuthParam, out *pb.CheckAuthResponse) error {
	_, err := a.AuthRedis.Get(ctx, in.Token)
	if nil == err {
		log.Println("AuthHandler CheckAuth reids error ", err)
		*out = pb.CheckAuthResponse{
			Status:  -1,
			Msg: "Token already logout",
			Data: nil,
		}
		return nil
	}
	claims, err := utils.JWTVerify(in.Token)
	if err != nil {
		*out = pb.CheckAuthResponse{
			Status: -1,
			Msg:    "Invalid token",
			Data:   nil,
		}
		log.Println("AuthHandler CheckAuth error: ", err)
		return err
	}
	*out = pb.CheckAuthResponse{
		Status: 0,
		Msg:    "CheckAuth success",
		Data: &pb.AuthData{
			UserID:   claims.UserID,
			UserType: claims.UserType,
			UserName: claims.UserName,
			Password: claims.Password,
		},
	}
	return nil
}

/*
Logout : invalid the token
*/
func (a *AuthHandler) Logout(ctx context.Context, in *pb.LogoutParam, out *pb.LogoutResponse) error {
	claims, err := utils.JWTVerify(in.Token)
	if nil != err {
		*out = pb.LogoutResponse{
			Status: 0,
			Msg:    "Logout success",
		}
		return nil
	}

	expTime := claims.StandardClaims.ExpiresAt
	now := time.Now().Unix()
	expireTime := expTime - now
	if expireTime < 0 {
		expireTime = 0
	}
	setErr := a.AuthRedis.Set(ctx, in.Token, "", strconv.FormatInt(expireTime, 10))
	if nil != setErr {
		log.Println("AuthHandler Logout redis error ", setErr)
		*out = pb.LogoutResponse{
			Status:  -1,
			Msg: "Redis set key value error",
		}
		return setErr
	}

	*out = pb.LogoutResponse{
		Status:  0,
		Msg: "Logout success",
	}
	return nil
}
