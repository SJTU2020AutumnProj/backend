package handler

import (
	pb "boxin/service/courseclass/proto/courseclass"
	repo "boxin/service/courseclass/repository"
	"context"
	"log"
	"time"

	"golang.org/x/crypto/openpgp/errors"
)

type CourseClassHandler struct {
	CourseClassRepository repo.CourseClassRepository
}

func (c *CourseClassHandler) AddCourseClass(ctx context.Context, req *pb.CourseClass, resp *pb.EditResponse) error {

	stime, _ := time.ParseInLocation("2006-01-02 15:04:05", req.StartTime, time.Local)
	etime, _ := time.ParseInLocation("2006-01-02 15:04:05", req.EndTime, time.Local)

	courseclass := repo.CourseClass{
		CourseName:   req.CourseName,
		Introduction: req.Introduction,
		TextBooks:    req.TextBooks,
		StartTime:    stime,
		EndTime:      etime,
	}
	if err := c.CourseClassRepository.AddCourseClass(ctx, courseclass); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CourseClassHandler AddCourseClass error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (c *CourseClassHandler) DeleteCourseClass(ctx context.Context, req *pb.CourseID, resp *pb.EditResponse) error {
	if err := c.CourseClassRepository.DeleteCourseClass(ctx, req.CourseID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CourseClassHandler DeleteCourseClass error: ", err)
		return errors.ErrKeyIncorrect
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (c *CourseClassHandler) UpdateCourseClass(ctx context.Context, req *pb.CourseClass, resp *pb.EditResponse) error {
	stime, _ := time.ParseInLocation("2006-01-02 15:04:05", req.StartTime, time.Local)
	etime, _ := time.ParseInLocation("2006-01-02 15:04:05", req.EndTime, time.Local)
	course := repo.CourseClass{
		CourseID:     req.CourseID,
		CourseName:   req.CourseName,
		Introduction: req.Introduction,
		TextBooks:    req.TextBooks,
		StartTime:    stime,
		EndTime:      etime,
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
	t := time.Now()
	stime := t.Format(course.StartTime.String())
	etime := t.Format(course.EndTime.String())
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchCourseClass error:", err)
		return err
	}
	*resp = pb.SearchCourseClassResponse{
		Status: 0,
		Courseclass: &pb.CourseClass{
			CourseID:     course.CourseID,
			CourseName:   course.CourseName,
			Introduction: course.Introduction,
			TextBooks:    course.TextBooks,
			StartTime:    stime,
			EndTime:      etime,
		},
	}
	return nil
}

func (c *CourseClassHandler) SearchCourseClasses(ctx context.Context, req *pb.CourseIDArray, resp *pb.SearchCourseClassesResponse) error {

	var courseclasses []*pb.CourseClass
	for i := range req.IDArray {
		course, err := c.CourseClassRepository.SearchCourseClass(ctx, req.IDArray[i])
		t := time.Now()
		stime := t.Format(course.StartTime.String())
		etime := t.Format(course.EndTime.String())
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("CourseClassHandler SearchCourseClasses error", err)
			return err
		}
		courseclasses = append(courseclasses, &pb.CourseClass{
			CourseID:     course.CourseID,
			CourseName:   course.CourseName,
			Introduction: course.Introduction,
			TextBooks:    course.TextBooks,
			StartTime:    stime,
			EndTime:      etime,
		})
	}
	*resp = pb.SearchCourseClassesResponse{
		Status:        0,
		Msg:           "Success",
		Courseclasses: courseclasses,
	}
	return nil
}

func (c *CourseClassHandler) AddTake(ctx context.Context, req *pb.Take, resp *pb.EditResponse) error {
	take := repo.Take{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		Role:     req.Role,
	}
	if err := c.CourseClassRepository.AddTake(ctx, take); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("TakeHandler AddTake error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (c *CourseClassHandler) DeleteTake(ctx context.Context, req *pb.UserCourse, resp *pb.EditResponse) error {
	if err := c.CourseClassRepository.DeleteTake(ctx, req.UserID, req.CourseID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("TakeHandler DeleteTake error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (c *CourseClassHandler) DeleteTakeByUser(ctx context.Context, req *pb.UserID, resp *pb.EditResponse) error {
	if err := c.CourseClassRepository.DeleteTakeByUser(ctx, req.UserID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("TakeHandler DeleteTakeByUser error", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (c *CourseClassHandler) DeleteTakeByCourseClass(ctx context.Context, req *pb.CourseID, resp *pb.EditResponse) error {
	if err := c.CourseClassRepository.DeleteTakeByCourseClass(ctx, req.CourseID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("TakeHandler DeleteTake error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (c *CourseClassHandler) SearchTakeByUser(ctx context.Context, req *pb.UserID, resp *pb.SearchTakeByUserResponse) error {
	course, err := c.CourseClassRepository.SearchTakeByUser(ctx, req.UserID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchTakeByUser error: ", err)
		return err
	}

	var ans []*pb.CourseClass
	for i := range course {
		ans = append(ans, &pb.CourseClass{
			CourseID:     course[i].CourseID,
			CourseName:   course[i].CourseName,
			Introduction: course[i].Introduction,
			TextBooks:    course[i].TextBooks,
			StartTime:    course[i].StartTime.String(),
			EndTime:      course[i].EndTime.String(),
		})
	}

	*resp = pb.SearchTakeByUserResponse{
		Status:  0,
		Msg:     "Success",
		Courses: ans,
	}
	return nil
}

// func (t *CourseClassHandler) SearchUserByCourseClass(ctx context.Context, req *pb.CourseID, resp *pb.SearchTakeResponse) error {
// 	take, err := t.TakeRepository.SearchTakeByCourseClass(ctx, req.CourseID)
// 	if nil != err {
// 		resp.Status = -1
// 		resp.Msg = "Error"
// 		log.Println("Handler SearchTakeByCourseClass error: ", err)
// 		return err
// 	}

// 	*resp = pb.SearchTakeResponse{
// 		Status: 0,
// 		Take: &pb.Take{
// 			UserID:   take.UserID,
// 			CourseID: take.CourseID,
// 			role:     take.role,
// 		},
// 	}
// 	return nil
// }
