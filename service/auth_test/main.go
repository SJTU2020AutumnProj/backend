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
	// validLogin(authService)
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
	username := "ck"
	password := "123"
	resp, err := login(authService, username, password)
	if nil == err {
		log.Println("Login success: ", resp.Token)
		token = resp.Token
	}
}

func checkAuth(authService auth.AuthService) {
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjMwNCwidXNlcm5hbWUiOiJjayIsInBhc3N3b3JkIjoiMTIzIiwidXNlcnR5cGUiOjAsImV4cCI6MTYwNzQzNDM1MiwiaXNzIjoiYm94aW4ifQ.5v6HW9aubw3f9C1XIGnA2xfBN7NdK_nldag-EU4mAdc"

	resp, err := authService.CheckAuth(
		context.Background(),
		&auth.CheckAuthParam{Token: token},
	)
	if nil != err {
		log.Println("CheckAuth error: ", err)
	} else {
		log.Println("CheckAuth success: ", resp.Data)
	}
}
