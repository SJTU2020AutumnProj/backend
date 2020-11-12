package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//UserInfo 用户信息
// type UserInfo struct {
// 	ID     uint
// 	Name   string
// 	Gender string
// 	Hobby  string
// }

//创建人：lsh
//创建时间：2020.11.10

var DB *gorm.DB

var err error

func InitORM() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	DB = db
	DB.AutoMigrate(&User{}) //自动迁移
	DB.SingularTable(true)  // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	// c1 := CourseClass{1, "语文", "很难", "数学书", time.Now(), time.Now()}
	// db.Create(&c1)

	fmt.Println("InitORM")

	// var users []User
	// db.Find(&users)
	// fmt.Println("ceshi ", users)

	// var c = new(Course)
	// fmt.Println(db.First(c))

	// var courses []Course
	// db.Find(&courses)
	// fmt.Println(courses)

	// 自动迁移
	// db.AutoMigrate(&UserInfo{})

	// u1 := UserInfo{1, "枯藤", "男", "篮球"}
	// u2 := UserInfo{2, "topgoer.com", "女", "足球"}
	// // 创建记录
	// db.Create(&u1)
	// db.Create(&u2)
	// // 查询
	// var u = new(UserInfo)
	// db.First(u)
	// fmt.Printf("%#v\n", u)
	// var uu UserInfo
	// db.Find(&uu, "hobby=?", "足球")
	// fmt.Printf("%#v\n", uu)
	// // 更新
	// db.Model(&u).Update("hobby", "双色球")
	// // 删除
	// db.Delete(&u)
}
