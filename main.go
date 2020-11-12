package main

import (
	"backend/database"
	"time"
)

func main() {
	database.InitORM()

	// testuser := database.GenerateUserModel(1, 0, "testuser", "123456", "sjtu", "518021", "188212", "aa@sjtu")
	// err := database.AddUser(testuser)
	// fmt.Println(err)

	// tests, err := database.GetUserByID(2)
	// fmt.Println("tests", tests)
	// fmt.Println(err)

	testCourse := database.GenerateCourseClassModel(1, "数学", "很难", "数学书", time.Now(), time.Now())
	database.AddCourseClass(testCourse)
}
