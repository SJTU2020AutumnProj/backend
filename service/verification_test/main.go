package main

import (
	verification "boxin/service/verification/proto/verification"
	"context"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	ServiceName = "go.micro.client.verification"
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
	verificationService := verification.NewVerificationService("go.micro.service.verification", server.Client())
	sendCode(verificationService, "chengke", "chengke3@163.com")
	verifyCode(verificationService, "chengke", "chengke3@163.com", "499613")
}

func sendCode(verificationService verification.VerificationService, username string, email string) {
	resp, err := verificationService.SendCodeEmail(
		context.Background(),
		&verification.SendCodeEmailParam{
			Email:    email,
			Username: username,
		},
	)
	if nil != err {
		log.Println("sendCode error, ", err)
		return
	}
	if resp.Status != 0 {
		log.Println("sendCode error, ", resp.Message)
		return
	}
	log.Println("sendCode success ", resp.Message)
}

func verifyCode(verificationService verification.VerificationService, username string, email string, code string) {
	resp, err := verificationService.VerifyCodeEmail(
		context.Background(),
		&verification.VerifyCodeEmailParam{
			Email:    email,
			Username: username,
			Code:     code,
		},
	)
	if nil != err {
		log.Println("verifyCode error, ", err)
		return
	}
	if resp.Status != 0 {
		log.Println("verifyCode error, ", resp.Message)
		return
	}
	log.Println("verifyCode success ", resp.Message)
}
