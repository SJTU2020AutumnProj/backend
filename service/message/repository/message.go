package repository

import (
	"context"
	"time"
	"log"
	"github.com/jinzhu/gorm"

	// Mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Message struct
type Message struct {
	MessageID   int32     `gorm:"auto_increment;column:message_id;primary_key:true;unique;index:"`
	MessageTime time.Time `gorm:"not null;column:message_time"`
	// 0 means broadcast, 1 means private message
	MessageType int32  `gorm:"not null;column:message_type"`
	UserID      int32  `gorm:"column:user_id"`
	CourseID    int32  `gorm:"column:course_id"`
	Title       string `gorm:"not null;column:title"`
	// 0 means not read yet, 1 means already read
	State       int32  `gorm:"not null;column:state"`
}

// TableName configure table name
func (Message) TableName() string {
	return "message"
}

// MessageRepository interface
type MessageRepository interface {
	AddMessage(ctx context.Context, message Message) (Message, error)
	SearchMessageByUserID(ctx context.Context, userID int32) ([]*Message, error)
	SearchMessageByCourseID(ctx context.Context, courseID int32) ([]*Message, error)
	SearchNewMessageByUserID(ctx context.Context, userID int32) ([]*Message, error)
}

// MessageRepositoryImpl implementation
type MessageRepositoryImpl struct {
	DB *gorm.DB
}

// AddMessage create a new message
func (repo *MessageRepositoryImpl) AddMessage(ctx context.Context, message Message) (Message, error) {
	if err := repo.DB.Create(&message).Error; nil != err {
		log.Println("Repository AddMessage error ", err)
		return Message{}, err
	}
	return message, nil
}

// SearchMessageByUserID search message by userID
func (repo *MessageRepositoryImpl) SearchMessageByUserID(ctx context.Context, userID int32) ([]*Message, error) {
	var messages []*Message

	result := repo.DB.Table("message").Where("user_id = ?", userID)

	if err := result.Find(&messages).Error; nil != err {
		log.Println("Repository SearchMessageByUserID error ", err)
		return []*Message{}, err
	}
	return messages, nil
}

// SearchMessageByCourseID search message by courseID
func (repo *MessageRepositoryImpl) SearchMessageByCourseID(ctx context.Context, courseID int32) ([]*Message, error) {
	var messages []*Message

	result := repo.DB.Table("message").Where("course_id = ?", courseID)

	if err := result.Find(&messages).Error; nil != err {
		log.Println("Repository SearchMessageByCourseID error ", err)
		return []*Message{}, err
	}
	return messages, nil
}

// SearchNewMessageByUserID search unread message by userID
func (repo *MessageRepositoryImpl) SearchNewMessageByUserID(ctx context.Context, userID int32) ([]*Message, error) {
	var messages []*Message

	result := repo.DB.Table("message").Where("course_id = ? and state = 0", userID)

	if err := result.Find(&messages).Error; nil != err {
		log.Println("Repository SearchNewMessageByUserID error ", err)
		return []*Message{}, err
	}
	return messages, nil
}
