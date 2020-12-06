package handler

import (
	// handler "boxin/service/courseclass/handler"
	user "boxin/service/user/proto/user"
	repo "boxin/service/user/repository"
	"context"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "github.com/smartystreets/goconvey/convey"
	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
)

const (
	// ServiceName = "go.micro.service.courseclass"
	MysqlUri = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	// EtcdAddr    = "localhost:2379"
)

func TestRegisterAdmin(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	u := &UserHandler{UserRepository: &repo.UserRepositoryImpl{DB: db}}
	var req user.RegisterUserParam
	var rsp user.RegisterUserResponse

	// var error_req user.RegisterUserParam
	// var error_rsp user.RegisterUserResponse

	tf := func(status user.RegisterUserResponse_Status) int32 {
		So(u.RegisterAdmin(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == user.RegisterUserResponse_SUCCESS {
			return rsp.UserID.UserID
		}
		return -1
	}

	tf_error := func(status user.RegisterUserResponse_Status) int32 {
		So(u.RegisterAdmin(context.TODO(), &req, &rsp), ShouldNotBeNil)
		So(rsp.Status, ShouldNotEqual, status)
		if rsp.Status == user.RegisterUserResponse_SUCCESS {
			return rsp.UserID.UserID
		}
		return -1
	}

	Convey("Test Register", t, func() {
		req.UserName = "测试用户1"
		req.Password = "111"
		req.School = "SJTU"
		req.ID = "111"
		req.Phone = "111"
		req.Email = "111@sjtu.edu.cn"

		id1 := tf(user.RegisterUserResponse_SUCCESS)
		So(id1, ShouldBeGreaterThanOrEqualTo, 0)

		req.UserName = "测试用户1"
		req.Password = "222"
		req.School = "SJTU"
		req.ID = "222"
		req.Phone = "222"
		req.Email = "222@sjtu.edu.cn"
		id2 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, -1)

		req.UserName = "测试用户3"
		req.Password = "333"
		req.School = "SJTU"
		req.ID = "333"
		req.Phone = "111"
		req.Email = "333@sjtu.edu.cn"
		id3 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id3, ShouldBeGreaterThanOrEqualTo, -1)

		req.UserName = "测试用户4"
		req.Password = "444"
		req.School = "SJTU"
		req.ID = "444"
		req.Phone = "444"
		req.Email = "111@sjtu.edu.cn"
		id4 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id4, ShouldBeGreaterThanOrEqualTo, -1)

		defer func() {
			So(db.Delete(&repo.User{}, id1).Error, ShouldBeNil)
		}()
	})
}

func TestRegisterTeacher(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	u := &UserHandler{UserRepository: &repo.UserRepositoryImpl{DB: db}}
	var req user.RegisterUserParam
	var rsp user.RegisterUserResponse

	// var error_req user.RegisterUserParam
	// var error_rsp user.RegisterUserResponse

	tf := func(status user.RegisterUserResponse_Status) int32 {
		So(u.RegisterTeacher(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == user.RegisterUserResponse_SUCCESS {
			return rsp.UserID.UserID
		}
		return -1
	}

	tf_error := func(status user.RegisterUserResponse_Status) int32 {
		So(u.RegisterTeacher(context.TODO(), &req, &rsp), ShouldNotBeNil)
		So(rsp.Status, ShouldNotEqual, status)
		if rsp.Status == user.RegisterUserResponse_SUCCESS {
			return rsp.UserID.UserID
		}
		return -1
	}

	Convey("Test Register", t, func() {
		req.UserName = "测试用户1"
		req.Password = "111"
		req.School = "SJTU"
		req.ID = "111"
		req.Phone = "111"
		req.Email = "111@sjtu.edu.cn"

		id1 := tf(user.RegisterUserResponse_SUCCESS)
		So(id1, ShouldBeGreaterThanOrEqualTo, 0)

		req.UserName = "测试用户1"
		req.Password = "222"
		req.School = "SJTU"
		req.ID = "222"
		req.Phone = "222"
		req.Email = "222@sjtu.edu.cn"
		id2 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, -1)

		req.UserName = "测试用户3"
		req.Password = "333"
		req.School = "SJTU"
		req.ID = "333"
		req.Phone = "111"
		req.Email = "333@sjtu.edu.cn"
		id3 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id3, ShouldBeGreaterThanOrEqualTo, -1)

		req.UserName = "测试用户4"
		req.Password = "444"
		req.School = "SJTU"
		req.ID = "444"
		req.Phone = "444"
		req.Email = "111@sjtu.edu.cn"
		id4 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id4, ShouldBeGreaterThanOrEqualTo, -1)

		defer func() {
			So(db.Delete(&repo.User{}, id1).Error, ShouldBeNil)
		}()
	})
}

