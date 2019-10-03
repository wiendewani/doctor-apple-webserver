package main

import (
	"../config"
	"../controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()

	router.GET("/user/:id", inDB.GetUser)

	router.Run(":2000")
}
