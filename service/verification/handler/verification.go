package handler

import (
	email "boxin/service/email/proto/email"
	pb "boxin/service/verification/proto/verification"
	repo "boxin/service/verification/repository"
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/micro/go-micro/v2/client"
)

/*
VerificationHandler handler
*/
type VerificationHandler struct {
	VerificationRepository repo.VerificationRepository
}

/*
Configuration of verification
*/
const (
	ExpireTime = "300"
)

/*
SendCodeEmail send verification code to the given email address
*/
func (v *VerificationHandler) SendCodeEmail(ctx context.Context, in *pb.SendCodeEmailParam, out *pb.SendCodeEmailResponse) error {
	randomNumber := strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)))
	content := "\tYour verification code is " + "<b>" + randomNumber + "</b>"
	title := "Your verification code"
	log.Println("VerificationHandler SendCodeEmail begin to send code email")
	emailService := email.NewEmailService("go.micro.service.email", client.DefaultClient)
	log.Println("VerificationHandler SendCodeEmail client constructed")
	resp, err := emailService.SendEmail(
		ctx,
		&email.SendEmailParam{
			Email:    in.Email,
			Username: in.Username,
			Title:    title,
			Content:  content,
		},
	)
	if nil != err {
		log.Println("VerificationHandler SendCodeEmail error ", resp.Message)
		*out = pb.SendCodeEmailResponse{
			Status:  -1,
			Message: "SendCodeEmail error",
		}
		return err
	}
	log.Println("VerificationHandler SendCodeEmail send email success")
	log.Println("VerificationHandler SendCodeEmail begin to store email address and code")
	setErr := v.VerificationRepository.Set(ctx, in.Email, randomNumber, ExpireTime)
	if nil != setErr {
		log.Println("VerificationHandler redis error ", setErr)
		*out = pb.SendCodeEmailResponse{
			Status:  -1,
			Message: "Redis set key value error",
		}
		return setErr
	}
	log.Println("VerificationHandler SendCodeEmail store email address and code success")
	*out = pb.SendCodeEmailResponse{
		Status:  0,
		Message: "SendCodeEmail success",
	}
	return nil
}

/*
VerifyCodeEmail verify the code
*/
func (v *VerificationHandler) VerifyCodeEmail(ctx context.Context, in *pb.VerifyCodeEmailParam, out *pb.VerifyCodeEmailResponse) error {
	actualCode, err := v.VerificationRepository.Get(ctx, in.Email)
	if nil != err {
		log.Println("VerificationHandler reids error ", err)
		*out = pb.VerifyCodeEmailResponse{
			Status:  -1,
			Message: "Redis error: " + err.Error(),
		}
		return err
	}
	if !(actualCode == in.Code) {
		log.Println("VerificationHandler verification code mismatch, given ", in.Code, "actual ", actualCode)
		*out = pb.VerifyCodeEmailResponse{
			Status:  -1,
			Message: "Verification code mismatch",
		}
		return nil
	}
	*out = pb.VerifyCodeEmailResponse{
		Status:  0,
		Message: "Verified",
	}
	return nil
}
