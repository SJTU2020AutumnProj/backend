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
	EtcdAddr    = "localhost:2379"
)

func main() {
	server := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)
	server.Init()
	courseClassService := courseclass.NewCourseClassService("go.micro.service.course", server.Client())
	addCourse(courseClassService, "数学", "很难", "数学书", time.Now(), time.Now())
	addCourse(courseClassService, "语文", "更难", "语文书", time.Now(), time.Now())
	getCourses(courseClassService, []int32{1, 2})
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
	t := time.Now()
	stime := t.Format(startTime.String())
	etime := t.Format(endTime.String())
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
