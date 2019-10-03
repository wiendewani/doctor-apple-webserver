package controllers

import (
	"doctor-apple-webservers/structs"
	"encoding/hex"
	"net/http"
	"time"
)

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

func (idb *InDB) GetSolusi(c *gin.Context) {
	var (
		solusi  structs.Solusi
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

func (idb *InDB) GetSolusi(c *gin.Context) {
	var (
		solusi  []structs.Solusi
		result gin.H
	)

	idb.DB.Find(&solusi)

	if len(solusi) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": solusi,
			"count":  len(solusi),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateSolusi(c *gin.Context) {
	var (
		solusi   structs.Solusi
		result gin.H
	)

	err := c.BindJSON(&solusi)

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

	idb.DB.Create(&solusi)

	result = gin.H{
		"result": solusi,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateSolusi(c *gin.Context) {

	Username := c.Query("Username")
	Password := c.PostForm("Password")

	var (
		solusi    structs.Solusi
		newSolusi structs.Solusi
		result  gin.H
	)

	err := idb.DB.First(&solusi, Username).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newPenyakit.Password = Password

	err = idb.DB.Model(&solusi).Updates(newSolusi).Error

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

func (idb *InDB) DeleteSolusi(c *gin.Context) {
	var (
		solusi   structs.Solusi
		result gin.H
	)

	Credential := c.Param("Credential")
	err := idb.DB.First(&solusi, Credential).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.DB.Delete(&solusi).Error

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
