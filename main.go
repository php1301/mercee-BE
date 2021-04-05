package main

import (
	"fmt"
	"mercee-be/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func main(){
	r := gin.Default()
	middleware.SetupLogOutput()
}