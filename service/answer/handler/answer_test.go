package handler

import (
	// handler "boxin/service/answer/handler"
	mongoDB "boxin/service/answer/mongoDB"
	answer "boxin/service/answer/proto/answer"
	repo "boxin/service/answer/repository"

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

func Test_CreateAndDeleteAnswer(t *testing.T) {
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
	collection := client.Database("jub").Collection("answer")

	a := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}
	var req answer.CreateAnswerParam
	var rsp answer.CreateAnswerResponse

	var dreq answer.AnswerID
	var drsp answer.DeleteAnswerResponse

	tf := func(status answer.CreateAnswerResponse_Status) int32 {
		So(a.CreateAnswer(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == answer.CreateAnswerResponse_SUCCESS {
			return rsp.AnswerID
		}
		return -1
	}

	tf2 := func(status answer.DeleteAnswerResponse_Status) int32 {
		So(a.DeleteAnswer(context.TODO(), &dreq, &drsp), ShouldBeNil)
		So(drsp.Status, ShouldEqual, status)
		if drsp.Status == answer.DeleteAnswerResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test Create and delete answer", t, func() {
		req.CommitTime = time.Now().Unix()
		req.Content = "哈哈哈哈"
		req.Note = "cahahah"

		id := tf(answer.CreateAnswerResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.AnswerID = rsp.AnswerID
		id2 := tf2(answer.DeleteAnswerResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, 0)
	})
}

func TestUpdateAnswer(t *testing.T) { //这里偷了个懒，直接用了AssignAnswer和DeleteAnswer
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
	collection := client.Database("jub").Collection("answer")

	h := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}
	var req answer.CreateAnswerParam
	var rsp answer.CreateAnswerResponse

	var dreq answer.AnswerID
	var drsp answer.DeleteAnswerResponse

	var ureq answer.AnswerInfo
	var ursp answer.UpdateAnswerResponse

	tf := func(status answer.UpdateAnswerResponse_Status) int32 {
		So(h.UpdateAnswer(context.TODO(), &ureq, &ursp), ShouldBeNil)
		So(ursp.Status, ShouldEqual, status)
		if ursp.Status == answer.UpdateAnswerResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test UpdateAnswer", t, func() {
		req.CommitTime = time.Now().Unix()
		req.Content = "哈哈哈哈"
		req.Note = "aaaaa"

		h.CreateAnswer(context.TODO(), &req, &rsp)

		ureq.AnswerID = rsp.AnswerID
		ureq.CommitTime = time.Now().Unix()
		ureq.Content = "哈哈哈哈哈"
		ureq.Note = "ajsjsj"
		id := tf(answer.UpdateAnswerResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.AnswerID = rsp.AnswerID
		h.DeleteAnswer(context.TODO(), &dreq, &drsp)
	})
}

func TestSearchAnswer(t *testing.T) {
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
	collection := client.Database("jub").Collection("answer")

	h := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}
	var req answer.CreateAnswerParam
	var rsp answer.CreateAnswerResponse

	var dreq answer.AnswerID
	var drsp answer.DeleteAnswerResponse

	var sreq answer.AnswerID
	var srsp answer.SearchAnswerResponse

	tf := func(status answer.SearchAnswerResponse_Status) int32 {
		So(h.SearchAnswer(context.TODO(), &sreq, &srsp), ShouldBeNil)
		So(srsp.Status, ShouldEqual, status)
		if srsp.Status == answer.SearchAnswerResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test SearchAnswer", t, func() {
		req.CommitTime = time.Now().Unix()
		req.Content = "哈哈哈哈"
		req.Note = "aaaaa"

		h.CreateAnswer(context.TODO(), &req, &rsp)

		sreq.AnswerID = rsp.AnswerID
		id := tf(answer.SearchAnswerResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.AnswerID = rsp.AnswerID
		h.DeleteAnswer(context.TODO(), &dreq, &drsp)
	})
}

