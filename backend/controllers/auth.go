package controllers

import (
	"fmt"
	"net/http"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Hompage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "ok",
	})
}

type UserRegister struct {
	Username    string `json:"user_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}
type UserLogin struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var userRegisterObj UserRegister
	err := c.ShouldBindJSON(&userRegisterObj)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid JSON body.",
		})
		return
	}
	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterObj.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error on hashing password: %s", err),
		})
		return
	}
	sql := `INSERT INTO users (username, password, displayname, email, dateofbirth) VALUES (?, ?, ?, ?, ?);`
	_, err = database.DB.Exec(sql, userRegisterObj.Username, string(hashedPassword), userRegisterObj.Username, userRegisterObj.Email, userRegisterObj.DateOfBirth)
	// Error on SQL
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
		})
		return
	}
	// Everything is ok
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   userRegisterObj,
	})
}

func Login(c *gin.Context) {
	var userLoginObj UserLogin
	err := c.ShouldBindJSON(&userLoginObj)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid JSON body.",
		})
		return
	}
	// Check if username exists
	var username string
	var password string
	sql := `SELECT username, password FROM users WHERE username = ?;`
	database.DB.QueryRow(sql, userLoginObj.Username).Scan(&username, &password)
	if username == "" {
		// User not found
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Invalid credentials.",
		})
		return
	}
	// If it exists, check password
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(userLoginObj.Password))
	if err != nil {
		// Password incorrect
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Invalid credentials.",
		})
		return
	}
	// If password is correct, return JWT token
	// Generate JWT token
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "jwt token.",
	})
}

func Logout(c *gin.Context) {
	//Change JWT as ""(nil)
}
