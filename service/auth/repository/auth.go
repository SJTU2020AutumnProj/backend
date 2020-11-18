/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-16 22:34:02
 * @LastEditors: Seven
 * @LastEditTime: 2020-11-18 09:06:16
 */
package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	UserID   int32  `gorm:"auto_increment;column:user_id;primary_key:true;unique;index:"`
	UserType int32  `gorm:"default:1;not null;column:user_type"`
	UserName string `gorm:"size:100;not null;column:user_name"`
	Password string `gorm:"size:100;not null;column:password"`
	School   string `gorm:"size:100;not null;column:school"`
	ID       string `gorm:"size:100;not null;column:ID"`
	Phone    string `gorm:"size:100;column:phone"`
	Email    string `gorm:"size:100;column:email"`
}

func (User) TableName() string {
	return "user"
}

type AuthRepository interface {
	Login(ctx context.Context, userName string, password string) (User, error)
}

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *AuthRepositoryImpl) Login(ctx context.Context, userName string, password string) (User, error) {
	var user User
	result := repo.DB.Where("user_name = ? and password = ?", userName, password)
	if nil != result.Error {
		return User{}, result.Error
	}
	if err := result.First(&user).Error; nil != err {
		return User{}, err
	}
	return user, nil
}
