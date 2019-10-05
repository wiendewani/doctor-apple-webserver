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
	router.GET("/ph/:id", inDB.GetDisease)
	router.POST("/ph", inDB.AnalyzeDiease)

	router.Run(":2221")
}
