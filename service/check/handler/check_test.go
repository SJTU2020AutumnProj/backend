package handler

import (
	// handler "boxin/service/check/handler"
	mongoDB "boxin/service/check/mongoDB"
	check "boxin/service/check/proto/check"
	repo "boxin/service/check/repository"

	// userhandler "boxin/service/user/handler"
	// user "boxin/service/user/proto/user"
	// userrepo "boxin/service/user/repository"
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

func Test_CreateAndDeleteCheck(t *testing.T) {
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
	collection := client.Database("jub").Collection("check")

	a := &CheckHandler{CheckRepository: &repo.CheckRepositoryImpl{DB: db}, CheckMongo: &mongoDB.CheckMongoImpl{CL: collection}}
	var req check.CreateCheckParam
	var rsp check.CreateCheckResponse

	var dreq check.CheckID
	var drsp check.DeleteCheckResponse

	tf := func(status check.CreateCheckResponse_Status) int32 {
		So(a.CreateCheck(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == check.CreateCheckResponse_SUCCESS {
			return rsp.CheckID
		}
		return -1
	}

	tf2 := func(status check.DeleteCheckResponse_Status) int32 {
		So(a.DeleteCheck(context.TODO(), &dreq, &drsp), ShouldBeNil)
		So(drsp.Status, ShouldEqual, status)
		if drsp.Status == check.DeleteCheckResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test Create and delete check", t, func() {
		req.TeacherID = 9373979
		req.CheckTime = time.Now().Unix()
		req.StudentID = 8918913
		req.Comment = "aaaa"
		req.Description = "bbb"
		req.HomeworkID = 21411407
		req.Score = 1000
		

		id := tf(check.CreateCheckResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.CheckID = rsp.CheckID
		id2 := tf2(check.DeleteCheckResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, 0)
	})
}

func TestUpdateCheck(t *testing.T) {
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
	collection := client.Database("jub").Collection("check")

	h := &CheckHandler{CheckRepository: &repo.CheckRepositoryImpl{DB: db}, CheckMongo: &mongoDB.CheckMongoImpl{CL: collection}}
	var req check.CreateCheckParam
	var rsp check.CreateCheckResponse

	var dreq check.CheckID
	var drsp check.DeleteCheckResponse

	var ureq check.UpdateCheckParam
	var ursp check.UpdateCheckResponse

	tf := func(status check.UpdateCheckResponse_Status) int32 {
		So(h.UpdateCheck(context.TODO(), &ureq, &ursp), ShouldBeNil)
		So(ursp.Status, ShouldEqual, status)
		if ursp.Status == check.UpdateCheckResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test UpdateCheck", t, func() {
		req.TeacherID = 9373979
		req.CheckTime = time.Now().Unix()
		req.StudentID = 8918913
		req.Comment = "aaaa"
		req.Description = "bbb"
		req.HomeworkID = 21411407
		req.Score = 1000

		h.CreateCheck(context.TODO(), &req, &rsp)

		ureq.CheckID = rsp.CheckID

		ureq.CheckTime = time.Now().Unix()
		ureq.Comment = "aaa9a"
		ureq.Description = "b9b"
		ureq.Score = 1999
		
		id := tf(check.UpdateCheckResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.CheckID = rsp.CheckID
		h.DeleteCheck(context.TODO(), &dreq, &drsp)
	})
}

func TestSearchCheckByID(t *testing.T) {
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
	collection := client.Database("jub").Collection("check")

	h := &CheckHandler{CheckRepository: &repo.CheckRepositoryImpl{DB: db}, CheckMongo: &mongoDB.CheckMongoImpl{CL: collection}}
	var req check.CreateCheckParam
	var rsp check.CreateCheckResponse

	var dreq check.CheckID
	var drsp check.DeleteCheckResponse

	var sreq check.CheckID
	var srsp check.SearchCheckByIDResponse

	tf := func(status check.SearchCheckByIDResponse_Status) int32 {
		So(h.SearchCheckByID(context.TODO(), &sreq, &srsp), ShouldBeNil)
		So(srsp.Status, ShouldEqual, status)
		if srsp.Status == check.SearchCheckByIDResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test SearchCheck", t, func() {
		req.TeacherID = 9373979
		req.CheckTime = time.Now().Unix()
		req.StudentID = 8918913
		req.Comment = "aaaa"
		req.Description = "bbb"
		req.HomeworkID = 21411407
		req.Score = 1000

		h.CreateCheck(context.TODO(), &req, &rsp)

		sreq.CheckID = rsp.CheckID
		id := tf(check.SearchCheckByIDResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.CheckID = rsp.CheckID
		h.DeleteCheck(context.TODO(), &dreq, &drsp)
	})
}