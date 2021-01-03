package handler

import (
	mongoDB "boxin/service/homework/mongoDB"
	homework "boxin/service/homework/proto/homework"
	repo "boxin/service/homework/repository"
	"context"
	"testing"
	"time"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "github.com/smartystreets/goconvey/convey"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
)

const (
	// ServiceName = "go.micro.service.courseclass"
	MysqlUri = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	// EtcdAddr    = "localhost:2379"
)

func TestAssignAndDeleteHomework(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("err")
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("jub").Collection("homework")

	h := &HomeworkHandler{HomeworkRepository: &repo.HomeworkRepositoryImpl{DB: db}, HomeworkMongo: &mongoDB.HomeworkMongoImpl{CL: collection}}
	var req homework.AssignHomeworkParam
	var rsp homework.AssignHomeworkResponse

	var dreq homework.HomeworkID
	var drsp homework.DeleteHomeworkResponse

	tf := func(status homework.AssignHomeworkResponse_Status) int32 {
		So(h.AssignHomework(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == homework.AssignHomeworkResponse_SUCCESS {
			return rsp.HomeworkID
		}
		return -1
	}

	tf2 := func(status homework.DeleteHomeworkResponse_Status) int32 {
		So(h.DeleteHomework(context.TODO(), &dreq, &drsp), ShouldBeNil)
		So(drsp.Status, ShouldEqual, status)
		if drsp.Status == homework.DeleteHomeworkResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test Assign and delete homework", t, func() {
		req.CourseID = 99632
		req.UserID = 99632
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()
		req.Title = "cccc"
		req.State = 1
		req.Description = "测试用例2"
		req.Content = "哈哈哈哈"

		id := tf(homework.AssignHomeworkResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.HomeworkID = rsp.HomeworkID
		id2 := tf2(homework.DeleteHomeworkResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, 0)
	})
}

func TestUpdateHomework(t *testing.T) { //这里偷了个懒，直接用了AssignHomework和DeleteHomework
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("err")
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("jub").Collection("homework")

	h := &HomeworkHandler{HomeworkRepository: &repo.HomeworkRepositoryImpl{DB: db}, HomeworkMongo: &mongoDB.HomeworkMongoImpl{CL: collection}}
	var req homework.AssignHomeworkParam
	var rsp homework.AssignHomeworkResponse

	var dreq homework.HomeworkID
	var drsp homework.DeleteHomeworkResponse

	var ureq homework.HomeworkInfo
	var ursp homework.UpdateHomeworkResponse

	tf := func(status homework.UpdateHomeworkResponse_Status) int32 {
		So(h.UpdateHomework(context.TODO(), &ureq, &ursp), ShouldBeNil)
		So(ursp.Status, ShouldEqual, status)
		if ursp.Status == homework.UpdateHomeworkResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test UpdateHomework", t, func() {
		req.CourseID = 99632
		req.UserID = 99632
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()
		req.Title = "Homework"
		req.State = 1
		req.Description = "测试用例2"
		req.Content = "哈哈哈哈"

		h.AssignHomework(context.TODO(), &req, &rsp)

		ureq.HomeworkID = rsp.HomeworkID
		ureq.CourseID = 99999
		ureq.UserID = 99999
		ureq.StartTime = time.Now().Unix()
		ureq.EndTime = time.Now().Unix()
		ureq.Description = "测试用例2"
		ureq.Content = "哈哈哈哈哈"
		id := tf(homework.UpdateHomeworkResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.HomeworkID = rsp.HomeworkID
		h.DeleteHomework(context.TODO(), &dreq, &drsp)
	})
}

func TestSearchHomework(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("err")
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("jub").Collection("homework")

	h := &HomeworkHandler{HomeworkRepository: &repo.HomeworkRepositoryImpl{DB: db}, HomeworkMongo: &mongoDB.HomeworkMongoImpl{CL: collection}}
	var req homework.AssignHomeworkParam
	var rsp homework.AssignHomeworkResponse

	var dreq homework.HomeworkID
	var drsp homework.DeleteHomeworkResponse

	var sreq homework.HomeworkID
	var srsp homework.SearchHomeworkResponse

	tf := func(status homework.SearchHomeworkResponse_Status) int32 {
		So(h.SearchHomework(context.TODO(), &sreq, &srsp), ShouldBeNil)
		So(srsp.Status, ShouldEqual, status)
		if srsp.Status == homework.SearchHomeworkResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test SearchHomework", t, func() {
		req.CourseID = 99632
		req.UserID = 99632
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()
		req.Title = "Homework"
		req.State = 1
		req.Description = "测试用例2"
		req.Content = "哈哈哈哈"

		h.AssignHomework(context.TODO(), &req, &rsp)

		sreq.HomeworkID = rsp.HomeworkID
		id := tf(homework.SearchHomeworkResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.HomeworkID = rsp.HomeworkID
		h.DeleteHomework(context.TODO(), &dreq, &drsp)
	})
}

func TestSearchHomeworkByUserID(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("err")
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("jub").Collection("homework")

	h := &HomeworkHandler{HomeworkRepository: &repo.HomeworkRepositoryImpl{DB: db}, HomeworkMongo: &mongoDB.HomeworkMongoImpl{CL: collection}}
	var req homework.AssignHomeworkParam
	var rsp homework.AssignHomeworkResponse

	var dreq homework.HomeworkID
	var drsp homework.DeleteHomeworkResponse

	var sreq homework.UserID
	var srsp homework.SearchHomeworkByUserIDResponse

	tf := func(status homework.SearchHomeworkByUserIDResponse_Status) int32 {
		So(h.SearchHomeworkByUserID(context.TODO(), &sreq, &srsp), ShouldBeNil)
		So(srsp.Status, ShouldEqual, status)
		if srsp.Status == homework.SearchHomeworkByUserIDResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test SearchHomework", t, func() {
		req.CourseID = 99632
		req.UserID = 99632
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()
		req.Title = "ccc"
		req.State = 1
		req.Description = "测试用例"
		req.Content = "哈哈哈哈哈"

		h.AssignHomework(context.TODO(), &req, &rsp)

		sreq.UserID = req.UserID
		id := tf(homework.SearchHomeworkByUserIDResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.HomeworkID = rsp.HomeworkID
		h.DeleteHomework(context.TODO(), &dreq, &drsp)
	})
}

func TestSearchHomeworkByCourseID(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("err")
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("jub").Collection("homework")

	h := &HomeworkHandler{HomeworkRepository: &repo.HomeworkRepositoryImpl{DB: db}, HomeworkMongo: &mongoDB.HomeworkMongoImpl{CL: collection}}
	var req homework.AssignHomeworkParam
	var rsp homework.AssignHomeworkResponse

	var dreq homework.HomeworkID
	var drsp homework.DeleteHomeworkResponse

	var sreq homework.CourseID
	var srsp homework.SearchHomeworkByCourseIDResponse

	tf := func(status homework.SearchHomeworkByCourseIDResponse_Status) int32 {
		So(h.SearchHomeworkByCourseID(context.TODO(), &sreq, &srsp), ShouldBeNil)
		So(srsp.Status, ShouldEqual, status)
		if srsp.Status == homework.SearchHomeworkByCourseIDResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test SearchHomework", t, func() {
		req.CourseID = 99632
		req.UserID = 99632
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()
		req.Title = "ccc"
		req.State = 1
		req.Description = "测试用例"
		req.Content = "哈哈哈哈哈"

		h.AssignHomework(context.TODO(), &req, &rsp)

		sreq.CourseID = req.CourseID
		id := tf(homework.SearchHomeworkByCourseIDResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.HomeworkID = rsp.HomeworkID
		h.DeleteHomework(context.TODO(), &dreq, &drsp)
	})
}
