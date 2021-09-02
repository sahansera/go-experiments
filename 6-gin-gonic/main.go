package main

import (
	"fmt"
	"gintest/config"
	"gintest/handler"
	"gintest/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	conf := config.BuildDBConfig()
	config.DB, err = gorm.Open(postgres.Open(config.DbURL(conf)), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}

	dbConn, _ := config.DB.DB()
	defer dbConn.Close()

	config.DB.AutoMigrate(&model.User{})

	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", handler.GetUsers)
		grp1.POST("user", handler.CreateUser)
	}
	r.Run()
}
