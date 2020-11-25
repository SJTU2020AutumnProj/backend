package main

import (
	courseclass "boxin/service/courseclass/proto/courseclass"
	"context"
	"log"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	ServiceName = "go.micro.client.courseclass"
	EtcdAddr    = "host.docker.internal:2379"
)

func main() {
	server := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)
	server.Init()
	courseClassService := courseclass.NewCourseClassService("go.micro.service.courseclass", server.Client())
	// log.Println("time.Now()", time.Now())

	// addCourse(courseClassService, "数学", "很难", "数学书", time.Now(), time.Now())
	// addCourse(courseClassService, "语文", "更难", "语文书", time.Now(), time.Now())
	// getCourses(courseClassService, []int32{1, 2})

	// addTake(courseClassService, 1, 1, 0)
	// addTake(courseClassService, 1, 2, 0)
	// addTake(courseClassService, 2, 1, 0)
	// addTake(courseClassService, 2, 2, 0)

	// getCourseByUser(courseClassService, 1)
	// getUserByCourse(courseClassService, 23)

	newCourse(courseClassService, 1, "新开课程测试", "测试", "书", time.Now(), time.Now())
	log.Println("完成test")
}

func getCourse(courseClassService courseclass.CourseClassService, courseID int32) (*courseclass.CourseClass, error) {
	resp, err := courseClassService.SearchCourseClass(context.Background(), &courseclass.CourseID{CourseID: courseID})
	if nil != err {
		log.Println("getCourse error:", err)
		return &courseclass.CourseClass{}, err
	}
	log.Println("Get course success:", resp)
	return resp.Courseclass, err
}

func addCourse(courseService courseclass.CourseClassService, courseName string, introduction string, textBooks string, startTime time.Time, endTime time.Time) {
	// t := time.Now()
	stime := startTime.Unix()
	etime := startTime.Unix()
	// log.Println("stime", stime)
	resp, err := courseService.AddCourseClass(
		context.Background(),
		&courseclass.CourseClass{
			CourseName:   courseName,
			Introduction: introduction,
			TextBooks:    textBooks,
			StartTime:    stime,
			EndTime:      etime,
		})
	if err != nil {
		log.Println("addCourse error:", err)
		return
	}
	log.Println("addCourse success:", resp)
}

func getCourses(courseService courseclass.CourseClassService, courseIDs []int32) ([]*courseclass.CourseClass, error) {
	resp, err := courseService.SearchCourseClasses(context.Background(), &courseclass.CourseIDArray{IDArray: courseIDs})
	if err != nil {
		log.Println("getCourses err:", err)
		return []*courseclass.CourseClass{}, err
	}
	log.Println("getCourses success: ", resp)
	return resp.Courseclasses, err
}

func addTake(courseService courseclass.CourseClassService, userID int32, courseID int32, role int32) {
	resp, err := courseService.AddTake(
		context.Background(),
		&courseclass.Take{
			UserID:   userID,
			CourseID: courseID,
			Role:     role,
		})

	if err != nil {
		log.Println("addTake error:", err)
		return
	}
	log.Println("addTake success:", resp)
}

func getCourseByUser(courseClassService courseclass.CourseClassService, userID int32) ([]*courseclass.CourseClass, error) {
	resp, err := courseClassService.SearchTakeByUser(context.Background(), &courseclass.UserID{UserID: userID})
	if nil != err {
		log.Println("getCourseByUser error:", err)
		return []*courseclass.CourseClass{}, err
	}
	log.Println("getCourseByUser success:", resp)
	return resp.Courses, err
}

func getUserByCourse(courseClassService courseclass.CourseClassService, courseID int32) ([]*courseclass.User, error) {
	resp, err := courseClassService.SearchTakeByCourse(context.Background(), &courseclass.CourseID{CourseID: courseID})
	if nil != err {
		log.Println("getUserByCourse error:", err)
		return []*courseclass.User{}, err
	}
	log.Println("getUserByCourse success:", resp)
	return resp.Users, err
}

func newCourse(courseClassService courseclass.CourseClassService, userID int32, courseName string, introduction string, textBooks string, startTime time.Time, endTime time.Time) {
	stime := startTime.Unix()
	etime := endTime.Unix()
	resp, err := courseClassService.NewCourse(
		context.Background(),
		&courseclass.NewCourseMessage{
			UserID:       userID,
			CourseName:   courseName,
			Introduction: introduction,
			TextBooks:    textBooks,
			StartTime:    stime,
			EndTime:      etime,
		})

	if nil != err {
		log.Println("NewCourse error:", err)
		return
	}
	log.Println("newCourse success", resp)
}
