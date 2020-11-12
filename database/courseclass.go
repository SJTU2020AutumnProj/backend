package database

import (
	"fmt"
	"time"
)

func (CourseClass) TableName() string {
	return "courseclass"
}

//database包，该包包含了数据库的ORM映射和增删改查的基本操作，类似于spring中的entity层

//course map
type CourseClass struct {
	CourseID     int32     `gorm:"auto_increment;column:course_id;primary_key:true"`
	CourseName   string    `gorm:"size:100;not null"`
	Introduction string    `gorm:"size:1000;not null"`
	TextBooks    string    `gorm:"size:1000;not null"`
	StartTime    time.Time `gorm:"not null"`
	EndTime      time.Time `gorm:"not null"`
}

//根据传入的数据生成CourseClass
func GenerateCourseClassModel(CourseID int32, CourseName string, Introduction string, TextBooks string, StartTime time.Time, EndTime time.Time) CourseClass {
	tmp := CourseClass{CourseID, CourseName, Introduction, TextBooks, StartTime, EndTime}
	fmt.Println("GenerateCourseClass", tmp)
	return tmp
}

//AddCourse 增加课程
func AddCourseClass(c CourseClass) error {
	if err := DB.Create(&c).Error; nil != err {
		return err
	}
	fmt.Println("AddCourseClass ok:", c.CourseID)
	return nil
}

//DeleteUserByID 根据用户ID删除用户
func DeleteCourseClassByID(ID int32) error {
	if err := DB.Delete(&CourseClass{}, ID).Error; nil != err {
		return err
	}
	return nil
}

func UpdateCourseClass(c CourseClass) error {
	tmp, err := GetCourseByID(c.CourseID)
	tmp.CourseName = c.CourseName
	tmp.Introduction = c.Introduction
	tmp.TextBooks = c.TextBooks
	tmp.StartTime = c.StartTime
	tmp.EndTime = c.EndTime

	if err = DB.Save(tmp).Error; nil != err {
		return err
	}
	return nil
}

// GetUserByID 用ID获取用户
func GetCourseByID(ID int32) (CourseClass, error) {
	var course CourseClass
	result := DB.First(&course, ID)
	return course, result.Error
}

func GetCourseTypeByID(ID int32) (int32, error) {
	var course CourseClass
	result := DB.First(&course, ID)
	return course.CourseID, result.Error
}
