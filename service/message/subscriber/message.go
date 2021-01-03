package subscriber

import (
	homework "boxin/service/homework/proto/homework"
	mongo "boxin/service/message/mongoDB"
	repo "boxin/service/message/repository"
	"context"
	"log"
	"time"
)

// MessageSub struct
type MessageSub struct {
	MessageRepository repo.MessageRepository
	MongoDB           mongo.MessageMongo
}

// Assigned add an assigned homework message
func (sub *MessageSub) Assigned(ctx context.Context, homework *homework.AssignedHomework) error {
	log.Println("Subscriber Assigned received message: ", homework)
	now := time.Unix(time.Now().Unix(), 0)
	message, err := sub.MessageRepository.AddMessage(
		ctx,
		repo.Message{
			MessageTime: now,
			MessageType: 0,
			CourseID:    homework.CourseID,
			Title:       homework.Title + " 已经发布",
			State:       0,
		},
	)
	if nil != err {
		log.Println("Subscriber Assigned add new message in Mysql error ", err)
		return err
	}
	content := homework.Title + " 已经发布，" + "截止时间为 " + time.Unix(homework.EndTime, 0).Format("2006-01-02 15:04:05")
	err = sub.MongoDB.AddMessage(
		ctx,
		mongo.Message{
			MessageID: message.MessageID,
			Content:   content,
		},
	)
	if nil != err {
		log.Println("Subscriber Assigned add new message in MongoDB error ", err)
		return err
	}
	return nil
}
