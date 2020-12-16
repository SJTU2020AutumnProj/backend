package handler

import (
	pb "boxin/service/verification/proto/verification"
	repo "boxin/service/verification/repository"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

/*
VerificationHandler handler
*/
type VerificationHandler struct {
	VerificationRepository repo.VerificationRepository
}

/*
Configuration of smtp service
*/
const (
	SMTPMailHost = "smtp.163.com"
	SMTPMailPort = "25"
	SMTPMailUser = "sjtuboxin@163.com"
	SMTPMailPwd  = "RBMJNDBOVZPEKFJQ"
	ExpireTime   = "300"
)

/*
SendCodeEmail send verification code to the given email address
*/
func (v *VerificationHandler) SendCodeEmail(ctx context.Context, in *pb.SendCodeEmailParam, out *pb.SendCodeEmailResponse) error {
	auth := smtp.PlainAuth(in.Username, SMTPMailUser, SMTPMailPwd, SMTPMailHost)
	nickname := "Boxin"
	contentType := "Content-Type: text/html; charset=UTF-8"
	randomNumber := strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)))
	body := "Dear " + in.Username + ", your verification code is " + "<b>" + randomNumber + "</b>"
	msg := []byte("To: " + in.Email + "\r\nFrom: " + nickname + "<" + SMTPMailUser + ">\r\nSubject: " + "Your verification code" +
		"\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(fmt.Sprintf("%s:%s", SMTPMailHost, SMTPMailPort), auth, SMTPMailUser, []string{in.Email}, msg)
	if nil != err {
		log.Println("VerificationHandler SendCodeEmail error ", err)
		*out = pb.SendCodeEmailResponse{
			Status:  -1,
			Message: "SMTP send email error",
		}
		return err
	}
	setErr := v.VerificationRepository.Set(ctx, in.Email, randomNumber, ExpireTime)
	if nil != setErr {
		log.Println("VerificationHandler redis error ", setErr)
		*out = pb.SendCodeEmailResponse{
			Status:  -1,
			Message: "Redis set key value error",
		}
		return setErr
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
