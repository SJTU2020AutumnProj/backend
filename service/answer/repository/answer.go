package repository

import (
	"context"
	// "log"
	"time"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Answer struct {
	AnswerID int32 `gorm:"auto_increment;column:answer_id;primary_key:true;unique;index:"`
	HomeworkID int32 `gorm:"not null;column:homework_id"`
	StudentID int32 `gorm:"not null;column:student_id"`
	Status int32 `gorm:"not null;column:status"`
	CommitTime time.Time `gorm:"not null;column:commit_time"`
}

func (Answer) TableName() string {
	return "answer"
}

type AnswerRepository interface {
	AddAnswer(ctx context.Context, answer Answer)(Answer, error)
	DeleteAnswer(ctx context.Context, answerID int32) error
	UpdateAnswer(ctx context.Context, answer Answer) error
	SearchAnswer(ctx context.Context, answerID int32) (Answer, error)
	SearchAnswerByHomeworkID(ctx context.Context, homeworkID int32) ([]*Answer,error)
	SearchAnswerByStudentID(ctx context.Context, studentID int32) ([]*Answer,error)
}

type AnswerRepositoryImpl struct{
	DB *gorm.DB
}

func(repo *AnswerRepositoryImpl) AddAnswer(ctx context.Context,answer Answer) (Answer, error){
	if err := repo.DB.Create(&answer).Error;nil != err {
		return Answer{},err
	}
	return answer,nil
}

func(repo *AnswerRepositoryImpl) DeleteAnswer(ctx context.Context,answerID int32) error{
	if err := repo.DB.Delete(&Answer{}, answerID).Error; nil != err {
		return err
	}
	return nil
}

func(repo *AnswerRepositoryImpl) SearchAnswer(ctx context.Context,answerID int32) (Answer, error){
	var answer Answer
	result := repo.DB.First(&answer, answerID)
	if nil != result.Error {
		return Answer{}, result.Error
	}
	return answer, result.Error
}

func(repo *AnswerRepositoryImpl) UpdateAnswer(ctx context.Context,answer Answer) error{
	tmp, err := repo.SearchAnswer(ctx, answer.AnswerID)
	tmp.HomeworkID = answer.HomeworkID
	tmp.StudentID = answer.StudentID
	tmp.Status = answer.Status
	tmp.CommitTime = answer.CommitTime
	if err = repo.DB.Save(tmp).Error; nil != err {
		return err
	}
	return nil
}

func(repo *AnswerRepositoryImpl) SearchAnswerByStudentID(ctx context.Context,studentID int32) ([]*Answer, error){
	var answers []*Answer

	result := repo.DB.Table("answer").Where("student_id = ?", studentID)

	if err := result.Find(&answers).Error; nil != err {
		return []*Answer{}, err
	}

	return answers, nil
}

func(repo *AnswerRepositoryImpl) SearchAnswerByHomeworkID(ctx context.Context,homeworkID int32) ([]*Answer, error){
	var answers []*Answer

	result := repo.DB.Table("answer").Where("homework_id = ?", homeworkID)

	if err := result.Find(&answers).Error; nil != err {
		return []*Answer{}, err
	}

	return answers, nil
}