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

type UserRepository interface {
	AddUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, userID int32) error
	UpdateUser(ctx context.Context, user User) error
	SearchUser(ctx context.Context, userID int32) (User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *UserRepositoryImpl) AddUser(ctx context.Context, user User) error {
	if err := repo.DB.Create(&user).Error; nil != err {
		return err
	}
	return nil
}

func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, userID int32) error {
	if err := repo.DB.Delete(&User{}, userID).Error; nil != err {
		return err
	}
	return nil
}

func (repo *UserRepositoryImpl) UpdateUser(ctx context.Context, user User) error {
	tmp, err := repo.SearchUser(ctx, user.UserID)
	tmp.UserType = user.UserType
	tmp.UserName = user.UserName
	tmp.Password = user.Password
	tmp.School = user.School
	tmp.ID = user.ID
	tmp.Phone = user.Phone
	tmp.Email = user.Email
	if err = repo.DB.Save(tmp).Error; nil != err {
		return err
	}
	return nil
}

func (repo *UserRepositoryImpl) SearchUser(ctx context.Context, userID int32) (User, error) {
	var user User
	result := repo.DB.First(&user, userID)
	if nil != result.Error {
		return User{}, result.Error
	}
	return user, result.Error
}
