package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetDisease(c *gin.Context) {
	var (
		diseases structs.Disease
		result   gin.H
	)

	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": diseases,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) AnalyzeDiease(c *gin.Context) {
	var (
		analyzer structs.Analyzer
		result   gin.H
	)

	err := c.BindJSON(&analyzer)

	if err != nil {
		c.JSON(http.StatusBadRequest, "JSON ERROR")
		return
	}

	result = gin.H{
		"result": analyzer,
	}

	c.JSON(200, result)
}
