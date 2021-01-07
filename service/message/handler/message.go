package handler

import (
	mongoDB "boxin/service/message/mongoDB"
	pb "boxin/service/message/proto/message"
	repo "boxin/service/message/repository"

	"context"
	"log"
)

// MessageHandler struct
type MessageHandler struct {
	MessageRepository repo.MessageRepository
	MessageMongo      mongoDB.MessageMongo
}

// GetMessageByUserID return messages according to the given userID
func (h *MessageHandler) GetMessageByUserID(ctx context.Context, in *pb.GetMessageByUserIDParam, out *pb.GetMessageByUserIDResponse) error {
	resp, err := h.MessageRepository.SearchMessageByUserID(ctx, in.UserID)
	if nil != err {
		log.Println("MessageHandler GetMessageByUserID error ", err)
		*out = pb.GetMessageByUserIDResponse{
			Status: -1,
			Msg:    "Get message by userID error",
		}
		return err
	}
	var result []*pb.Message
	for i := range resp {
		messageContent, err := h.MessageMongo.SearchMessage(ctx, resp[i].MessageID)
		if nil != err {
			log.Println("MessageHandler GetMessageByUserID error ", err)
			out = &pb.GetMessageByUserIDResponse{
				Status: -1,
				Msg:    "Get message by userID error",
			}
			return err
		}
		currentResult := &pb.Message{
			MessageID:   resp[i].MessageID,
			MessageTime: resp[i].MessageTime.Unix(),
			MessageType: resp[i].MessageType,
			UserID:      resp[i].UserID,
			CourseID:    resp[i].CourseID,
			Title:       resp[i].Title,
			Content:     messageContent.Content,
			State:       resp[i].State,
		}
		result = append(result, currentResult)
	}
	*out = pb.GetMessageByUserIDResponse{
		Status: 0,
		Msg:    "Get message by userID success",
		Data:   result,
	}
	log.Println("MessageHandler GetMessageByUserID success ", result)
	return nil
}

// GetMessageByCourseID return messages
func (h *MessageHandler) GetMessageByCourseID(ctx context.Context, in *pb.GetMessageByCourseIDParam, out *pb.GetMessageByCourseIDResponse) error {
	resp, err := h.MessageRepository.SearchMessageByCourseID(ctx, in.CourseID)
	if nil != err {
		log.Println("MessageHandler GetMessageByCourseID error ", err)
		*out = pb.GetMessageByCourseIDResponse{
			Status: -1,
			Msg:    "Get message by courseID error",
		}
		return err
	}
	var result []*pb.Message
	for i := range resp {
		messageContent, err := h.MessageMongo.SearchMessage(ctx, resp[i].MessageID)
		if nil != err {
			log.Println("MessageHandler GetMessageByCourseID error ", err)
			*out = pb.GetMessageByCourseIDResponse{
				Status: -1,
				Msg:    "Get message by courseID error",
			}
			return err
		}
		currentResult := &pb.Message{
			MessageID:   resp[i].MessageID,
			MessageTime: resp[i].MessageTime.Unix(),
			MessageType: resp[i].MessageType,
			UserID:      resp[i].UserID,
			CourseID:    resp[i].CourseID,
			Title:       resp[i].Title,
			Content:     messageContent.Content,
			State:       resp[i].State,
		}
		result = append(result, currentResult)
	}
	*out = pb.GetMessageByCourseIDResponse{
		Status: 0,
		Msg:    "Get message by courseID success",
		Data:   result,
	}
	log.Println("MessageHandler GetMessageByCourseID success ", result)
	return nil
}
