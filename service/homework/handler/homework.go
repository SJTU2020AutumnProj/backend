package handler

import (
	"boxin/service/answer/proto/answer"
	mongoDB "boxin/service/homework/mongoDB"
	pb "boxin/service/homework/proto/homework"
	repo "boxin/service/homework/repository"
	"boxin/service/user/proto/user"

	"boxin/service/courseclass/proto/courseclass"

	// "boxin/service/homework/proto/homework"
	"context"
	"log"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

// HomeworkHandler struct
type HomeworkHandler struct {
	HomeworkRepository repo.HomeworkRepository
	HomeworkMongo      mongoDB.HomeworkMongo
	// 由go-micro封装，用于发送消息的接口，老版本叫micro.Publisher
	HomeworkAssignedPubEvent micro.Event
	HomeworkAnswerPubEvent   micro.Event
	CheckPubEvent            micro.Event
}

const (
	// HomeworkAssignedTopic topic of AssignHomework message
	HomeworkAssignedTopic = "assigned"
	// HomeworkAnswerPubTopic topic of ReleaseHomeworkAnswer message
	HomeworkAnswerPubTopic = "published"
	//CheckPubTopic topic of ReleaseCheck message
	CheckPubTopic = "checkReleased"
)

// AssignHomework assign homework
func (h *HomeworkHandler) AssignHomework(ctx context.Context, req *pb.AssignHomeworkParam, resp *pb.AssignHomeworkResponse) error {
	stime := time.Unix(req.StartTime, 0)
	etime := time.Unix(req.EndTime, 0)
	homework := repo.Homework{
		CourseID:  req.CourseID,
		UserID:    req.UserID,
		StartTime: stime,
		EndTime:   etime,
		Title:     req.Title,
		State:     req.State,
		AnswerID:  req.AnswerID,
		Score:     req.Score,
	}
	var resp_homework repo.Homework
	var err error
	if resp_homework, err = h.HomeworkRepository.AddHomework(ctx, homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler AssignHomework error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	resp.HomeworkID = resp_homework.HomeworkID

	mongo_homework := mongoDB.Homework{
		HomeworkID:  resp_homework.HomeworkID,
		Description: req.Description,
		Content:     req.Content,
		Note:        req.Note,
	}
	if err = h.HomeworkMongo.AddHomework(ctx, mongo_homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler AssignHomework error: ", err)
		return err
	}

	//初始化user_homework表
	courseClassService := courseclass.NewCourseClassService("go.micro.service.courseclass", client.DefaultClient)
	searchResult, err1 := courseClassService.SearchTakeByCourse(ctx, &courseclass.CourseID{CourseID: req.CourseID})
	if nil != err1 {
		log.Println("HomeworkHandler AssignHomework error: ", err1)
		resp.Status = -1
		resp.Msg = "Error"
		return err1
	}
	users := searchResult.Users

	for i := range users {
		userID := users[i].UserID
		h.HomeworkRepository.AddUserHomework(ctx, userID, resp.HomeworkID)
	}

	assignedHomework := &pb.AssignedHomework{
		HomeworkID:  resp_homework.HomeworkID,
		CourseID:    req.CourseID,
		UserID:      req.UserID,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Title:       req.Title,
		State:       req.State,
		AnswerID:    req.AnswerID,
		Score:       req.Score,
		Description: req.Description,
		Note:        req.Note,
	}
	if err = h.HomeworkAssignedPubEvent.Publish(ctx, assignedHomework); err != nil {
		log.Println("HomeworkHandler AssignHomework send message error ", err)
	}

	return nil
}

func (h *HomeworkHandler) DeleteHomework(ctx context.Context, req *pb.HomeworkID, resp *pb.DeleteHomeworkResponse) error {
	if err := h.HomeworkRepository.DeleteHomework(ctx, req.HomeworkID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler DeleteHomework error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"

	if err := h.HomeworkMongo.DeleteHomework(ctx, req.HomeworkID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler DeleteHomework error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (h *HomeworkHandler) UpdateHomework(ctx context.Context, req *pb.HomeworkInfo, resp *pb.UpdateHomeworkResponse) error {
	stime := time.Unix(req.StartTime, 0)
	etime := time.Unix(req.EndTime, 0)
	homework := repo.Homework{
		HomeworkID: req.HomeworkID,
		CourseID:   req.CourseID,
		UserID:     req.UserID,
		StartTime:  stime,
		EndTime:    etime,
		Title:      req.Title,
		State:      req.State,
		AnswerID:   req.AnswerID,
		Score:      req.Score,
	}

	if err := h.HomeworkRepository.UpdateHomework(ctx, homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler UpdateHomework error:", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"

	mongo_homework := mongoDB.Homework{
		HomeworkID:  req.HomeworkID,
		Description: req.Description,
		Content:     req.Content,
		Note:        req.Note,
	}

	if err := h.HomeworkMongo.UpdateHomework(ctx, mongo_homework); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler DeleteHomework error: ", err)
		return err
	}
	resp.Status = 0
	resp.Msg = "Success"
	return nil
}

func (h *HomeworkHandler) SearchHomework(ctx context.Context, req *pb.HomeworkID, resp *pb.SearchHomeworkResponse) error {
	// stime := time.Unix(req.StartTime, 0)
	// etime := time.Unix(req.EndTime, 0)
	homework, err := h.HomeworkRepository.SearchHomework(ctx, req.HomeworkID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomework error:", err)
		return err
	}

	mongo_homework, err := h.HomeworkMongo.SearchHomework(ctx, req.HomeworkID)

	*resp = pb.SearchHomeworkResponse{
		Status: 0,
		Homework: &pb.HomeworkInfo{
			HomeworkID:  homework.HomeworkID,
			CourseID:    homework.CourseID,
			UserID:      homework.UserID,
			StartTime:   homework.StartTime.Unix(),
			EndTime:     homework.EndTime.Unix(),
			Title:       homework.Title,
			State:       homework.State,
			AnswerID:    homework.AnswerID,
			Score:       homework.Score,
			Description: mongo_homework.Description,
			Content:     mongo_homework.Content,
			Note:        mongo_homework.Note,
		},
	}
	return nil
}

func (h *HomeworkHandler) GetHomeworkByTeacherID(ctx context.Context, req *pb.TeacherID, resp *pb.GetHomeworkByTeacherIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByTeacherID(ctx, req.TeacherID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomeByTeacherID error: ", err)
		return err
	}

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		homework_json, err := h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchHomework error:", err)
			return err
		}
		ans = append(ans, &pb.HomeworkInfo{
			HomeworkID:  homeworks[i].HomeworkID,
			CourseID:    homeworks[i].CourseID,
			UserID:      homeworks[i].UserID,
			StartTime:   homeworks[i].StartTime.Unix(),
			EndTime:     homeworks[i].EndTime.Unix(),
			Title:       homeworks[i].Title,
			State:       homeworks[i].State,
			AnswerID:    homeworks[i].AnswerID,
			Score:       homeworks[i].Score,
			Description: homework_json.Description,
			Content:     homework_json.Content,
			Note:        homework_json.Note,
		})
	}

	*resp = pb.GetHomeworkByTeacherIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return nil
}

func (h *HomeworkHandler) GetHomeworkByTeacherIDAndCourseID(ctx context.Context, req *pb.GetHomeworkByTeacherIDAndCourseIDParam, resp *pb.GetHomeworkByTeacherIDAndCourseIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByCourseID(ctx, req.CourseID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomeByCourseID error: ", err)
		return err
	}

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		homework_json, err := h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchHomework error:", err)
			return err
		}
		ans = append(ans, &pb.HomeworkInfo{
			HomeworkID: homeworks[i].HomeworkID,
			CourseID:   homeworks[i].CourseID,
			UserID:     homeworks[i].UserID,
			StartTime:  homeworks[i].StartTime.Unix(),
			EndTime:    homeworks[i].EndTime.Unix(),
			Title:      homeworks[i].Title,
			State:      homeworks[i].State,
			AnswerID:   homeworks[i].AnswerID,
			Score:      homeworks[i].Score,

			Description: homework_json.Description,
			Content:     homework_json.Content,
			Note:        homework_json.Note,
		})
	}

	*resp = pb.GetHomeworkByTeacherIDAndCourseIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return nil
}

func (h *HomeworkHandler) GetHomeworkByStudentID(ctx context.Context, req *pb.StudentID, resp *pb.GetHomeworkByStudentIDResponse) error {
	homeworkIDs, err := h.HomeworkRepository.SearchHomeworkIDByUserID(ctx, req.StudentID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler GetHomeworkByStudentID error: ", err)
		return err
	}

	var ans []*pb.HomeworkInfo
	for i := range homeworkIDs {
		hw, err := h.HomeworkRepository.SearchHomework(ctx, homeworkIDs[i])
		homework_json, err := h.HomeworkMongo.SearchHomework(ctx, homeworkIDs[i])
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler GetHomeworkByStudentID error:", err)
			return err
		}
		ans = append(ans, &pb.HomeworkInfo{
			HomeworkID:  hw.HomeworkID,
			CourseID:    hw.CourseID,
			UserID:      hw.UserID,
			StartTime:   hw.StartTime.Unix(),
			EndTime:     hw.EndTime.Unix(),
			Title:       hw.Title,
			State:       hw.State,
			AnswerID:    hw.AnswerID,
			Score:       hw.Score,
			Description: homework_json.Description,
			Content:     homework_json.Content,
			Note:        homework_json.Note,
		})
	}

	*resp = pb.GetHomeworkByStudentIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return nil
}

