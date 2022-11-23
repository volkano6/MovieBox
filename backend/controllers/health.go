package controllers

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Hompage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "ok", 
	})
}

type UserRegister struct{
	USERNAME string `json:"user_name"`
	EMAIL string `json:"email"`
	DATEOFBIRTH string `json:"date_of_birth"`
	PASSWORD string `json:"password"`
}
type UserLogin struct{
	USERNAME string `json:"user_name"`
	EMAIL string `json:"email"`
	PASSWORD string `json:"password"`
}

func Register(c *gin.Context) {
	var userRegisterObj UserRegister
	if err:= c.ShouldBindJSON(&userRegisterObj); err == nil {
		fmt.Printf("user obj - %+v \n", userRegisterObj)
	}else{
		fmt.Printf("error - %+v \n", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": userRegisterObj,	
	})
}
// Can we write together Login and Logout ? 
func Login(c *gin.Context) {
	var userLoginObj UserLogin
	if err:= c.ShouldBindJSON(&userLoginObj); err == nil {
		fmt.Printf("user obj - %+v \n", userLoginObj)
	}else{
		fmt.Printf("error - %+v \n", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": userLoginObj,	
	})
}

func Logout(c *gin.Context) {
	//Change JWT as ""(nil)
}