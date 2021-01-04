package repository

import (
	"context"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Check struct
type Check struct {
	CheckID int32 `gorm:"auto_increment;column:check_id;primary_key:true;unique;index:"`
	CheckTime time.Time `gorm:"not null;column:commit_time"`
	Score int32 `gorm:"not null;column:score"`
}

// TableName configure table name
func (Check) TableName() string {
	return "check"
}

// CheckRepository interface
type CheckRepository interface {
	AddCheck(ctx context.Context, check Check)(Check, error)
	DeleteCheck(ctx context.Context, checkID int32) error
	UpdateCheck(ctx context.Context, check Check) error
	SearchCheckByID(ctx context.Context, checkID int32) (Check, error)
}

// CheckRepositoryImpl implementation
type CheckRepositoryImpl struct{
	DB *gorm.DB
}

// AddCheck add a check in Mysql
func(repo *CheckRepositoryImpl) AddCheck(ctx context.Context,check Check) (Check, error){
	if err := repo.DB.Create(&check).Error;nil != err {
		log.Println("CheckRepository AddCheck error ", err)
		return check,err
	}
	return check,nil
}

// DeleteCheck delete a check in Mysql by its ID
func(repo *CheckRepositoryImpl) DeleteCheck(ctx context.Context,checkID int32) error{
	if err := repo.DB.Delete(&Check{}, checkID).Error; nil != err {
		log.Println("CheckRepository DeleteCheck error ", err)
		return err
	}
	return nil
}

// SearchCheckByID search a check in Mysql by its ID
func(repo *CheckRepositoryImpl) SearchCheckByID(ctx context.Context,checkID int32) (Check, error){
	var check Check
	result := repo.DB.First(&check, checkID)
	if nil != result.Error {
		log.Println("CheckRepository SearchCheckByID error ", result.Error)
		return Check{}, result.Error
	}
	return check, result.Error
}

// UpdateCheck update a check in Mysql
func(repo *CheckRepositoryImpl) UpdateCheck(ctx context.Context,check Check) error{
	if err := repo.DB.Model(&check).Updates(check).Error; nil != err {
		log.Println("CheckRepository UpdateCheck error ", err)
		return err
	}
	return nil
}


