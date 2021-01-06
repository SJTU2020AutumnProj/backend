package handler

import (
	pb "boxin/service/courseclass/proto/courseclass"
	repo "boxin/service/courseclass/repository"
	"boxin/service/user/proto/user"
	"context"
	"log"
	"time"

	"golang.org/x/crypto/openpgp/errors"

	"github.com/micro/go-micro/v2/client"
)

type CourseClassHandler struct {
	CourseClassRepository repo.CourseClassRepository
}

// func (c *CourseClassHandler) AddCourseClass(ctx context.Context, req *pb.CourseClass, resp *pb.EditResponse) error {

// 	// timeNow := time.Unix(timestamp, 0)

// 	stime := time.Unix(req.StartTime, 0)
// 	etime := time.Unix(req.EndTime, 0)

// 	// log.Println("repo.stime", stime)

// 	courseclass := repo.CourseClass{
// 		CourseName:   req.CourseName,
// 		Introduction: req.Introduction,
// 		TextBooks:    req.TextBooks,
// 		StartTime:    stime,
// 		EndTime:      etime,
// 	}
// 	if err := c.CourseClassRepository.AddCourseClass(ctx, courseclass); nil != err {
// 		resp.Status = -1
// 		resp.Msg = "Error"
// 		log.Println("CourseClassHandler AddCourseClass error: ", err)
// 		return err
// 	}
// 	resp.Status = 0
// 	resp.Msg = "Success"
// 	return nil
// }

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
		State:        req.State,
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
			State:        course.State,
		},
	}
	return nil
}

func (c *CourseClassHandler) SearchCourseClasses(ctx context.Context, req *pb.CourseIDArray, resp *pb.SearchCourseClassesResponse) error {

	var courseclasses []*pb.CourseClass
	for i := range req.IDArray {
		course, err := c.CourseClassRepository.SearchCourseClass(ctx, req.IDArray[i])
		// t := time.Now()
		// stime := t.Format(course.StartTime.String())
		// etime := t.Format(course.EndTime.String())
		stime := course.StartTime.Unix()
		etime := course.EndTime.Unix()
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
			State:        course.State,
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
		log.Println("CourseHandler AddTake error: ", err)
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
		log.Println("CourseHandler DeleteTake error: ", err)
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
		log.Println("CourseHandler DeleteTakeByUser error", err)
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
		log.Println("CourseHandler DeleteTake error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (c *CourseClassHandler) SearchTakeByUser(ctx context.Context, req *pb.UserID, resp *pb.SearchTakeByUserResponse) error {
	courses, err := c.CourseClassRepository.SearchTakeByUser(ctx, req.UserID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchTakeByUser error: ", err)
		return err
	}

	var ans []*pb.CourseClass
	for i := range courses {
		ans = append(ans, &pb.CourseClass{
			CourseID:     courses[i].CourseID,
			CourseName:   courses[i].CourseName,
			Introduction: courses[i].Introduction,
			TextBooks:    courses[i].TextBooks,
			StartTime:    courses[i].StartTime.Unix(),
			EndTime:      courses[i].EndTime.Unix(),
			State:        courses[i].State,
		})
	}

	*resp = pb.SearchTakeByUserResponse{
		Status:  0,
		Msg:     "Success",
		Courses: ans,
	}
	return nil
}


func (c *CourseClassHandler) SearchTakeByCourse(ctx context.Context, req *pb.CourseID, resp *pb.SearchTakeByCourseResponse) error {
	userService := user.NewUserService("go.micro.service.user", client.DefaultClient)

	userIDs, err := c.CourseClassRepository.SearchTakeByCourseClass(ctx, req.CourseID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchTakeByCourse error: ", err)
		return err
	}

	users, err1 := userService.SearchUsers(ctx, &user.UserIDArray{UserIDArray: userIDs})

	if nil != err1 {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchTakeByCourse error: ", err)
		return err
	}

	var ans []*pb.User
	for i := range users.Users {
		ans = append(ans, &pb.User{
			UserID:   users.Users[i].UserID,
			UserType: users.Users[i].UserType,
			UserName: users.Users[i].UserName,
			School:   users.Users[i].School,
			Id:       users.Users[i].ID,
			Phone:    users.Users[i].Phone,
			Email:    users.Users[i].Email,
			Name: 	  users.Users[i].Name,
		})
	}

	*resp = pb.SearchTakeByCourseResponse{
		Status: 0,
		Msg:    "Success",
		Users:  ans,
	}
	return nil
}

func (c *CourseClassHandler) SearchUserNotInCourse(ctx context.Context,req *pb.CourseID,resp *pb.SearchUserNotInCourseResponse)error{
	userService := user.NewUserService("go.micro.service",client.DefaultClient)
	getAllUserResponse,err := userService.GetAllUsers(context.Background(),&user.GetAllUsersParam{})
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchUserNotInCourse error: ", err)
		return err
	}
	allUsers := getAllUserResponse.Users

	userIDs, err := c.CourseClassRepository.SearchTakeByCourseClass(ctx, req.CourseID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchUserNotInCourse error: ", err)
		return err
	}

	//去除已经选了这门课的学生
	for i := range allUsers{
		for j:= range userIDs{
			if allUsers[i].UserID == userIDs[j]{
				allUsers = allUsers[:i+copy(allUsers[i:], allUsers[i+1:])]
			}
		}
	}
	var ans []*pb.User

	for i:= range allUsers{
		ans = append(ans, &pb.User{
			UserID:	 allUsers[i].UserID,
			UserType:   allUsers[i].UserType,
			UserName:	allUsers[i].UserName,
			Password:	allUsers[i].Password,
			School:		allUsers[i].School,
			Id:			allUsers[i].ID,
			Phone:		allUsers[i].Phone,
			Email:		allUsers[i].Email,
			Name:		allUsers[i].Name,
		})
	}

	*resp = pb.SearchUserNotInCourseResponse{
		Status: 0,
		Msg:    "Success",
		Users:  ans,
	}
	return nil
}

func (c *CourseClassHandler) NewCourse(ctx context.Context, req *pb.NewCourseMessage, resp *pb.NewCourseResponse) error {
	stime := time.Unix(req.StartTime, 0)
	etime := time.Unix(req.EndTime, 0)

	courseclass := repo.CourseClass{
		CourseName:   req.CourseName,
		Introduction: req.Introduction,
		TextBooks:    req.TextBooks,
		StartTime:    stime,
		EndTime:      etime,
		State:        req.State,
	}

	var newCourse repo.CourseClass
	var err1 error

	if newCourse, err1 = c.CourseClassRepository.NewCourse(ctx, courseclass); nil != err1 {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("CourseClassHandler AddCourseClass error: ", err1)
		return err1
	}

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
			State:        newCourse.State,
		},
	}

	take := repo.Take{
		UserID:   req.UserID,
		CourseID: newCourse.CourseID,
		Role:     1,
	}

	if err := c.CourseClassRepository.AddTake(ctx, take); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("TakeHandler NewCourse:AddTake error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil

}
