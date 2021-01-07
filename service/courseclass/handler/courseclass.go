package handler

import (
	pb "boxin/service/courseclass/proto/courseclass"
	repo "boxin/service/courseclass/repository"
	// "boxin/service/user/proto/user"
	"context"
	"log"
	"time"

	// "golang.org/x/crypto/openpgp/errors"

	// "github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
)

type CourseClassHandler struct {
	CourseClassRepository repo.CourseClassRepository
}

func (c *CourseClassHandler) DeleteCourseClass(ctx context.Context, req *pb.CourseID, resp *pb.EditResponse) error {
	err := c.CourseClassRepository.DeleteCourseClass(ctx, req.CourseID)
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (c *CourseClassHandler) UpdateCourseClass(ctx context.Context, req *pb.CourseClass, resp *pb.EditResponse) error {
	//timeNow := time.Unix(timestamp, 0)
	stime := time.Unix(req.StartTime, 0)
	etime := time.Unix(req.EndTime, 0)
	course := repo.CourseClass{
		CourseID:     req.CourseID,
		CourseName:   req.CourseName,
		Introduction: req.Introduction,
		TextBooks:    req.TextBooks,
		StartTime:    stime,
		EndTime:      etime,
		State: req.State,
	}
	if err := c.CourseClassRepository.UpdateCourseClass(ctx, course); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CourseClassHandler UpdateCourseClass error:", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (c *CourseClassHandler) SearchCourseClass(ctx context.Context, req *pb.CourseID, resp *pb.SearchCourseClassResponse) error {

	course, err := c.CourseClassRepository.SearchCourseClass(ctx, req.CourseID)
	// t := time.Now()
	// stime := t.Format(course.StartTime.String())
	// etime := t.Format(course.EndTime.String())
	stime := course.StartTime.Unix()
	etime := course.EndTime.Unix()
	
	*resp = pb.SearchCourseClassResponse{
		Status: 0,
		Courseclass: &pb.CourseClass{
			CourseID:     course.CourseID,
			CourseName:   course.CourseName,
			Introduction: course.Introduction,
			TextBooks:    course.TextBooks,
			StartTime:    stime,
			EndTime:      etime,
			State: course.State,
		},
	}
	return err
}

func (c *CourseClassHandler) SearchCourseClasses(ctx context.Context, req *pb.CourseIDArray, resp *pb.SearchCourseClassesResponse) error {

	var courseclasses []*pb.CourseClass
	var err error
	var course repo.CourseClass
	for i := range req.IDArray {
		course, err = c.CourseClassRepository.SearchCourseClass(ctx, req.IDArray[i])
		// t := time.Now()
		// stime := t.Format(course.StartTime.String())
		// etime := t.Format(course.EndTime.String())
		stime := course.StartTime.Unix()
		etime := course.EndTime.Unix()

		courseclasses = append(courseclasses, &pb.CourseClass{
			CourseID:     course.CourseID,
			CourseName:   course.CourseName,
			Introduction: course.Introduction,
			TextBooks:    course.TextBooks,
			StartTime:    stime,
			EndTime:      etime,
			State: course.State,
		})
	}
	*resp = pb.SearchCourseClassesResponse{
		Status:        0,
		Msg:           "Success",
		Courseclasses: courseclasses,
	}
	return err
}

func (c *CourseClassHandler) AddTake(ctx context.Context, req *pb.Take, resp *pb.EditResponse) error {
	take := repo.Take{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		Role:     req.Role,
	}
	err := c.CourseClassRepository.AddTake(ctx, take)
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (c *CourseClassHandler) DeleteTake(ctx context.Context, req *pb.UserCourse, resp *pb.EditResponse) error {
	err := c.CourseClassRepository.DeleteTake(ctx, req.UserID, req.CourseID)
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (c *CourseClassHandler) DeleteTakeByUser(ctx context.Context, req *pb.UserID, resp *pb.EditResponse) error {
	err := c.CourseClassRepository.DeleteTakeByUser(ctx, req.UserID)
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (c *CourseClassHandler) DeleteTakeByCourseClass(ctx context.Context, req *pb.CourseID, resp *pb.EditResponse) error {
	err := c.CourseClassRepository.DeleteTakeByCourseClass(ctx, req.CourseID)
	resp.Status = 0
	resp.Msg = "Success"
	return err
}

func (c *CourseClassHandler) SearchTakeByUser(ctx context.Context, req *pb.UserID, resp *pb.SearchTakeByUserResponse) error {
	courses, err := c.CourseClassRepository.SearchTakeByUser(ctx, req.UserID)

	var ans []*pb.CourseClass
	for i := range courses {
		ans = append(ans, &pb.CourseClass{
			CourseID:     courses[i].CourseID,
			CourseName:   courses[i].CourseName,
			Introduction: courses[i].Introduction,
			TextBooks:    courses[i].TextBooks,
			StartTime:    courses[i].StartTime.Unix(),
			EndTime:      courses[i].EndTime.Unix(),
			State: courses[i].State,
		})
	}

	*resp = pb.SearchTakeByUserResponse{
		Status:  0,
		Msg:     "Success",
		Courses: ans,
	}
	return err
}

const (
	ServiceName = "go.micro.client.user"
	EtcdAddr    = "localhost:2379"
)

func (c *CourseClassHandler) NewCourse(ctx context.Context, req *pb.NewCourseMessage, resp *pb.NewCourseResponse) error {
	stime := time.Unix(req.StartTime, 0)
	etime := time.Unix(req.EndTime, 0)

	courseclass := repo.CourseClass{
		CourseName:   req.CourseName,
		Introduction: req.Introduction,
		TextBooks:    req.TextBooks,
		StartTime:    stime,
		EndTime:      etime,
		State: 		req.State,
	}

	var newCourse repo.CourseClass
	var err1 error

	newCourse, err1 = c.CourseClassRepository.NewCourse(ctx, courseclass)

	log.Println(newCourse.CourseID)
	*resp = pb.NewCourseResponse{
		Status: 0,
		Msg:    "Sucess",
		Courseclass: &pb.CourseClass{
			CourseID:     newCourse.CourseID,
			CourseName:   newCourse.CourseName,
			Introduction: newCourse.Introduction,
			TextBooks:    newCourse.TextBooks,
			StartTime:    newCourse.StartTime.Unix(),
			EndTime:      newCourse.EndTime.Unix(),
			State: newCourse.State,
		},
	}

	take := repo.Take{
		UserID:   req.UserID,
		CourseID: newCourse.CourseID,
		Role:     1,
	}

	err1 = c.CourseClassRepository.AddTake(ctx, take)
	resp.Status = 0
	resp.Msg = "Success"
	return err1
}