func (h *HomeworkHandler) GetHomeworkByStudentIDAndCourseID(ctx context.Context, req *pb.GetHomeworkByStudentIDAndCourseIDParam, resp *pb.GetHomeworkByStudentIDAndCourseIDResponse) error {
	//hwIDs是根据userID在user_homework中筛选的homeworkID
	hwIDs, err := h.HomeworkRepository.SearchHomeworkIDByUserID(ctx, req.StudentID)
	//homeworks是根据courseID在homework表中筛选的Homework
	homeworks, err := h.HomeworkRepository.SearchHomeworkByCourseID(ctx, req.CourseID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler GetHomeworkByStudentIDAndCourseID error: ", err)
		return err
	}

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		for j := range hwIDs {
			if homeworks[i].HomeworkID == hwIDs[j] {
				homework_json, err := h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
				if nil != err {
					resp.Status = -1
					resp.Msg = "Error"
					log.Println("Handler SearchHomework error:", err)
					return err
				}

				ans = append(ans, &pb.HomeworkInfo{
					HomeworkID:  homeworks[i].HomeworkID,
					CourseID:    homeworks[i].CourseID,
					UserID:      homeworks[i].UserID,
					StartTime:   homeworks[i].StartTime.Unix(),
					EndTime:     homeworks[i].EndTime.Unix(),
					Title:       homeworks[i].Title,
					State:       homeworks[i].State,
					AnswerID:    homeworks[i].AnswerID,
					Score:       homeworks[i].Score,
					Description: homework_json.Description,
					Content:     homework_json.Content,
					Note:        homework_json.Note,
				})
			}
		}
	}

	*resp = pb.GetHomeworkByStudentIDAndCourseIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return nil
}

