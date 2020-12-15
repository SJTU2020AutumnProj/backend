package repository

import (
	"context"
	// "log"
	"time"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Check struct {
	CheckID int32 `gorm:"auto_increment;column:check_id;primary_key:true;unique;index:"`
	HomeworkID int32 `gorm:"not null;column:homework_id"`
	TeacherID int32 `gorm:"not null;column:teacher_id"`
	CheckTime time.Time `gorm:"not null;column:commit_time"`
}

func (Check) TableName() string {
	return "check"
}

type CheckRepository interface {
	AddCheck(ctx context.Context, check Check)(Check, error)
	DeleteCheck(ctx context.Context, checkID int32) error
	UpdateCheck(ctx context.Context, check Check) error
	SearchCheck(ctx context.Context, checkID int32) (Check, error)
	SearchCheckByHomeworkID(ctx context.Context, homeworkID int32) ([]*Check,error)
	SearchCheckByTeacherID(ctx context.Context, teacherID int32) ([]*Check,error)
}

type CheckRepositoryImpl struct{
	DB *gorm.DB
}

func(repo *CheckRepositoryImpl) AddCheck(ctx context.Context,check Check) (Check, error){
	if err := repo.DB.Create(&check).Error;nil != err {
		return Check{},err
	}
	return check,nil
}

func(repo *CheckRepositoryImpl) DeleteCheck(ctx context.Context,checkID int32) error{
	if err := repo.DB.Delete(&Check{}, checkID).Error; nil != err {
		return err
	}
	return nil
}

func(repo *CheckRepositoryImpl) SearchCheck(ctx context.Context,checkID int32) (Check, error){
	var check Check
	result := repo.DB.First(&check, checkID)
	if nil != result.Error {
		return Check{}, result.Error
	}
	return check, result.Error
}

func(repo *CheckRepositoryImpl) UpdateCheck(ctx context.Context,check Check) error{
	tmp, err := repo.SearchCheck(ctx, check.CheckID)
	tmp.HomeworkID = check.HomeworkID
	tmp.TeacherID = check.TeacherID
	tmp.CheckTime = check.CheckTime
	if err = repo.DB.Save(tmp).Error; nil != err {
		return err
	}
	return nil
}

func(repo *CheckRepositoryImpl) SearchCheckByTeacherID(ctx context.Context,teacherID int32) ([]*Check, error){
	var checks []*Check

	result := repo.DB.Table("check").Where("teacher_id = ?", teacherID)

	if err := result.Find(&checks).Error; nil != err {
		return []*Check{}, err
	}

	return checks, nil
}

func(repo *CheckRepositoryImpl) SearchCheckByHomeworkID(ctx context.Context,homeworkID int32) ([]*Check, error){
	var checks []*Check

	result := repo.DB.Table("check").Where("homework_id = ?", homeworkID)

	if err := result.Find(&checks).Error; nil != err {
		return []*Check{}, err
	}

	return checks, nil
}