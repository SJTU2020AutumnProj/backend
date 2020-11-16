package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	UserID   int32  `gorm:"auto_increment;column:user_id;primary_key:true;unique;index:"`
	UserType int32  `gorm:"default:1;not null"`
	UserName string `gorm:"size:100;not null"`
	Password string `gorm:"size:100;not null"`
	School   string `gorm:"size:100;not null"`
	ID       string `gorm:"size:100;not null"`
	Phone    string `gorm:"size:100"`
	Email    string `gorm:"size:100"`
}

func (User) TableName() string {
	return "user"
}

type UserRepository interface {
	AddUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, userID int32) error
	UpdateUser(ctx context.Context, user User) error
	SearchUser(ctx context.Context, userID int32) (User, error)
	GenerateUser(
		userID int32,
		userType int32,
		userName string,
		password string,
		school string,
		ID string,
		phone string,
		email string) User
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
	return user, result.Error
}

func (repo *UserRepositoryImpl) GenerateUser(
	userID int32,
	userType int32,
	userName string,
	password string,
	school string,
	ID string,
	phone string,
	email string) User {
		return User{
			userID,
			userType,
			userName,
			password,
			school,
			ID,
			phone,
			email}
}
