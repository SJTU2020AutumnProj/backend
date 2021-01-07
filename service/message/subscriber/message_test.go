package subscriber

import (
	// "boxin/service/message/handler"
	homework "boxin/service/homework/proto/homework"
	mongoDB "boxin/service/message/mongoDB"
	// message "boxin/service/message/proto/message"
	repo "boxin/service/message/repository"
	// "boxin/service/message/subscriber"
	"context"
	"log"

	"testing"
	. "github.com/smartystreets/goconvey/convey"

	// 引入插件
	// "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	// 引入公共的自定义配置函数
	// "boxin/utils/tracer"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/broker"
	// "github.com/micro/go-micro/v2/broker/nats"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
	// "github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration
const (
	ServiceName = "go.micro.service.message"
	MysqlURI    = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddr    = "localhost:2379"
	MongoURI    = "mongodb://localhost:27017"
	NatsURI     = "nats://localhost:4222"
)

func TestAssigned(t *testing.T){
	//连接mysql数据库
	db, err := gorm.Open("mysql", MysqlURI)
	if nil != err {
		panic(err)
	}
	log.Println("Connected to Mysql")
	defer db.Close()

	//连接mongoDB数据库

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(MongoURI)
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	defer client.Disconnect(context.Background())
	collection := client.Database("jub").Collection("message")

	m := &MessageSub{
		MongoDB:      &mongoDB.MessageMongoImpl{CL: collection},
		MessageRepository: &repo.MessageRepositoryImpl{DB: db},
	}

	var homework *homework.AssignedHomework

	tf := func(err error) int32{
		So(m.Assigned(context.Background(),homework),ShouldBeNil)
		return 1
	}

	Convey("Test Assigned", t, func(){
		homework.HomeworkID = 1
		homework.CourseID = 1 
		homework.UserID = 1

		id:=tf(nil)
		So(id,ShouldBeGreaterThanOrEqualTo, 0)

	})
	
}