func TestPostAnswerByStudent(t *testing.T){
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
	collection := client.Database("jub").Collection("answer")

	a := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}

	var req answer.PostAnswerParam
	var rsp answer.PostAnswerResponse

	var dreq answer.AnswerID
	var drsp answer.DeleteAnswerResponse

	tf := func(status answer.PostAnswerResponse_Status) int32 {
		So(a.PostAnswerByStudent(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == answer.PostAnswerResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test PostAnswerByStudet", t, func() {
		req.HomeworkID = int32(time.Now().Unix())%99773
		req.UserID = int32(time.Now().Unix())%99773
		req.CommitTime = time.Now().Unix()
		req.Content = "哈哈哈哈"
		req.Note = "aaaaa"

		id := tf(answer.PostAnswerResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.AnswerID = rsp.AnswerID
		a.DeleteAnswer(context.TODO(), &dreq, &drsp)
	})
}

func TestPostAnswerByTeacher(t *testing.T) {
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
	collection := client.Database("jub").Collection("answer")

	a := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}

	var req answer.PostAnswerParam
	var rsp answer.PostAnswerResponse

	var dreq answer.AnswerID
	var drsp answer.DeleteAnswerResponse

	tf := func(status answer.PostAnswerResponse_Status) int32 {
		So(a.PostAnswerByTeacher(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == answer.PostAnswerResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test PostAnswerByTeacher", t, func() {
		req.HomeworkID = int32(time.Now().Unix())%99773
		req.UserID = int32(time.Now().Unix())%99773
		req.CommitTime = time.Now().Unix()
		req.Content = "哈哈哈哈"
		req.Note = "aaaaa"

		id := tf(answer.PostAnswerResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.AnswerID = rsp.AnswerID
		a.DeleteAnswer(context.TODO(), &dreq, &drsp)
	})
}

func TestSearchAnswerByUserID (t*testing.T){
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
	collection := client.Database("jub").Collection("answer")

	a := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}

	var req answer.PostAnswerParam
	var rsp answer.PostAnswerResponse

	var dreq answer.AnswerID
	var drsp answer.DeleteAnswerResponse

	var sreq answer.UserID
	var srsp answer.SearchAnswerByUserIDResponse

	tf := func(status answer.SearchAnswerByUserIDResponse_Status) int32 {
		So(a.SearchAnswerByUserID(context.TODO(), &sreq, &srsp), ShouldBeNil)
		So(srsp.Status, ShouldEqual, status)
		if srsp.Status == answer.SearchAnswerByUserIDResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("TestSearchAnswerByUserID", t, func() {
		req.HomeworkID = int32(time.Now().Unix())%9283983
		req.UserID = int32(time.Now().Unix())%922743728
		req.CommitTime = time.Now().Unix()
		req.Content = "哈哈哈哈"
		req.Note = "aaaaa"

		a.PostAnswerByStudent(context.TODO(), &req, &rsp)

		sreq.UserID = req.UserID
		id := tf(answer.SearchAnswerByUserIDResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.AnswerID = rsp.AnswerID
		a.DeleteAnswer(context.TODO(), &dreq, &drsp)
	})
}

func TestSearchAnswerByHomeworkID (t *testing.T) {
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
	collection := client.Database("jub").Collection("answer")

	a := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}

	var req answer.PostAnswerParam
	var rsp answer.PostAnswerResponse

	var dreq answer.AnswerID
	var drsp answer.DeleteAnswerResponse

	var sreq answer.HomeworkID
	var srsp answer.SearchAnswerByHomeworkIDResponse

	tf := func(status answer.SearchAnswerByHomeworkIDResponse_Status) int32 {
		So(a.SearchAnswerByHomeworkID(context.TODO(), &sreq, &srsp), ShouldBeNil)
		So(srsp.Status, ShouldEqual, status)
		if srsp.Status == answer.SearchAnswerByHomeworkIDResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("TestSearchAnswerByHomeworkID", t, func() {
		req.HomeworkID = int32(time.Now().Unix())%92878464
		req.UserID = int32(time.Now().Unix())%92374626
		req.CommitTime = time.Now().Unix()
		req.Content = "哈哈哈哈"
		req.Note = "aaaaa"

		a.PostAnswerByStudent(context.TODO(), &req, &rsp)

		sreq.HomeworkID = req.HomeworkID
		id := tf(answer.SearchAnswerByHomeworkIDResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.AnswerID = rsp.AnswerID
		a.DeleteAnswer(context.TODO(), &dreq, &drsp)
	})
}

func TestSearchAnswerByUserIDAndHomeworkID(t *testing.T){
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
	collection := client.Database("jub").Collection("answer")

	a := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}

	var req answer.PostAnswerParam
	var rsp answer.PostAnswerResponse

	var dreq answer.AnswerID
	var drsp answer.DeleteAnswerResponse

	var sreq answer.UserIDAndHomeworkID
	var srsp answer.SearchAnswerByUserIDAndHomeworkIDResponse

	tf := func(status answer.SearchAnswerByUserIDAndHomeworkIDResponse_Status) int32 {
		So(a.SearchAnswerByUserIDAndHomeworkID(context.TODO(), &sreq, &srsp), ShouldBeNil)
		So(srsp.Status, ShouldEqual, status)
		if srsp.Status == answer.SearchAnswerByUserIDAndHomeworkIDResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("SearchAnswerByUserIDAndHomeworkID", t, func() {
		req.HomeworkID = int32(time.Now().Unix())%92878223
		req.UserID = int32(time.Now().Unix())%92374123
		req.CommitTime = time.Now().Unix()
		req.Content = "哈哈哈哈"
		req.Note = "aaaaa"

		a.PostAnswerByStudent(context.TODO(), &req, &rsp)

		sreq.HomeworkID = req.HomeworkID
		sreq.UserID = req.UserID
		id := tf(answer.SearchAnswerByUserIDAndHomeworkIDResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.AnswerID = rsp.AnswerID
		a.DeleteAnswer(context.TODO(), &dreq, &drsp)
	})
}

// func TestSearchAnswerByUserID(t *testing.T) {
// 	db, err := gorm.Open("mysql", MysqlUri)
// 	if nil != err {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// 设置客户端连接配置
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	// 连接到MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		fmt.Println("err")
// 	}

// 	// 检查连接
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		fmt.Println("err")
// 	}
// 	fmt.Println("Connected to MongoDB!")
// 	collection := client.Database("jub").Collection("answer")
// 	a := &AnswerHandler{AnswerRepository: &repo.AnswerRepositoryImpl{DB: db}, AnswerMongo: &mongoDB.AnswerMongoImpl{CL: collection}}

// 	var req answer.CreateAnswerParam
// 	var rsp answer.CreateAnswerResponse

// 	var dreq answer.AnswerID
// 	var drsp answer.DeleteAnswerResponse

// 	var sreq answer.UserID
// 	var srsp answer.SearchAnswerByUserIDResponse

// 	tf := func(status answer.SearchAnswerByUserIDResponse_Status) int32 {
// 		So(a.SearchAnswerByUserID(context.TODO(), &sreq, &srsp), ShouldBeNil)
// 		So(srsp.Status, ShouldEqual, status)
// 		if srsp.Status == answer.SearchAnswerByUserIDResponse_SUCCESS {
// 			return 0
// 		}
// 		return -1
// 	}

// 	Convey("Test SearchAnswer", t, func() {
// 		req.CommitTime = time.Now().Unix()
// 		req.Content = "哈哈哈哈"
// 		req.Note = "aaaaa"

// 		a.CreateAnswer(context.TODO(), &req, &rsp)

// 		sreq.UserID = rsp.UserID
// 		id := tf(answer.SearchAnswerByUserIDResponse_SUCCESS)
// 		So(id, ShouldBeGreaterThanOrEqualTo, 0)

// 		dreq.AnswerID = rsp.AnswerID
// 		a.DeleteAnswer(context.TODO(), &dreq, &drsp)
// 	})
// }
