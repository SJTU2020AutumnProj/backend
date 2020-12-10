package handler

import (
	homework "boxin/service/homework/proto/homework"
	repo "boxin/service/homework/repository"
	mongoDB "boxin/service/homework/mongoDB"
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

	h := &HomeworkHandler{HomeworkRepository: &repo.HomeworkRepositoryImpl{DB: db},HomeworkMongo: &mongoDB.HomeworkMongoImpl{CL:collection}}
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

	tf2 := func (status homework.DeleteHomeworkResponse_Status) int32 {
		So(h.DeleteHomework(context.TODO(), &dreq,&drsp),ShouldBeNil)
		So(drsp.Status,ShouldEqual, status)
		if drsp.Status == homework.DeleteHomeworkResponse_SUCCESS {
			return 0
		}
		return -1
	}

	Convey("Test NewCourseClass", t, func() {
		req.CourseID = 99632
		req.TeacherID = 99632
		req.StartTime = time.Now().Unix()
		req.EndTime = time.Now().Unix()
		req.HomeworkJson = "测试文字2"

		id := tf(homework.AssignHomeworkResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		dreq.HomeworkID = rsp.HomeworkID
		id2 := tf2(homework.DeleteHomeworkResponse_SUCCESS)
		So(id2, ShouldBeGreaterThanOrEqualTo, 0)
		// defer func() {
		// 	db.Where("user_id = ?", req.UserID).Delete(&repo.Take{})
		// 	So(db.Delete(&repo.CourseClass{}, rsp.Courseclass.CourseID).Error, ShouldBeNil)
		// }()
	})
}

