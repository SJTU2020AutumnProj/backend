package handler

import (
	pb "boxin/service/email/proto/email"
	"context"
	"log"

	// "time"

	"github.com/go-gomail/gomail"
)

/*
EmailHandler handler
*/
type EmailHandler struct {
	Val int
}

/*
Configuration of smtp service
*/
const (
	// SMTPMailHost = "smtp.163.com"
	// SMTPMailPort = 25
	SMTPMailHost = "smtp.gmail.com"
	SMTPMailPort = 587
	// SMTPMailUser = "sjtuboxin@163.com"
	SMTPMailUser = "xiaoxiangjiaoJUB@gmail.com"
	SMTPMailName = "boxin"
	// SMTPMailPwd  = "RBMJNDBOVZPEKFJQ"
	SMTPMailPwd = "ASDFGHJKL;"
)

/*
SendEmail send email to the given email address
*/
func (e *EmailHandler) SendEmail(ctx context.Context, in *pb.SendEmailParam, out *pb.SendEmailResponse) error {
	// auth := smtp.PlainAuth(in.Username, SMTPMailUser, SMTPMailPwd, SMTPMailHost)
	// nickname := "Boxin"
	// contentType := "Content-Type: text/html; charset=UTF-8"
	// body := "Dear " + in.Username + "\n" + in.Content
	// msg := []byte("To: " + in.Email + "\r\nFrom: " + nickname + "<" + SMTPMailUser + ">\r\nSubject: " + in.Title +
	// 	"\r\n" + contentType + "\r\n\r\n" + body)
	// err := smtp.SendMail(fmt.Sprintf("%s:%s", SMTPMailHost, SMTPMailPort), auth, SMTPMailUser, []string{in.Email}, msg)
	// if nil != err {
	// 	log.Println("EmailHandler SendCodeEmail error ", err)
	// *out = pb.SendEmailResponse{
	// 	Status:  -1,
	// 	Message: "SMTP send email error",
	// }
	// return err
	// }
	// *out = pb.SendEmailResponse{
	// 	Status:  0,
	// 	Message: "Send email success",
	// }
	// return nil
	// time.Sleep(30 * time.Second)
	log.Println("EmailHandler SendCodeEmail get request")
	m := gomail.NewMessage()

	m.SetAddressHeader("From", SMTPMailUser, SMTPMailName)

	m.SetHeader("To", m.FormatAddress(in.Email, in.Username))

	m.SetHeader("Cc", m.FormatAddress(SMTPMailUser, SMTPMailName))

	m.SetHeader("Subject", in.Title)

	m.SetBody("text/html", "Dear "+in.Username+":<br><br>"+"&nbsp;&nbsp;"+in.Content)
	log.Println("EmailHandler SendCodeEmail begin to send code")
	d := gomail.NewPlainDialer(SMTPMailHost, SMTPMailPort, SMTPMailUser, SMTPMailPwd)
	log.Println("EmailHandler SendCodeEmail client constructed")
	if err := d.DialAndSend(m); err != nil {
		log.Println("EmailHandler SendCodeEmail error ", err)
		*out = pb.SendEmailResponse{
			Status:  -1,
			Message: "SMTP send email error",
		}
		return err
	}
	log.Println("EmailHandler SendCodeEmail send success")
	*out = pb.SendEmailResponse{
		Status:  0,
		Message: "Send email success",
	}
	log.Println("EmailHandler SendCodeEmail return")
	return nil
}
