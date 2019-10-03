package controllers

import (
	"../structs"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"net/http"
	"time"
)

func (idb *InDB) GetPenyakit(c *gin.Context) {
	var (
		penyakit  structs.Penyakit
		result gin.H
	)

	username := c.Param("username")
	err := idb.DB.Where("username = ?", username).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		users  []structs.User
		result gin.H
	)

	idb.DB.Find(&users)

	if len(users) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  len(users),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePenyakit(c *gin.Context) {
	var (
		penyakit   structs.Penyakit
		result gin.H
	)

	err := c.BindJSON(&penyakit)

	if err != nil {
		c.JSON(http.StatusBadRequest, "JSON ERROR")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(penyakit.Password), bcrypt.DefaultCost)
	penyakit.Password = string(hash)

	t := time.Now()
	h := sha3.New256()
	h.Write([]byte(t.Format(time.Stamp) + "/U/" + penyakit.Username))
	penyakit.Credential = hex.EncodeToString(h.Sum(nil))

	idb.DB.Create(&penyakit)

	result = gin.H{
		"result": penyakit,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePenyakit(c *gin.Context) {

	Username := c.Query("Username")
	Password := c.PostForm("Password")

	var (
		penyakit    structs.Penyakit
		newPenyakit structs.Penyakit
		result  gin.H
	)

	err := idb.DB.First(&penyakit, Username).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newPenyakit.Password = Password

	err = idb.DB.Model(&user).Updates(newUser).Error

	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "update success",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePenyakit(c *gin.Context) {
	var (
		penyakit   structs.Penyakit
		result gin.H
	)

	Credential := c.Param("Credential")
	err := idb.DB.First(&penyakit, Credential).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.DB.Delete(&penyakit).Error

	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "penyakit deleted",
		}
	}

	c.JSON(http.StatusOK, result)
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
