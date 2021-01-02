package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
User : map `user` table into struct
*/
type User struct {
	UserID   int32  `gorm:"auto_increment;column:user_id;primary_key:true;unique;index:"`
	UserType int32  `gorm:"default:1;not null;column:user_type"`
	UserName string `gorm:"size:100;not null;column:user_name"`
	Password string `gorm:"size:100;not null;column:password"`
	School   string `gorm:"size:100;not null;column:school"`
	ID       string `gorm:"size:100;not null;column:ID"`
	Phone    string `gorm:"size:100;column:phone"`
	Email    string `gorm:"size:100;column:email"`
	Name string `gorm:"size:100;not null;column:name"`
}

/*
TableName : map table `user` to struct User
*/
func (User) TableName() string {
	return "user"
}

/*
UserRepository : define functions about table `user`
*/
type UserRepository interface {
	AddUser(ctx context.Context, user User) (User, error)
	DeleteUser(ctx context.Context, userID int32) error
	UpdateUser(ctx context.Context, user User) error
	SearchUser(ctx context.Context, userID int32) (User, error)
	SearchUserByUserName(ctx context.Context, userName string) (User, error)
	SearchUserByPhone(ctx context.Context, phone string) (User, error)
	SearchUserByEmail(ctx context.Context, email string) (User, error)
}

/*
UserRepositoryImpl : implementation to UserRepository
*/
type UserRepositoryImpl struct {
	DB *gorm.DB
}

/*
AddUser : add a tuple to table `user` with given information
*/
func (repo *UserRepositoryImpl) AddUser(ctx context.Context, user User) (User, error) {
	if err := repo.DB.Create(&user).Error; nil != err {
		return User{}, err
	}
	return user, nil
}

/*
DeleteUser : delete a user from table `user` by the given userID
*/
func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, userID int32) error {
	if err := repo.DB.Delete(&User{}, userID).Error; nil != err {
		return err
	}
	return nil
}

/*
UpdateUser : update an tuple in table `user` with the given information
*/
func (repo *UserRepositoryImpl) UpdateUser(ctx context.Context, user User) error {
	tmp, err := repo.SearchUser(ctx, user.UserID)
	tmp.UserType = user.UserType
	tmp.UserName = user.UserName
	tmp.Password = user.Password
	tmp.School = user.School
	tmp.ID = user.ID
	tmp.Phone = user.Phone
	tmp.Email = user.Email
	tmp.Name = user.Name
	if err = repo.DB.Save(tmp).Error; nil != err {
		return err
	}
	return nil
}

/*
SearchUser : search in table `user` with the given userID
*/
func (repo *UserRepositoryImpl) SearchUser(ctx context.Context, userID int32) (User, error) {
	var user User
	result := repo.DB.First(&user, userID)
	if nil != result.Error {
		return User{}, result.Error
	}
	return user, result.Error
}

/*
SearchUserByUserName : search in table `user` with the given username
*/
func (repo *UserRepositoryImpl) SearchUserByUserName(ctx context.Context, userName string) (User, error) {
	var user User
	result := repo.DB.First(&user, "user_name = ?", userName)
	return user, result.Error
}

/*
SearchUserByPhone : search in table `user` with the given phone
*/
func (repo *UserRepositoryImpl) SearchUserByPhone(ctx context.Context, phone string) (User, error) {
	var user User
	result := repo.DB.First(&user, "phone = ?", phone)
	return user, result.Error
}

/*
SearchUserByEmail : search in table `user` with the given email
*/
func (repo *UserRepositoryImpl) SearchUserByEmail(ctx context.Context, email string) (User, error) {
	var user User
	result := repo.DB.First(&user, "email = ?", email)
	return user, result.Error
}
