package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"im/global"
	"im/model"
	"im/router"
)

func main() {
	initDB()
	initRouter()
}

func initRouter() {
	r := gin.Default()
	// 注册路由
	router.RegisterRouter(r)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func initDB() {
	dsn := "root:pass1234@tcp(106.14.18.18)/im?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	global.DB = db
	// 同步表结构
	syncTable()
}

func syncTable() {
	err := global.DB.AutoMigrate(
		&model.User{},
		&model.UserLoginInfo{},
	)
	if err != nil {
		panic(err)
	}
}
