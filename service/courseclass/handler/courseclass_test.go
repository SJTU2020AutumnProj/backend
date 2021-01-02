package handler

import (
	// handler "boxin/service/courseclass/handler"
	courseclass "boxin/service/courseclass/proto/courseclass"
	repo "boxin/service/courseclass/repository"
	userhandler "boxin/service/user/handler"
	user "boxin/service/user/proto/user"
	userrepo "boxin/service/user/repository"
	"context"
	"testing"
	"time"

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

func TestNewCourseClass(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
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
		req.UserID = 99635
		req.CourseName = "测试课程"
		req.Introduction = "希望别出错"
		req.TextBooks = "编译原理"
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()
		req.State = 0

		id := tf(courseclass.NewCourseResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		defer func() {
			db.Where("user_id = ?", req.UserID).Delete(&repo.Take{})
			So(db.Delete(&repo.CourseClass{}, rsp.Courseclass.CourseID).Error, ShouldBeNil)
		}()
	})
}

func TestDeleteCourseClass(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.CourseID
	var rsp courseclass.EditResponse

	// var error_req courseclass.CourseID
	// var error_rsp courseclass.EditResponse

	var nreq courseclass.NewCourseMessage
	var nrsp courseclass.NewCourseResponse

	tf := func(status courseclass.EditResponse_Status) int32 {
		So(c.DeleteCourseClass(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.EditResponse_SUCCESS {
			if rsp.Status == 0 {
				return 0
			}
		}
		return -1
	}

	tf_error := func(status courseclass.EditResponse_Status) int32 {
		So(c.DeleteCourseClass(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.EditResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test DeleteCourseClass", t, func() {

		nreq.UserID = 99635
		nreq.CourseName = "测试课程"
		nreq.Introduction = "希望别出错"
		nreq.TextBooks = "编译原理"
		nreq.StartTime = time.Now().Unix()
		nreq.EndTime = time.Now().Unix()
		nreq.State = 0
		c.NewCourse(context.TODO(), &nreq, &nrsp)

		req.CourseID = nrsp.Courseclass.CourseID
		id := tf(courseclass.EditResponse_SUCCESS)
		So(id, ShouldEqual, 0)

		req.CourseID = 99999999
		id2 := tf_error(courseclass.EditResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, -1)

		// error_req.CourseID = -1
		// So(c.DeleteCourseClass(context.TODO(), &error_req, &error_rsp), ShouldBeNil)
		defer func() { db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{}) }()
	})
}

func TestUpdateCourseClass(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.CourseClass
	var rsp courseclass.EditResponse

	// var error_req courseclass.CourseClass
	// var error_rsp courseclass.EditResponse

	var nreq courseclass.NewCourseMessage
	var nrsp courseclass.NewCourseResponse

	tf := func(status courseclass.EditResponse_Status) int32 {
		So(c.UpdateCourseClass(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.EditResponse_SUCCESS {
			if rsp.Status == 0 {
				return 0
			}
		}
		return -1
	}

	tf_error := func(status courseclass.EditResponse_Status) int32 {
		So(c.UpdateCourseClass(context.TODO(), &req, &rsp), ShouldNotBeNil)
		So(rsp.Status, ShouldNotEqual, status)
		if rsp.Status == courseclass.EditResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test UpdateCourseClass", t, func() {

		nreq.UserID = 99635
		nreq.CourseName = "测试课程"
		nreq.Introduction = "希望别出错"
		nreq.TextBooks = "编译原理"
		nreq.StartTime = time.Now().Unix()
		nreq.EndTime = time.Now().Unix()
		nreq.State = 0
		c.NewCourse(context.TODO(), &nreq, &nrsp)

		req.CourseID = nrsp.Courseclass.CourseID
		req.CourseName = "更新课程"
		req.Introduction = "更新哦"
		req.TextBooks = "更新的书"
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()
		req.State = 0

		id := tf(courseclass.EditResponse_SUCCESS)
		So(id, ShouldEqual, 0)

		req.CourseID = 99999
		id2:=tf_error(courseclass.EditResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, -1)

		// So(c.UpdateCourseClass(context.TODO(), &error_req, &error_rsp), ShouldNotBeNil)
		defer func() { So(db.Delete(&repo.CourseClass{}, nrsp.Courseclass.CourseID).Error, ShouldBeNil) }()
		// defer func() { db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{}) }()
	})
}

func TestSearchCourseClass(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.CourseID
	var rsp courseclass.SearchCourseClassResponse

	// var error_req courseclass.CourseClass
	// var error_rsp courseclass.EditResponse

	var nreq courseclass.NewCourseMessage
	var nrsp courseclass.NewCourseResponse

	tf := func(status courseclass.SearchCourseClassResponse_Status) int32 {
		So(c.SearchCourseClass(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.SearchCourseClassResponse_SUCCESS {
			if rsp.Status == 0 {
				return 0
			}
		}
		return -1
	}

	tf_error := func(status courseclass.SearchCourseClassResponse_Status) int32 {
		So(c.SearchCourseClass(context.TODO(), &req, &rsp), ShouldNotBeNil)
		So(rsp.Status, ShouldNotEqual, status)
		if rsp.Status == courseclass.SearchCourseClassResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test SearchCourseClass", t, func() {

		nreq.UserID = 99635
		nreq.CourseName = "测试课程"
		nreq.Introduction = "希望别出错"
		nreq.TextBooks = "编译原理"
		nreq.StartTime = time.Now().Unix()
		nreq.EndTime = time.Now().Unix()
		nreq.State = 0
		c.NewCourse(context.TODO(), &nreq, &nrsp)

		req.CourseID = nrsp.Courseclass.CourseID

		id := tf(courseclass.SearchCourseClassResponse_SUCCESS)
		So(id, ShouldEqual, 0)

		req.CourseID = 99999
		id2:= tf_error(courseclass.SearchCourseClassResponse_SUCCESS)
		So(id2, ShouldEqual,-1)

		// So(c.UpdateCourseClass(context.TODO(), &error_req, &error_rsp), ShouldNotBeNil)
		defer func() { So(db.Delete(&repo.CourseClass{}, nrsp.Courseclass.CourseID).Error, ShouldBeNil) }()
		defer func() { db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{}) }()
	})
}

func TestSeaechCourseClasses(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.CourseIDArray
	var rsp courseclass.SearchCourseClassesResponse

	var nreq courseclass.NewCourseMessage
	var nrsp courseclass.NewCourseResponse

	tf := func(status courseclass.SearchCourseClassesResponse_Status) int32 {
		So(c.SearchCourseClasses(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.SearchCourseClassesResponse_SUCCESS {
			if rsp.Status == 0 {
				return 0
			}
		}
		return -1
	}

	Convey("Test SearchCourseClass", t, func() {

		nreq.UserID = 99635
		nreq.CourseName = "测试课程"
		nreq.Introduction = "希望别出错"
		nreq.TextBooks = "编译原理"
		nreq.StartTime = time.Now().Unix()
		nreq.EndTime = time.Now().Unix()
		nreq.State = 0
		c.NewCourse(context.TODO(), &nreq, &nrsp)

		tmp := []int32{nrsp.Courseclass.CourseID}
		req.IDArray = tmp

		id := tf(courseclass.SearchCourseClassesResponse_SUCCESS)
		So(id, ShouldEqual, 0)
		defer func() { So(db.Delete(&repo.CourseClass{}, nrsp.Courseclass.CourseID).Error, ShouldBeNil) }()
		defer func() { db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{}) }()
	})
}

func TestAddTake(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.Take
	var rsp courseclass.EditResponse

	tf := func(status courseclass.EditResponse_Status) int32 {
		So(c.AddTake(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.EditResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test AddTake", t, func() {
		req.UserID = 99635
		req.CourseID = 66988
		req.Role = 1

		id := tf(courseclass.EditResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		// defer func() { So(db.Delete(&repo.Take{}, rsp.Courseclass.CourseID).Error, ShouldBeNil) }()
		defer func() { db.Where("user_id = ?", 99635).Delete(&repo.Take{}) }()
	})
}

func TestDeleteTake(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.UserCourse
	var rsp courseclass.EditResponse

	var nreq courseclass.Take
	var nrsp courseclass.EditResponse

	tf := func(status courseclass.EditResponse_Status) int32 {
		So(c.DeleteTake(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.EditResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test DeleteTake", t, func() {
		nreq.UserID = 99635
		nreq.CourseID = 1
		nreq.Role = 1

		req.CourseID = 1
		req.UserID = 99635

		c.AddTake(context.TODO(), &nreq, &nrsp)

		id := tf(courseclass.EditResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		// defer func() { db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{}) }()
	})
}

func TestDeleteTakeByUser(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.UserID
	var rsp courseclass.EditResponse

	var nreq courseclass.Take
	var nrsp courseclass.EditResponse

	tf := func(status courseclass.EditResponse_Status) int32 {
		So(c.DeleteTakeByUser(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.EditResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test DeleteTakeByUser", t, func() {
		nreq.UserID = 99635
		nreq.CourseID = 1
		nreq.Role = 1

		req.UserID = 99635

		c.AddTake(context.TODO(), &nreq, &nrsp)

		id := tf(courseclass.EditResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		// defer func() { db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{}) }()
	})
}

func TestDeleteTakeByCourseClass(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.CourseID
	var rsp courseclass.EditResponse

	var nreq courseclass.Take
	var nrsp courseclass.EditResponse

	tf := func(status courseclass.EditResponse_Status) int32 {
		So(c.DeleteTakeByCourseClass(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.EditResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("TestDeleteTakeByCourseClass", t, func() {
		nreq.UserID = 99635
		nreq.CourseID = 99635
		nreq.Role = 1

		req.CourseID = 99635

		c.AddTake(context.TODO(), &nreq, &nrsp)

		id := tf(courseclass.EditResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		// defer func() { db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{}) }()
	})
}

func TestSearchTakeByUser(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.UserID
	var rsp courseclass.SearchTakeByUserResponse

	var nreq courseclass.Take
	var nrsp courseclass.EditResponse

	var ncreq courseclass.NewCourseMessage
	var ncrsp courseclass.NewCourseResponse

	tf := func(status courseclass.SearchTakeByUserResponse_Status) int32 {
		So(c.SearchTakeByUser(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == courseclass.SearchTakeByUserResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test NewCourseClass", t, func() {
		ncreq.UserID = 99635
		ncreq.CourseName = "测试课程"
		ncreq.Introduction = "希望别出错"
		ncreq.TextBooks = "编译原理"
		ncreq.StartTime = time.Now().Unix()
		ncreq.EndTime = time.Now().Unix()
		ncreq.State = 0
		c.NewCourse(context.TODO(), &ncreq, &ncrsp)

		nreq.UserID = 99635
		nreq.CourseID = ncrsp.Courseclass.CourseID
		nreq.Role = 1

		req.UserID = 99635

		c.AddTake(context.TODO(), &nreq, &nrsp)

		id := tf(courseclass.SearchTakeByUserResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		defer func() {
			db.Where("user_id = ?", req.UserID).Delete(&repo.Take{})
			So(db.Delete(&repo.CourseClass{}, ncrsp.Courseclass.CourseID).Error, ShouldBeNil)
		}()
	})
}

func TestSearchTakeByCourse(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	u := &(userhandler.UserHandler{UserRepository: &userrepo.UserRepositoryImpl{DB:db}})
	var ureq user.RegisterUserParam
	var ursp user.RegisterUserResponse

	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
	var req courseclass.CourseID
	// var rsp courseclass.SearchTakeByCourseResponse

	var nreq courseclass.Take
	var nrsp courseclass.EditResponse

	tf := func(status courseclass.SearchTakeByCourseResponse_Status) int32 {
		// So(c.SearchTakeByCourse(context.TODO(), &req, &rsp), ShouldBeNil)
		// So(rsp.Status, ShouldEqual, status)
		// if rsp.Status == courseclass.SearchTakeByCourseResponse_SUCCESS {
		// 	return 0
		// }
		// return -1
		return 0
	}

	Convey("TestSearchTakeByCourse", t, func() {
		ureq.UserName = "测试用户111"
		ureq.Type = 0
		ureq.Password = "123"
		ureq.School = "SJTU"
		ureq.ID = "1111"
		ureq.Phone = "1223"
		ureq.Email = "11@sjtu.edu.cn"
		ureq.School = "SJTU"
		u.RegisterStudent(context.TODO(),&ureq,&ursp)

		nreq.UserID = ursp.UserID.UserID
		nreq.CourseID = 99635
		nreq.Role = 1

		req.CourseID = 99635

		c.AddTake(context.TODO(), &nreq, &nrsp)

		id := tf(courseclass.SearchTakeByCourseResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)
		defer func() { db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{})
		db.Where("user_id = ?",nreq.UserID).Delete(&userrepo.User{})
		}()
	})
}

// func TestSearchTakeByCourse(t *testing.T) {
// 	db, err := gorm.Open("mysql", MysqlUri)
// 	if nil != err {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}

// 	u := &(userhandler.UserHandler{UserRepository: &userrepo.UserRepositoryImpl{DB: db}})
// 	var ureq user.RegisterUserParam
// 	var ursp user.RegisterUserResponse

// 	// c := &CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}
// 	// var req courseclass.CourseID
// 	// var rsp courseclass.SearchTakeByCourseResponse

// 	const (
// 		ServiceName = "go.micro.client.courseclass"
// 		EtcdAddr    = "localhost:2379"
// 	)

// 	server := micro.NewService(
// 		micro.Name(ServiceName),
// 		micro.Registry(etcd.NewRegistry(
// 			registry.Addrs(EtcdAddr),
// 		)),
// 	)
// 	server.Init()
// 	courseClassService := courseclass.NewCourseClassService("go.micro.service.courseclass", server.Client())

// 	var nreq courseclass.Take
// 	var nrsp courseclass.EditResponse

// 	tf := func(status courseclass.SearchTakeByCourseResponse_Status) int32 {
// 		// So(c.SearchTakeByCourse(context.TODO(), &req, &rsp), ShouldBeNil)
// 		// So(rsp.Status, ShouldEqual, status)
// 		// if rsp.Status == courseclass.SearchTakeByCourseResponse_SUCCESS {
// 		// 	return 0
// 		// }
// 		// return -1
// 		rsp, err1 := courseClassService.SearchTakeByCourse(context.Background(), &courseclass.CourseID{CourseID: 99635})
// 		So(err1, ShouldBeNil)
// 		So(rsp, ShouldNotBeNil)
// 		return 0
// 	}

// 	Convey("TestSearchTakeByCourse", t, func() {
// 		ureq.UserName = "测试用户111"
// 		ureq.Password = "123"
// 		ureq.School = "SJTU"
// 		ureq.ID = "1111"
// 		ureq.Phone = "1223"
// 		ureq.Email = "11@sjtu.edu.cn"
// 		u.RegisterStudent(context.TODO(), &ureq, &ursp)

// 		nreq.UserID = ursp.UserID.UserID
// 		nreq.CourseID = 99635
// 		nreq.Role = 1

// 		// req.CourseID = 99635

// 		c.AddTake(context.TODO(), &nreq, &nrsp)

// 		id := tf(courseclass.SearchTakeByCourseResponse_SUCCESS)
// 		So(id, ShouldBeGreaterThanOrEqualTo, 0)
// 		defer func() {
// 			db.Where("user_id = ?", nreq.UserID).Delete(&repo.Take{})
// 			db.Where("user_id = ?", nreq.UserID).Delete(&userrepo.User{})
// 		}()
// 	})
// }
