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

func (idb *InDB) GetPenyebab(c *gin.Context) {
	var (
		penyabab  structs.Penyebab
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

func (idb *InDB) GetPenyebab(c *gin.Context) {
	var (
		penyebab  []structs.Penyebab
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

func (idb *InDB) CreatePenyebab(c *gin.Context) {
	var (
		penyebab   structs.Penyebab
		result gin.H
	)

	err := c.BindJSON(&penyebab)

	if err != nil {
		c.JSON(http.StatusBadRequest, "JSON ERROR")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(penyebab.Password), bcrypt.DefaultCost)
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

func (idb *InDB) UpdatePenyebab(c *gin.Context) {

	Username := c.Query("Username")
	Password := c.PostForm("Password")

	var (
		penyebab    structs.Penyebab
		newPenyebab structs.Penyebab
		result  gin.H
	)

	err := idb.DB.First(&penyebab, Username).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newPenyebab.Password = Password

	err = idb.DB.Model(&penyebab).Updates(newPenyebab).Error

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

func (idb *InDB) DeletePenyebab(c *gin.Context) {
	var (
		penyebab   structs.Penyebab
		result gin.H
	)

	Credential := c.Param("Credential")
	err := idb.DB.First(&penyebab, Credential).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.DB.Delete(&penyebab).Error

	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Penyebab deleted",
		}
	}

	c.JSON(http.StatusOK, result)
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