func (h *HomeworkHandler) GetHomeworkByCourseID(ctx context.Context, req *pb.CourseID, resp *pb.GetHomeworkByCourseIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByCourseID(ctx, req.CourseID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler GetHomeByCourseID error: ", err)
		return err
	}

	var ans []*pb.HomeworkInfo
	for i := range homeworks {
		homework_json, err := h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler SearchHomework error:", err)
			return err
		}
		ans = append(ans, &pb.HomeworkInfo{
			HomeworkID:  homeworks[i].HomeworkID,
			CourseID:    homeworks[i].CourseID,
			UserID:      homeworks[i].UserID,
			StartTime:   homeworks[i].StartTime.Unix(),
			EndTime:     homeworks[i].EndTime.Unix(),
			Title:       homeworks[i].Title,
			State:       homeworks[i].State,
			AnswerID:    homeworks[i].AnswerID,
			Score:       homeworks[i].Score,
			Description: homework_json.Description,
			Content:     homework_json.Content,
			Note:        homework_json.Note,
		})
	}

	*resp = pb.GetHomeworkByCourseIDResponse{
		Status:    0,
		Msg:       "Success",
		Homeworks: ans,
	}
	return nil
}

func (h *HomeworkHandler) GetUserByHomeworkID(ctx context.Context, req *pb.HomeworkID, resp *pb.GetUserByHomeworkIDResponse) error {
	userService := user.NewUserService("go.micro.service.user", client.DefaultClient)
	var userInfos []*pb.UserInfo

	uh, err := h.HomeworkRepository.SearchUserHomeworkByHomeworkID(ctx, req.HomeworkID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler GetUserByHomeworkID error:", err)
		return err
	}

	for i := range uh {
		// var userID *user.UserID
		// userID.UserID = uh[i].UserID
		u, err := userService.SearchUser(ctx, &user.UserID{UserID: uh[i].UserID})
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler GetUserByHomeworkID error:", err)
			return err
		}
		// var userInfo *pb.UserInfo
		// userInfo.UserID = users.User.UserID
		// userInfo.UserType = users.User.UserType
		// userInfo.UserName = users.User.UserName
		// userInfo.Password = users.User.Password
		// userInfo.School = users.User.School
		// userInfo.ID = users.User.ID
		// userInfo.Phone = users.User.Phone
		// userInfo.Email = users.User.Email
		// userInfo.Name = users.User.Name
		// userInfo.HomeworkID = uh[i].HomeworkID
		// userInfo.AnswerID = uh[i].AnswerID
		// userInfo.CheckID = uh[i].CheckID
		// userInfo.State = uh[i].CheckID

		// userInfos = append(userInfos,userInfo)

		userInfo := &pb.UserInfo{
			UserID:     u.User.UserID,
			UserType:   u.User.UserType,
			UserName:   u.User.UserName,
			Password:   u.User.Password,
			School:     u.User.School,
			ID:         u.User.ID,
			Phone:      u.User.Phone,
			Email:      u.User.Email,
			Name:       u.User.Name,
			HomeworkID: uh[i].HomeworkID,
			AnswerID:   uh[i].AnswerID,
			CheckID:    uh[i].CheckID,
			State:      uh[i].State,
		}

		userInfos = append(userInfos, userInfo)
	}
	*resp = pb.GetUserByHomeworkIDResponse{
		Status:   0,
		Msg:      "Success",
		UserInfo: userInfos,
	}
	return nil
}

