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
	CommitTime time.Time `gorm:"not null;column:commit_time"`
}

type UserHomework struct {
	HomeworkID int32 `gorm:"column:homework_id;primary_key:true;index:"`
	UserID int32 `gorm:"column:user_id;primary_key:true;index:"`
	AnswerID int32 `gorm:"column:answer_id"`
	CheckID int32 `gorm:"column:check_id"`
	State int32 `gorm:"column:state"`
}


func (Answer) TableName() string {
	return "answer"
}

func (UserHomework) TableName() string {
	return "user_homework"
}

type AnswerRepository interface {
	AddAnswer(ctx context.Context, answer Answer)(Answer, error)
	DeleteAnswer(ctx context.Context, answerID int32) error
	UpdateAnswer(ctx context.Context, answer Answer) error
	SearchAnswer(ctx context.Context, answerID int32) (Answer, error)
	SearchAnswerByHomeworkID(ctx context.Context, homeworkID int32) ([]*Answer,error)
	SearchAnswerByUserID(ctx context.Context, userID int32) ([]*Answer,error)
	PostAnswerByStudent(ctx context.Context, userID int32,homeworkID int32, state int32,answer Answer) (Answer,error)
	PostAnswerByTeacher(ctx context.Context, userID int32,homeworkID int32, answer Answer) (Answer,error)
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
	tmp.CommitTime = answer.CommitTime
	if err = repo.DB.Model(&tmp).Updates(tmp).Error; nil != err {
		return err
	}
	return nil
}

func (repo *AnswerRepositoryImpl) SearchAnswerByUserID(ctx context.Context,userID int32) ([]*Answer,error){
	var user_homework []*Answer

	result := repo.DB.Table("user_homework").Where("user_id = ?", userID)

	if err := result.Find(&user_homework).Error; nil != err {
		return []*Answer{}, err
	}

	var answers []*Answer

	for i := range user_homework {
		tmp := repo.DB.Table("answer").Where("answer_id = ?", user_homework[i].AnswerID)
		var answer Answer
		if err := tmp.First(&answer).Error; nil != err {
			return []*Answer{}, err
		}
		answers = append(answers, &answer)
	}
	return answers, nil
}

func (repo *AnswerRepositoryImpl) SearchAnswerByHomeworkID(ctx context.Context, homeworkID int32) ([]*Answer, error){
	var user_homework []*Answer

	result := repo.DB.Table("user_homework").Where("homework_id = ?", homeworkID)

	if err := result.Find(&user_homework).Error; nil != err {
		return []*Answer{}, err
	}

	var answers []*Answer

	for i := range user_homework {
		tmp := repo.DB.Table("answer").Where("answer_id = ?", user_homework[i].AnswerID)
		var answer Answer
		if err := tmp.First(&answer).Error; nil != err {
			return []*Answer{}, err
		}
		answers = append(answers, &answer)
	}
	return answers, nil
}

func (repo *AnswerRepositoryImpl) PostAnswerByStudent(ctx context.Context,userID int32,homeworkID int32, state int32,answer Answer) (Answer, error){
	if err := repo.DB.Create(&answer).Error;nil != err {
		return Answer{},err
	}

	um := UserHomework{
		HomeworkID:	homeworkID,
		UserID:		userID,
		AnswerID: 	answer.AnswerID,
		CheckID: 	-1,
		State:		state,
	}

	if err := repo.DB.Create(&um).Error;nil != err {
		return Answer{},err
	}
	return answer,nil
}

func (repo *AnswerRepositoryImpl) PostAnswerByTeacher(ctx context.Context,userID int32,homeworkID int32,answer Answer) (Answer, error){
	if err := repo.DB.Create(&answer).Error;nil != err {
		return Answer{},err
	}

	um := UserHomework{
		HomeworkID:	homeworkID,
		UserID:		userID,
		AnswerID: 	answer.AnswerID,
		CheckID: 	-1,
		State:		-1,
	}

	if err := repo.DB.Save(um).Error; nil != err {
		return Answer{},err
	}
	return answer,nil
}