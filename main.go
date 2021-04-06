package main

import (
	"fmt"
	"mercee-be/config"
	"mercee-be/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	middleware.SetupLogOutput()

	db, err := gorm.Open(mysql.Open(config.DbUrl(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("Status", err)
	}
	db.AutoMigrate()
}
