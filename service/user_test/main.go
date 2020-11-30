package main

import (
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

func main() {
	server := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)
	server.Init()
	userService := user.NewUserService("go.micro.service.user", server.Client())
	// authService := auth.NewAuthService("go.micro.service.auth", server.Client())
	var userNames []string
	var phones []string
	var emails []string
	for i := 0; i < 100; i++ {
		userName := CreateRandomString(5)
		phone := CreateRandomString(5)
		email := CreateRandomString(5)
		userNames = append(userNames, userName)
		phones = append(phones, phone)
		emails = append(emails, email)
		registerStudent(userService, userName, CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), phone, email)
	}
	for i := range userNames {
		registerStudent(userService, userNames[i], CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), CreateRandomString(5))
	}
	for i := range phones {
		registerStudent(userService, CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), phones[i], CreateRandomString(5))
	}
	for i := range emails {
		registerStudent(userService, CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), CreateRandomString(5), emails[i])
	}
	// getUser(userService, 1)
	// addUser(userService, 1, "ck", "12345", "SJTU", "518021910095", "19901714261", "chengke3@163.com")
	// updateUser(userService, 2, 0, "coco", "12345", "SJTU", "518021910095", "19901714261", "chengke3@163.com")
	// getUsers(userService, []int32{1, 2})
	// registerStudent(userService, "ck", "12345", "SJTU", "518021910095", "19901714261", "chengke3@163.com")
	// registerStudent(userService, "dsy", "12345", "SJTU", "id unknown", "phone unknown", "email unknown")
	// registerStudent(userService, "chengke", "12345", "SJTU", "518021910095", "13818646704", "chengke3@163.com")
	// getUser(userService, 3)
	// deleteUser(userService, 3)
	// updateUser(userService, 1, 2, "ck", "123", "SJTU", "518021910095", "19901714261", "chengke3@163.com")
	// user1 := login(authService, "ck", "12345").UserID
	// getUser(userService, user1)
	// updateUser(userService, user1, 2, "ck", "123", "SJTU", "518021910095", "19901714261", "chengke3@163.com")
	// user1 = login(authService, "ck", "123").UserID
	// deleteUser(userService, user1)
	// user2 := login(authService, "dsy", "12345").UserID
	// user3 := login(authService, "chengke", "12345").UserID
	// getUsers(userService, []int32{user2, user3})
}

// func getUser(userService user.UserService, userID int32) (*user.User, error) {
// 	resp, err := userService.SearchUser(context.Background(), &user.UserID{UserID: userID})
// 	if nil != err {
// 		log.Println("getUser error: ", err)
// 		return &user.User{}, err
// 	}
// 	log.Println("getUser success: ", resp)
// 	return resp.User, err
// }

// func deleteUser(userService user.UserService, userID int32) {
// 	resp, err := userService.DeleteUser(context.Background(), &user.UserID{UserID: userID})
// 	if nil != err {
// 		log.Println("deleteUser error: ", err)
// 	}
// 	log.Println("deleteUser success: ", resp)
// }

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

func registerStudent(
	userService user.UserService,
	userName string,
	password string,
	school string,
	id string,
	phone string,
	email string) {
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
	} else {
		log.Println("registerStudent success: ", resp)
	}
}

// func login(authService auth.AuthService, userName string, password string) *auth.UserData {
// 	resp, err := authService.Login(context.Background(), &auth.LoginParam{UserName: userName, Password: password})
// 	if nil != err {
// 		log.Println("login error: ", err)
// 		return &auth.UserData{}
// 	}
// 	log.Println("login success: ", resp)
// 	return resp.Data
// }

// func getUsers(userService user.UserService, userIDs []int32) ([]*user.User, error) {
// 	resp, err := userService.SearchUsers(context.Background(), &user.UserIDArray{IDArray: userIDs})
// 	if err != nil {
// 		log.Println("getUsers error: ", err)
// 		return []*user.User{}, err
// 	}
// 	log.Println("getUsers success: ", resp)
// 	return resp.Users, err
// }

// func updateUser(userService user.UserService, userID int32, userType int32, userName string, password string, school string, ID string, phone string, email string) {
// 	user, err := getUser(userService, userID)
// 	if err != nil {
// 		log.Println("updateUser error: ", err)
// 	}
// 	user.UserType = userType
// 	user.UserName = userName
// 	user.Password = password
// 	user.School = school
// 	user.Id = ID
// 	user.Phone = phone
// 	user.Email = email
// 	resp, erro := userService.UpdateUser(context.Background(), user)
// 	if erro != nil {
// 		log.Println("updateUser error: ", erro)
// 	}
// 	log.Println("updateUser success: ", resp)
// }

// func addUser(userService user.UserService, userType int32, userName string, password string, school string, ID string, phone string, email string) {
// 	resp, err := userService.AddUser(
// 		context.Background(),
// 		&user.User{
// 			UserType: userType,
// 			UserName: userName,
// 			Password: password,
// 			School:   school,
// 			Id:       ID,
// 			Phone:    phone,
// 			Email:    email,
// 		})
// 	if err != nil {
// 		log.Println("addUser error: ", err)
// 		return
// 	}
// 	log.Println("addUser success: ", resp)
// }
