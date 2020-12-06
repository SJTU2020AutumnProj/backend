package repository

import (
	"context"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CourseClass struct {
	CourseID     int32     `gorm:"auto_increment;column:course_id;primary_key:true"`
	CourseName   string    `gorm:"size:100;not null"`
	Introduction string    `gorm:"size:1000;not null"`
	TextBooks    string    `gorm:"size:1000;not null"`
	StartTime    time.Time `gorm:"not null"`
	EndTime      time.Time `gorm:"not null"`
}

type Take struct {
	UserID   int32 `gorm:"column:user_id;not null;primary_key;index"`
	CourseID int32 `gorm:"column:course_id;not null;primary_key;index"`
	Role     int32 `gorm:"column:role;not null;"`
}

func (CourseClass) TableName() string {
	return "courseclass"
}

type CourseClassRepository interface {
	AddCourseClass(ctx context.Context, courseclass CourseClass) error
	DeleteCourseClass(ctx context.Context, courseID int32) error
	UpdateCourseClass(ctx context.Context, courseclass CourseClass) error
	SearchCourseClass(ctx context.Context, courseID int32) (CourseClass, error)

	NewCourse(ctx context.Context, courseclass CourseClass) (CourseClass, error)
	// GenerateCourseClass(
	// 	courseID int32,
	// 	courseName string,
	// 	introduction string,
	// 	textBooks string,
	// 	startTime string,
	// 	endTime string) CourseClass

	AddTake(ctx context.Context, take Take) error
	DeleteTake(ctx context.Context, userID int32, courseID int32) error
	DeleteTakeByUser(ctx context.Context, userID int32) error
	DeleteTakeByCourseClass(ctx context.Context, courseID int32) error
	SearchTakeByUser(ctx context.Context, userID int32) ([]*CourseClass, error)
	SearchTakeByCourseClass(ctx context.Context, courseID int32) ([]int32, error)
	// GenerateTake(
	// 	userID int32,
	// 	courseID int32,
	// 	role int32) Take
}

func (Take) TableName() string {
	return "take"
}

type CourseClassRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *CourseClassRepositoryImpl) AddCourseClass(ctx context.Context, courseclass CourseClass) error {
	if err := repo.DB.Create(&courseclass).Error; nil != err {
		return err
	}
	return nil
}

func (repo *CourseClassRepositoryImpl) DeleteCourseClass(ctx context.Context, courseID int32) error {
	if err := repo.DB.Delete(&CourseClass{}, courseID).Error; nil != err {
		return err
	}
	return nil
}

func (repo *CourseClassRepositoryImpl) UpdateCourseClass(ctx context.Context, c CourseClass) error {
	tmp, err := repo.SearchCourseClass(ctx, c.CourseID)
	tmp.CourseName = c.CourseName
	tmp.Introduction = c.Introduction
	tmp.TextBooks = c.TextBooks
	tmp.StartTime = c.StartTime
	tmp.EndTime = c.EndTime

	if err = repo.DB.Save(tmp).Error; nil != err {
		return err
	}
	return nil
}

func (repo *CourseClassRepositoryImpl) SearchCourseClass(ctx context.Context, courseID int32) (CourseClass, error) {
	var courseclass CourseClass
	result := repo.DB.First(&courseclass, courseID)
	if result.Error != nil {
		log.Println("Repository SearchCourseClass", result.Error)
	}
	return courseclass, result.Error
}

// func (repo *CourseClassRepositoryImpl) GenerateCourseClass(
// 	courseID int32,
// 	courseName string,
// 	introduction string,
// 	textBooks string,
// 	startTime time.Time,
// 	endTime time.Time) CourseClass {
// 	return CourseClass{
// 		courseID,
// 		courseName,
// 		introduction,
// 		textBooks,
// 		startTime,
// 		endTime}
// }

func (repo *CourseClassRepositoryImpl) AddTake(ctx context.Context, take Take) error {
	if err := repo.DB.Create(&take).Error; nil != err {
		return err
	}
	return nil
}

func (repo *CourseClassRepositoryImpl) DeleteTake(ctx context.Context, userID int32, courseID int32) error {
	if err := repo.DB.Where("user_id = ?", userID).Where("course_id = ?", courseID).Delete(&Take{}).Error; nil != err {
		return err
	}
	return nil
}

func (repo *CourseClassRepositoryImpl) DeleteTakeByUser(ctx context.Context, userID int32) error {
	if err := repo.DB.Where("user_id = ?", userID).Delete(&Take{}).Error; nil != err {
		return err
	}
	return nil
}

func (repo *CourseClassRepositoryImpl) DeleteTakeByCourseClass(ctx context.Context, courseID int32) error {
	if err := repo.DB.Where("course_id = ?", courseID).Delete(&Take{}).Error; nil != err {
		return err
	}
	return nil
}

func (repo *CourseClassRepositoryImpl) SearchTakeByUser(ctx context.Context, userID int32) ([]*CourseClass, error) {
	var takes []*Take

	// result := repo.DB.Table("take").Where("user_id = ?", userID).Find(&courses)
	// if nil != result.Error {
	// 	return courses, result.Error
	// }
	result := repo.DB.Table("take").Where("user_id = ?", userID)

	if err := result.Find(&takes).Error; nil != err {
		return []*CourseClass{}, err
	}

	var courses []*CourseClass

	for i := range takes {
		tmp := repo.DB.Table("courseclass").Where("course_id = ?", takes[i].CourseID)
		var course CourseClass
		if err := tmp.First(&course).Error; nil != err {
			return []*CourseClass{}, err
		}
		courses = append(courses, &course)
	}
	return courses, nil

	// result := repo.DB.Where(map[string]interface{}{"user_id": userID}).Find(&courses)
	// var ans []*CourseClass
	// for i := range courses {
	// 	ans[i] = &courses[i]
	// }

	// if nil != result.Error {
	// 	return ans, result.Error
	// }
	// return ans, result.Error
}

//返回userID的数组
func (repo *CourseClassRepositoryImpl) SearchTakeByCourseClass(ctx context.Context, courseID int32) ([]int32, error) {
	var tmp []*Take
	var ans []int32
	result := repo.DB.Find(&tmp, "course_id = ?", courseID)

	for i := range tmp {
		ans = append(ans, tmp[i].UserID)
	}

	if nil != result.Error {
		return ans, result.Error
	}
	return ans, result.Error
}

func (repo *CourseClassRepositoryImpl) NewCourse(ctx context.Context, courseclass CourseClass) (CourseClass, error) {
	if err := repo.DB.Create(&courseclass).Error; nil != err {
		return courseclass, err
	}
	return courseclass, nil
}

// func (repo *CourseClassRepositoryImpl) GenerateTakeClass(
// 	userID int32,
// 	courseID int32,
// 	role int32,
// ) Take {
// 	return Take{
// 		userID,
// 		courseID,
// 		role}
// }
