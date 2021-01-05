package main

import (
	handler "boxin/service/user/handler"
	user "boxin/service/user/proto/user"
	repo "boxin/service/user/repository"
	"context"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddUser(t *testing.T) {
	// var u handler.UserHandler

	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	u := &handler.UserHandler{UserRepository: &repo.UserRepositoryImpl{DB: db}}
	var req user.User
	var rsp user.EditResponse

	tf := func(status user.EditResponse_Status) int32 {
		So(u.AddUser(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == user.EditResponse_SUCCESS {
			return req.UserID
		}
		return -1
	}

	Convey("Test Add User", t, func() {
		req.UserType = 0
		req.UserName = "unitTest"
		req.Password = "unitTest"
		req.School = "SJTU"
		req.Id = "110"
		req.Phone = "54749110"
		req.Email = "110@sjtu.edu.cn"
		id := tf(user.EditResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		defer func() { So(db.Delete(&repo.User{}, req.UserID).Error, ShouldBeNil) }()
	})

}
