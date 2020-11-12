package database

func (Take) TableName() string {
	return "take"
}

//take map
type Take struct {
	UserID   int32 `gorm:"column:user_id;not null;primary_key;index"`
	CourseID int32 `gorm:"column:course_id;not null;primary_key;index"`
	role     int32 `gorm:"column:role;not null;"`
}

func GetCourseByUserID(uid int32) []CourseClass {
	var tmp []CourseClass
	DB.Find(&tmp, "UserID = ?", "uid")
	return tmp
}

func GetUserByCourseID(cid int32) []User {
	var tmp []User
	DB.Find(&tmp, "CourseID = ?", "cid")
	return tmp
}
