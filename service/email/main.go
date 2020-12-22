package main

import (
	"boxin/service/email/handler"
	email "boxin/service/email/proto/email"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"
)

/*
Configuration of Verification service
*/
const (
	ServiceName = "go.micro.service.email"
	EtcdAddr    = "localhost:2379"
)

func main() {
	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)

	service.Init()

	emaillHandler := &handler.EmailHandler{Val: 0}

	if err := email.RegisterEmailServiceHandler(service.Server(), emaillHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