func (h *HomeworkHandler) GetHomeworkByCourseIDAndUserID(ctx context.Context, req *pb.CourseIDAndUserID, resp *pb.GetHomeworkByCourseIDAndUserIDResponse) error {
	homeworks, err := h.HomeworkRepository.SearchHomeworkByCourseID(ctx, req.CourseID)
	var ans []*pb.HomeworkAndUserInfo

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler GetHomeworkByCourseIDAndUserID error: ", err)
		return err
	}

	for i := range homeworks {
		uh, err := h.HomeworkRepository.SearchUserHomeworkByHomeworkID(ctx, homeworks[i].HomeworkID)
		if nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("Handler GetHomeworkByCourseIDAndUserID error: ", err)
			return err
		}
		for j := range uh {
			if uh[j].UserID == req.UserID {
				homework_json, err := h.HomeworkMongo.SearchHomework(ctx, homeworks[i].HomeworkID)
				if nil != err {
					resp.Status = -1
					resp.Msg = "Error"
					log.Println("Handler GetHomeworkByCourseIDAndUserID error:", err)
					return err
				}
				ans = append(ans, &pb.HomeworkAndUserInfo{
					HomeworkID:        homeworks[i].HomeworkID,
					CourseID:          homeworks[i].CourseID,
					TeacherID:         homeworks[i].UserID,
					StartTime:         homeworks[i].StartTime.Unix(),
					EndTime:           homeworks[i].EndTime.Unix(),
					Title:             homeworks[i].Title,
					State:             homeworks[i].State,
					AnswerID:          homeworks[i].AnswerID,
					Score:             homeworks[i].Score,
					Description:       homework_json.Description,
					Content:           homework_json.Content,
					Note:              homework_json.Note,
					StudentID:         uh[j].UserID,
					CheckID:           uh[j].CheckID,
					UserHomeworkState: uh[j].State,
				})
			}
		}
	}

	*resp = pb.GetHomeworkByCourseIDAndUserIDResponse{
		Status:              0,
		Msg:                 "Success",
		HomeworkAndUserInfo: ans,
	}
	return nil
}

func (h *HomeworkHandler) GetUserHomework(ctx context.Context, req *pb.GetUserHomeworkParam, resp *pb.GetUserHomeworkResponse) error {
	uh, err := h.HomeworkRepository.SearchUserHomework(ctx, req.UserID, req.HomeworkID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler GetUserHomework error: ", err)
		return err
	}

	if nil == uh {
		*resp = pb.GetUserHomeworkResponse{
			Status: 0,
			UserHomework: &pb.UserHomework{
				UserID:     req.UserID,
				HomeworkID: req.HomeworkID,
				AnswerID:   -1,
				CheckID:    -1,
				State:      0,
			},
		}
	}

	*resp = pb.GetUserHomeworkResponse{
		Status: 0,
		UserHomework: &pb.UserHomework{
			UserID:     req.UserID,
			HomeworkID: req.HomeworkID,
			AnswerID:   uh.AnswerID,
			CheckID:    uh.CheckID,
			State:      uh.State,
		},
	}
	return nil
}

//老师上传作业答案
func (h *HomeworkHandler) PostHomeworkAnswer(ctx context.Context, req *pb.PostParam, resp *pb.PostHomeworkAnswerResponse) error {

	answerService := answer.NewAnswerService("go.micro.service.answer", client.DefaultClient)

	aw, err := answerService.PostAnswerByTeacher(ctx, &answer.PostAnswerParam{HomeworkID: req.HomeworkID, UserID: req.UserID, CommitTime: req.CommitTime, Content: req.Content, Note: req.Note})

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler PostHomeworkAnswer error: ", err)
		return err
	}

	resp.AnswerID = aw.AnswerID

	err1 := h.HomeworkRepository.PostHomeworkAnswer(ctx, req.HomeworkID, aw.AnswerID)
	if nil != err1 {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler PostHomeworkAnswer error: ", err1)
		return err1
	}
	return nil
}

