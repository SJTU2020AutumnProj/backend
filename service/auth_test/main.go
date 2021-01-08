package main

import (
	auth "boxin/service/auth/proto/auth"
	"context"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	ServiceName = "go.micro.client.auth"
	EtcdAddr    = "localhost:2379"
)

type User struct {
	UserID   int32
	UserName string
	UserType int32
	Password string
	School   string
	ID       string
	Phone    string
	Email    string
}

type LoginResponse struct {
	Data  User
	Token string
}

var token string

func main() {
	server := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)
	server.Init()
	authService := auth.NewAuthService("go.micro.service.auth", server.Client())
	validLogin(authService)
	checkAuth(authService)
	logout(authService)
	checkAuth(authService)
}

func login(authService auth.AuthService, username string, password string) (LoginResponse, error) {
	resp, err := authService.Login(
		context.Background(),
		&auth.LoginParam{
			UserName: username,
			Password: password,
		},
	)
	if nil != err {
		log.Println("Login error: ", resp.Msg)
		return LoginResponse{}, err
	}
	user := User{
		UserID:   resp.Data.UserID,
		UserName: resp.Data.UserName,
		UserType: resp.Data.UserType,
		Password: resp.Data.Password,
		ID:       resp.Data.Id,
		School:   resp.Data.School,
		Email:    resp.Data.Email,
		Phone:    resp.Data.Phone,
	}
	loginResponse := LoginResponse{
		Data:  user,
		Token: resp.Token,
	}
	return loginResponse, nil
}

func validLogin(authService auth.AuthService) {
	username := "chengke"
	password := "12345"
	resp, err := login(authService, username, password)
	if nil == err {
		log.Println("Login success: ", resp.Token)
		token = resp.Token
	}
}

func checkAuth(authService auth.AuthService) {
	resp, _ := authService.CheckAuth(
		context.Background(),
		&auth.CheckAuthParam{Token: token},
	)
	if 0 != resp.Status {
		log.Println("CheckAuth error: ", resp.Msg)
	} else {
		log.Println("CheckAuth success: ", resp.Msg)
	}
}

func logout(authService auth.AuthService) {
	resp, err := authService.Logout(
		context.Background(),
		&auth.LogoutParam{Token: token},
	)
	if nil != err {
		log.Println("Logout error: ", err)
	} else {
		log.Println("Logout success: ", resp.Msg)
	}
}
