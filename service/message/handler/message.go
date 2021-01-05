package handler

import (
	mongoDB "boxin/service/message/mongoDB"
	pb "boxin/service/message/proto/message"
	repo "boxin/service/message/repository"

	"context"
	"log"
	"time"
)

// MessageHandler struct
type MessageHandler struct {
	MessageRepository repo.MessageRepository
	MessageMongo      mongoDB.MessageMongo
}

// GetMessageByUserID return messages according to the given userID
func (h *MessageHandler) GetMessageByUserID(ctx context.Context, in *pb.GetMessageByUserIDParam, out *pb.GetMessageByUserIDResponse) error {
	resp, err := h.MessageRepository.SearchMessageByUserID(ctx, in.UserID);
	if nil != err {
		log.Println("MessageHandler GetMessageByUserID error ", err)
		out = &pb.GetMessageByUserIDResponse {
			Status: -1,
			Msg: "Get message by userID error",
		}
		return err
	}
	var result []*pb.Message
	for i := range resp {
		messageContent := h.MessageMongo.SearchMessage(ctx, resp[i].MessageID)
	}
}

// GetMessageByCourseID return messages
func (h *MessageHandler) GetMessageByCourseID(ctx context.Context, in *pb.GetMessageByCourseIDParam, out *pb.GetMessageByCourseIDResponse) error {
}