//老师公布作业答案
func (h *HomeworkHandler) ReleaseHomeworkAnswer(ctx context.Context, req *pb.ReleaseParam, resp *pb.ReleaseHomeworkAnswerResponse) error {
	if err := h.HomeworkRepository.ReleaseHomeworkAnswer(ctx, req.HomeworkID); nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("HomeworkHandler ReleaseHomeworkAnswer error:", err)
		return err
	}
	*resp = pb.ReleaseHomeworkAnswerResponse{
		Status: 0,
		Msg:    "Success",
	}
	homework, searchErr := h.HomeworkRepository.SearchHomework(ctx, req.HomeworkID)
	if nil != searchErr {
		log.Println("HomeworkHandler ReleaseHomeworkAnswer error ", searchErr)
		return nil
	}
	homeworkAnswerPub := &pb.HomeworkAnswerPub{
		HomeworkID: homework.HomeworkID,
		AnswerID:   homework.AnswerID,
		TeacherID:  req.TeacherID,
		CourseID:   homework.CourseID,
		Title:      homework.Title,
		PubTime:    req.PubTime,
	}
	if err := h.HomeworkAnswerPubEvent.Publish(ctx, homeworkAnswerPub); err != nil {
		log.Println("HomeworkHandler ReleaseHomeworkAnswer send message error ", err)
	}

	return nil
}

func (h *HomeworkHandler) StudentSearchHomework(ctx context.Context, req *pb.StudentSearchHomeworkParam, resp *pb.StudentSearchHomeworkResponse) error {
	homework, err := h.HomeworkRepository.SearchHomework(ctx, req.HomeworkID)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomework error:", err)
		return err
	}

	mongo_homework, err := h.HomeworkMongo.SearchHomework(ctx, req.HomeworkID)

	*resp = pb.StudentSearchHomeworkResponse{
		Status: 0,
		Homework: &pb.HomeworkInfo{
			HomeworkID:  homework.HomeworkID,
			CourseID:    homework.CourseID,
			UserID:      homework.UserID,
			StartTime:   homework.StartTime.Unix(),
			EndTime:     homework.EndTime.Unix(),
			Title:       homework.Title,
			State:       homework.State,
			AnswerID:    homework.AnswerID,
			Score:       homework.Score,
			Description: mongo_homework.Description,
			Content:     mongo_homework.Content,
			Note:        mongo_homework.Note,
		},
	}

	err = h.HomeworkRepository.UpdateUserHomeworkState(ctx, req.UserID, req.HomeworkID, 1)
	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler SearchHomework error:", err)
		return err
	}
	return nil
}

//老师公布批改情况,即修改user_homework表中的state为4
func (h *HomeworkHandler) ReleaseCheck(ctx context.Context, req *pb.ReleaseCheckParam, resp *pb.ReleaseCheckResponse) error {
	userIDs, err := h.HomeworkRepository.SearchUserIDByHomeworkID(ctx, req.HomeworkID)

	if nil != err {
		resp.Status = -1
		resp.Msg = "Error"
		log.Println("Handler ReleaseCheck error:", err)
		return err
	}

	for i := range userIDs {
		if err := h.HomeworkRepository.UpdateUserHomeworkState(ctx, userIDs[i], req.HomeworkID, 4); nil != err {
			resp.Status = -1
			resp.Msg = "Error"
			log.Println("HomeworkHandler ReleaseCheck error:", err)
			return err
		}
	}
	*resp = pb.ReleaseCheckResponse{
		Status: 0,
		Msg:    "Success",
	}
	homework, searchErr := h.HomeworkRepository.SearchHomework(ctx, req.HomeworkID)
	if nil != searchErr {
		log.Println("HomeworkHandler ReleaseCheck error ", searchErr)
		return nil
	}
	releasedCheck := &pb.ReleasedCheck{
		HomeworkID:  req.HomeworkID,
		TeacherID:   req.TeacherID,
		StudentID:   userIDs,
		CourseID:    homework.CourseID,
		ReleaseTime: req.ReleaseTime,
		Title:       homework.Title,
	}
	if err := h.CheckPubEvent.Publish(ctx, releasedCheck); err != nil {
		log.Println("HomeworkHandler ReleaseCheck send message error ", err)
	}
	return nil
}
