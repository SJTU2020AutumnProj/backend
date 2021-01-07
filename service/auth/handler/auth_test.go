package handler

import (
	// handler "boxin/service/auth/handler"

	auth "boxin/service/auth/proto/auth"
	repo "boxin/service/auth/repository"
	"strconv"
	"time"

	// redis "boxin/service/auth/redis"

	userhandler "boxin/service/user/handler"
	user "boxin/service/user/proto/user"
	userrepo "boxin/service/user/repository"
	redigo "github.com/garyburd/redigo/redis"

	// userhandler "boxin/service/user/handler"
	// user "boxin/service/user/proto/user"
	// userrepo "boxin/service/user/repository"
	"context"
	"testing"

	// "time"

	// "fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "github.com/smartystreets/goconvey/convey"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
)

const (
	// ServiceName = "go.micro.service.courseclass"
	// MysqlUri = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	// EtcdAddr    = "localhost:2379"

	ServiceName = "go.micro.service.auth"
	MysqlUri    = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	RedisHost   = "127.0.0.1"
	RedisPort   = 6379
	EtcdAddr    = "localhost:2379"
)

// type UserHandler struct {
// 	UserRepository userrepo.UserRepository
// }

func TestLoginAndLogOut(t *testing.T) {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	c, err := redigo.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// authHandler := &handler.AuthHandler{
	// 	AuthRepository: &repo.AuthRepositoryImpl{DB: db},
	// 	AuthRedis: &redis.AuthRedisImpl{CONN: c},
	// }

	a := &AuthHandler{AuthRepository: &repo.AuthRepositoryImpl{DB: db}}
	u := &userhandler.UserHandler{UserRepository: &userrepo.UserRepositoryImpl{DB: db}}
	var areq user.RegisterUserParam
	var arsp user.RegisterUserResponse

	var req auth.LoginParam
	var rsp auth.LoginResponse

	// var lreq auth.LogoutParam
	// var lrsp auth.LogoutResponse

	tf := func(status auth.LoginResponse_Status) int32 {
		So(a.Login(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == auth.LoginResponse_SUCCESS {
			return 1
		}
		return -1
	}

	// tf2 := func(status auth.LogoutResponse_Status) int32 {
	// 	So(a.Logout(context.TODO(), &lreq, &lrsp), ShouldBeNil)
	// 	So(lrsp.Status, ShouldEqual, status)
	// 	if lrsp.Status == auth.LogoutResponse_SUCCESS {
	// 		return 0
	// 	}
	// 	return -1
	// }

	Convey("Test Login", t, func() {
		areq.UserName = strconv.FormatInt(time.Now().Unix()%54643435,10)

		areq.Password = "111"
		areq.School = "SJTU"
		areq.ID = strconv.FormatInt(time.Now().Unix()%5464325,10)
		areq.Phone = strconv.FormatInt(time.Now().Unix()%343467,10)
		areq.Email = strconv.FormatInt(time.Now().Unix()%546455,10)
		areq.Name = strconv.FormatInt(time.Now().Unix()%5432435,10)
		
		u.RegisterAdmin(context.TODO(), &areq, &arsp)

		req.UserName = areq.UserName
		req.Password = areq.Password

		id := tf(auth.LoginResponse_SUCCESS)
		So(id, ShouldBeGreaterThanOrEqualTo, 0)

		// lreq.Token = rsp.Token
		// id2 := tf2(auth.LogoutResponse_SUCCESS)
		// So(id2, ShouldBeGreaterThanOrEqualTo, 0)
	})
}