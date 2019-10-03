package controllers

package controllers

import (
"../config"
"github.com/gin-gonic/gin"
)

func UserRouter() {
	db := config.DBInit()
	inDB := &InDB{DB: db}

	router := gin.Default()

	router.GET("/user/:id", inDB.GetUser)
	router.GET("/user", inDB.GetUsers)
	router.POST("/user", inDB.CreateUser)
	router.PUT("/user", inDB.UpdateUser)
	router.DELETE("/user/:id", inDB.DeleteUser)
}

