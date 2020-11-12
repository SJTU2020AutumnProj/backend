package database

import "fmt"

func (User) TableName() string {
	return "user"
}

//user map
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

//根据传入的数据生成User
func GenerateUserModel(UserID int32, UserType int32, UserName string, Password string, School string, ID string, Phone string, Email string) User {
	tmp := User{UserID, UserType, UserName, Password, School, ID, Phone, Email}
	fmt.Println("Generate user", tmp)
	return tmp
}

//AddUser 增加用户
func AddUser(u User) error {
	fmt.Println("AddUser", u.UserID)
	if err := DB.Create(&u).Error; nil != err {
		return err
	}
	fmt.Println("AddUser ok:", u.UserID)
	return nil
}

//DeleteUserByID 根据用户ID删除用户
func DeleteUserByID(ID int32) error {
	if err := DB.Delete(&User{}, ID).Error; nil != err {
		return err
	}
	return nil
}

func UpdateUser(u User) error {
	tmp, err := GetUserByID(u.UserID)
	tmp.UserType = u.UserType
	tmp.UserName = u.UserName
	tmp.Password = u.Password
	tmp.School = u.School
	tmp.ID = u.ID
	tmp.Phone = u.Phone
	tmp.Email = u.Email

	if err = DB.Save(tmp).Error; nil != err {
		return err
	}
	return nil
}

// GetUserByID 用ID获取用户
func GetUserByID(ID int32) (User, error) {
	fmt.Println("get user by ID:", ID)
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

func GetUserTypeByID(ID int32) (int32, error) {
	var user User
	result := DB.First(&user, ID)
	return user.UserType, result.Error
}
