package mongoDB

import (
	"context"
	"log"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Message struct
type Message struct {
	MessageID   int32 `json:"message_id,omitempty" bson:"message_id,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
}

// TableName configure table name
func (Message) TableName() string {
	return "message"
}

// MessageMongo interface
type MessageMongo interface {
	AddMessage(ctx context.Context, message Message) error
	SearchMessage(ctx context.Context, messageID int32) (Message,error)
}

// MessageMongoImpl struct
type MessageMongoImpl struct {
	CL *mongo.Collection
}

// AddMessage add a new message in mongoDB
func (repo *MessageMongoImpl) AddMessage(ctx context.Context, message Message) error {
	id := message.MessageID
	content := message.Content
	m := Message{id, content}

	_, err := repo.CL.InsertOne(ctx, m)
	if err != nil {
		log.Println("MongoDB Addmessage error ", err)
		return err
	}
	return nil
}

// SearchMessage search message by messageID in mongoDB
func (repo *MessageMongoImpl) SearchMessage(ctx context.Context, messageID int32) (Message,error){
	filter := bson.M{"message_id": messageID}
	var result Message
	err := repo.CL.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("MongoDB SearchMessage error ", err)
		return result,err
	}
	return result,nil
}





