package main

import (
	handler "boxin/service/courseclass/handler"
	courseclass "boxin/service/courseclass/proto/courseclass"
	repo "boxin/service/courseclass/repository"
	"context"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCourseClass(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &handler.CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.NewCourseMessage
	var rsp courseclass.NewCourseResponse

	tf := func(status courseclass.NewCourseResponse_Status) int32 {
		So(c.NewCourse(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.NewCourseResponse_SUCCESS {
			return rsp.Courseclass.CourseID
		}
		return -1
	}

	Convey("Test NewCourseClass", t, func() {
		req.UserID = 1
		req.CourseName = "测试课程"
		req.Introduction = "希望别出错"
		req.TextBooks = "编译原理"
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()

		id := tf(courseclass.NewCourseResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		defer func() { So(db.Delete(&repo.CourseClass{}, rsp.Courseclass.CourseID).Error, ShouldBeNil) }()
	})
}
