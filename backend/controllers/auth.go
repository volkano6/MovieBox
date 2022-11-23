package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hompage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "ok",
	})
}

type UserRegister struct {
	Username    string `json:"user_name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Password    string `json:"password"`
}
type UserLogin struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var userRegisterObj UserRegister
	if err := c.ShouldBindJSON(&userRegisterObj); err == nil {
		// Valid JSON body
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userRegisterObj,
		})
		return
	} else {
		// JSON body error
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
		})
		return
	}
}

func Login(c *gin.Context) {
	var userLoginObj UserLogin
	if err := c.ShouldBindJSON(&userLoginObj); err == nil {
		// TODO: Return JWT token
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userLoginObj,
		})
		return
	} else {
		// Invalid credentials
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
		})
		return
	}
}

func Logout(c *gin.Context) {
	//Change JWT as ""(nil)
}
