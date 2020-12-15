package main

import (
	auth "boxin/service/auth/proto/auth"
	user "boxin/service/user/proto/user"
	"context"
	"log"

	"bytes"
	"crypto/rand"
	"math/big"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	ServiceName = "go.micro.client.user"
	EtcdAddr    = "localhost:2379"
)

var userNames []string
var passwords []string
var userIDs []int32

func main() {
	server := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)
	server.Init()
	userService := user.NewUserService("go.micro.service.user", server.Client())
	authService := auth.NewAuthService("go.micro.service.auth", server.Client())

	registerTest(userService)
	loginTest(authService)
}

/*
CreateRandomString : create random string by the given length
*/
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func registerTest(userService user.UserService) {
	var phones []string
	var emails []string
	for i := 0; i < 100; i++ {
		userName := CreateRandomString(5)
		password := CreateRandomString(5)
		phone := CreateRandomString(5)
		email := CreateRandomString(5)
		userNames = append(userNames, userName)
		passwords = append(passwords, password)
		phones = append(phones, phone)
		emails = append(emails, email)
		userID := registerStudent(userService, userName, password, CreateRandomString(5), CreateRandomString(5), phone, email)
		userIDs = append(userIDs, userID)
	}
	for i := range userNames {
		registerStudent(userService, userNames[i], passwords[i], CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), CreateRandomString(5))
	}
	for i := range phones {
		registerStudent(userService, CreateRandomString(5), passwords[i], CreateRandomString(5), CreateRandomString(5), phones[i], CreateRandomString(5))
	}
	for i := range emails {
		registerStudent(userService, CreateRandomString(5), passwords[i], CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), emails[i])
	}
}

func registerStudent(
	userService user.UserService,
	userName string,
	password string,
	school string,
	id string,
	phone string,
	email string) int32{
	resp, err := userService.RegisterStudent(
		context.Background(),
		&user.RegisterUserParam{
			UserName: userName,
			Password: password,
			School:   school,
			ID:       id,
			Phone:    phone,
			Email:    email,
		},
	)
	if nil != err {
		log.Println("registeStudent error: ", err)
		return 0
	} else {
		log.Println("registerStudent success: ", resp)
		result := resp.UserID.UserID
		return result
	}
}

func loginTest(authService auth.AuthService) {
	for i := range userNames {
		resp, err := authService.Login(
			context.Background(),
			&auth.LoginParam{
				UserName: userNames[i],
				Password: passwords[i],
			},
		)
		if nil != err {
			log.Println("login error: ", err)
		} else {
			log.Println("login success: ", resp.Data.UserID)
		}
		if(resp.Data.UserID != userIDs[i]) {
			log.Println("login error: userID different")
		}
	}
}
