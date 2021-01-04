package repository

import (
	"context"
	// "log"
	"time"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Homework struct {
	HomeworkID int32 `gorm:"auto_increment;column:homework_id;primary_key:true;unique;index"`
	CourseID int32 `gorm:"not null;column:course_id"`
	UserID int32 `gorm:"not null;column:user_id"`
	StartTime time.Time `gorm:"not null;column:start_time"`
	EndTime time.Time `gorm:"not null;column:end_time"`
	Title string `gorm:"not null;column:title"`
	State int32 `gorm:"not null;column:state"`
	AnswerID int32 `gorm:"column:answer_id"`
}

type UserHomework struct {
	HomeworkID int32 `gorm:"column:homework_id;primary_key:true;index:"`
	UserID int32 `gorm:"column:user_id;primary_key:true;index:"`
	AnswerID int32 `gorm:"column:answer_id"`
	CheckID int32 `gorm:"column:check_id"`
	State int32 `gorm:"column:state"`
}

func (Homework) TableName() string {
	return "homework"
}

func (UserHomework) TableName() string {
	return "user_homework"
}

type HomeworkRepository interface {
	AddHomework(ctx context.Context,homework Homework) (Homework, error)
	DeleteHomework(ctx context.Context,homeworkID int32) error
	UpdateHomework(ctx context.Context,homework Homework) error
	SearchHomework(ctx context.Context,homeworkID int32) (Homework, error)
	SearchHomeworkByUserID(ctx context.Context,userID int32) ([]*Homework, error)
	SearchHomeworkByCourseID(ctx context.Context,courseID int32) ([]*Homework, error)
	PostHomeworkAnswer(ctx context.Context,homeworkID int32, answerID int32) error
	ReleaseHomeworkAnswer(ctx context.Context,homeworkID int32) error
}

type HomeworkRepositoryImpl struct {
	DB *gorm.DB
}

func(repo *HomeworkRepositoryImpl) AddHomework(ctx context.Context,homework Homework) (Homework, error){
	if err := repo.DB.Create(&homework).Error;nil != err {
		return Homework{},err
	}
	return homework,nil
}

func(repo *HomeworkRepositoryImpl) DeleteHomework(ctx context.Context,homeworkID int32) error{
	if err := repo.DB.Delete(&Homework{}, homeworkID).Error; nil != err {
		return err
	}
	return nil
}

func(repo *HomeworkRepositoryImpl) UpdateHomework(ctx context.Context,homework Homework) error{
	tmp, err := repo.SearchHomework(ctx, homework.HomeworkID)
	tmp.CourseID = homework.CourseID
	tmp.UserID = homework.UserID
	tmp.StartTime = homework.StartTime
	tmp.EndTime = homework.EndTime
	tmp.Title = homework.Title
	tmp.State = homework.State
	tmp.AnswerID = homework.AnswerID
	// if err = repo.DB.Save(tmp).Error; nil != err {
	// 	return err
	// }
	// return nil
	if err = repo.DB.Model(&tmp).Updates(tmp).Error; nil != err {
		return err
	}
	return nil
}

func(repo *HomeworkRepositoryImpl) SearchHomework(ctx context.Context,homeworkID int32) (Homework, error){
	var homework Homework
	result := repo.DB.First(&homework, homeworkID)
	if nil != result.Error {
		return Homework{}, result.Error
	}
	return homework, result.Error
}

func(repo *HomeworkRepositoryImpl) SearchHomeworkByUserID(ctx context.Context,userID int32) ([]*Homework, error){
	var homeworks []*Homework

	result := repo.DB.Table("homework").Where("user_id = ?", userID)

	if err := result.Find(&homeworks).Error; nil != err {
		return []*Homework{}, err
	}
	return homeworks, nil
}

func(repo *HomeworkRepositoryImpl) SearchHomeworkByCourseID(ctx context.Context,courseID int32) ([]*Homework, error){
	var homeworks []*Homework

	result := repo.DB.Table("homework").Where("course_id = ?", courseID)

	if err := result.Find(&homeworks).Error; nil != err {
		return []*Homework{}, err
	}
	return homeworks, nil
}

//这个函数仅仅把homework表中的answer_id填上
func (repo*HomeworkRepositoryImpl) PostHomeworkAnswer(ctx context.Context,homeworkID int32, answerID int32) error {
	tmp, err := repo.SearchHomework(ctx, homeworkID)
	tmp.AnswerID = answerID
	if err = repo.DB.Model(&tmp).Updates(tmp).Error; nil != err {
		return err
	}
	return nil
}

//这个函数把homework表中的state更改
func (repo*HomeworkRepositoryImpl) ReleaseHomeworkAnswer(ctx context.Context,homeworkID int32) error{
	tmp, err := repo.SearchHomework(ctx, homeworkID)
	tmp.State = 2
	if err = repo.DB.Model(&tmp).Updates(tmp).Error; nil != err {
		return err
	}
	return nil
}