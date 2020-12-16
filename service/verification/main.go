package main

import (
	"boxin/service/verification/handler"
	verification "boxin/service/verification/proto/verification"
	repo "boxin/service/verification/repository"
	"log"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"
)

/*
Configuration of Verification service
*/
const (
	ServiceName = "go.micro.service.verification"
	RedisHost   = "127.0.0.1"
	RedisPort   = 6379
	EtcdAddr    = "localhost:2379"
)

func main() {
	c, err := redigo.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)

	service.Init()

	verificationHandler := &handler.VerificationHandler{VerificationRepository: &repo.VerificationRepositoryImpl{CONN: c}}

	if err := verification.RegisterVerificationServiceHandler(service.Server(), verificationHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
