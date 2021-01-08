package main

import (
	email "boxin/service/email/proto/email"
	"context"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	ServiceName = "go.micro.client.email"
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
	emailService := email.NewEmailService("go.micro.service.email", server.Client())
	send(emailService, "chengke3@163.com", "chengke", "Hello", "This is a hello email")
}

func send(emailService email.EmailService, Email string, username string, title string, content string) {
	resp, err := emailService.SendEmail(
		context.Background(),
		&email.SendEmailParam{
			Email:    Email,
			Username: username,
			Title:    title,
			Content:  content,
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