func TestRegisterStudent(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	u := &UserHandler{UserRepository: &repo.UserRepositoryImpl{DB: db}}
	var req user.RegisterUserParam
	var rsp user.RegisterUserResponse

	// var error_req user.RegisterUserParam
	// var error_rsp user.RegisterUserResponse

	tf := func(status user.RegisterUserResponse_Status) int32 {
		So(u.RegisterStudent(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == user.RegisterUserResponse_SUCCESS {
			return rsp.UserID.UserID
		}
		return -1
	}

	tf_error := func(status user.RegisterUserResponse_Status) int32 {
		So(u.RegisterStudent(context.TODO(), &req, &rsp), ShouldNotBeNil)
		So(rsp.Status, ShouldNotEqual, status)
		if rsp.Status == user.RegisterUserResponse_SUCCESS {
			return rsp.UserID.UserID
		}
		return -1
	}

	Convey("Test Register", t, func() {
		req.UserName = "测试用户1"
		req.Password = "111"
		req.School = "SJTU"
		req.ID = "111"
		req.Phone = "111"
		req.Email = "111@sjtu.edu.cn"

		id1 := tf(user.RegisterUserResponse_SUCCESS)
		So(id1, ShouldBeGreaterThanOrEqualTo, 0)

		req.UserName = "测试用户1"
		req.Password = "222"
		req.School = "SJTU"
		req.ID = "222"
		req.Phone = "222"
		req.Email = "222@sjtu.edu.cn"
		id2 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, -1)

		req.UserName = "测试用户3"
		req.Password = "333"
		req.School = "SJTU"
		req.ID = "333"
		req.Phone = "111"
		req.Email = "333@sjtu.edu.cn"
		id3 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id3, ShouldBeGreaterThanOrEqualTo, -1)

		req.UserName = "测试用户4"
		req.Password = "444"
		req.School = "SJTU"
		req.ID = "444"
		req.Phone = "444"
		req.Email = "111@sjtu.edu.cn"
		id4 := tf_error(user.RegisterUserResponse_SUCCESS)
		So(id4, ShouldBeGreaterThanOrEqualTo, -1)

		defer func() {
			So(db.Delete(&repo.User{}, id1).Error, ShouldBeNil)
		}()
	})
}

func TestUpdateUser(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	u := &UserHandler{UserRepository: &repo.UserRepositoryImpl{DB: db}}
	var req user.UpdateUserParam
	var rsp user.UpdateUserResponse

	var nreq user.RegisterUserParam
	var nrsp user.RegisterUserResponse

	tf := func(status user.UpdateUserResponse_Status) int32 {
		So(u.UpdateUser(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == user.UpdateUserResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test Update", t, func() {
		nreq.UserName = "测试用户1"
		nreq.Password = "111"
		nreq.School = "SJTU"
		nreq.ID = "111"
		nreq.Phone = "111"
		nreq.Email = "111@sjtu.edu.cn"

		u.RegisterStudent(context.TODO(), &nreq, &nrsp)

		id1 := nrsp.UserID.UserID

		req.UserID = id1
		req.UserType = 1
		req.UserName = "测试用户3"
		req.School = "SJTU"
		req.ID = "333"
		req.Phone = "111"
		req.Email = "333@sjtu.edu.cn"

		id2 := tf(user.UpdateUserResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, 0)

		defer func() {
			So(db.Delete(&repo.User{}, id1).Error, ShouldBeNil)
		}()
	})
}

func TestSearchUser(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	u := &UserHandler{UserRepository: &repo.UserRepositoryImpl{DB: db}}
	var req user.UserID
	var rsp user.SearchUserResponse

	var nreq user.RegisterUserParam
	var nrsp user.RegisterUserResponse

	tf := func(status user.SearchUserResponse_Status) int32 {
		So(u.SearchUser(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == user.SearchUserResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test SearchUser", t, func() {
		nreq.UserName = "测试用户1"
		nreq.Password = "111"
		nreq.School = "SJTU"
		nreq.ID = "111"
		nreq.Phone = "111"
		nreq.Email = "111@sjtu.edu.cn"

		u.RegisterStudent(context.TODO(), &nreq, &nrsp)

		id1 := nrsp.UserID.UserID

		req.UserID = id1

		id2 := tf(user.SearchUserResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, 0)

		defer func() {
			So(db.Delete(&repo.User{}, id1).Error, ShouldBeNil)
		}()
	})
}

func TestSearchUsers(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	u := &UserHandler{UserRepository: &repo.UserRepositoryImpl{DB: db}}
	var req user.UserIDArray
	var rsp user.SearchUsersResponse

	var nreq user.RegisterUserParam
	var nrsp user.RegisterUserResponse

	tf := func(status user.SearchUserResponse_Status) int32 {
		So(u.SearchUsers(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == user.SearchUsersResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test SearchUser", t, func() {
		nreq.UserName = "测试用户1"
		nreq.Password = "111"
		nreq.School = "SJTU"
		nreq.ID = "111"
		nreq.Phone = "111"
		nreq.Email = "111@sjtu.edu.cn"

		u.RegisterStudent(context.TODO(), &nreq, &nrsp)

		id1 := nrsp.UserID.UserID

		tmp := []int32{id1}
		req.UserIDArray = tmp

		id2 := tf(user.SearchUserResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, 0)

		defer func() {
			So(db.Delete(&repo.User{}, id1).Error, ShouldBeNil)
		}()
	})
}
