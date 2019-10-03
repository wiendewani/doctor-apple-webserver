package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import (
	"../structs"
)

func (idb *InDB) GetUserDetail(c *gin.Context) {
	var (
		userDetail structs.UserDetail
		result     gin.H
	)

	Credential := c.Param("Credential")
	err := idb.DB.Where("Credential = ?", Credential).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": userDetail,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateUserDetail(c *gin.Context) {
	var (
		userDetail structs.UserDetail
		result     gin.H
	)

	Credential := c.Param("Credential")
	FullName := c.Param("FullName")

	userDetail.Credential = Credential
	userDetail.FullName = FullName

	idb.DB.Create(&userDetail)

	result = gin.H{
		"result": userDetail,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateUserDetail(c *gin.Context) {

}
