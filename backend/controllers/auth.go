package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Hompage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "ok",
	})
}

type UserRegister struct {
	Username    string `json:"user_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}
type UserLogin struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var userRegisterObj UserRegister
	err := c.ShouldBindJSON(&userRegisterObj)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage(fmt.Sprintf("Invalid JSON body. %s", err)))
		return
	}
	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterObj.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(fmt.Sprintf("Error on hashing password: %s", err)))
		return
	}
	sql := `INSERT INTO users (username, password, displayname, email, dateofbirth) VALUES (?, ?, ?, ?, ?);`
	_, err = database.DB.Exec(sql, userRegisterObj.Username, string(hashedPassword), userRegisterObj.Username, userRegisterObj.Email, userRegisterObj.DateOfBirth)
	// Error on SQL
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Everything is ok
	c.JSON(http.StatusOK, OkMessage(userRegisterObj))
}

func Login(c *gin.Context) {
	var userLoginObj UserLogin
	err := c.ShouldBindJSON(&userLoginObj)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage(fmt.Sprintf("Invalid JSON body: %s", err)))
		return
	}
	// Check if username exists
	var id string
	var username string
	var password string
	sql := `SELECT id, username, password FROM users WHERE username = ?;`
	database.DB.QueryRow(sql, userLoginObj.Username).Scan(&id, &username, &password)
	if username == "" {
		// User not found
		c.JSON(http.StatusNotFound, ErrorMessage("Invalid credentials."))
		return
	}
	// If it exists, check password
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(userLoginObj.Password))
	if err != nil {
		// Password incorrect
		c.JSON(http.StatusNotFound, ErrorMessage("Invalid credentials."))
		return
	}
	// If password is correct, generate and return JWT token
	token, err := createTokenJWT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(fmt.Sprintf("Could not generate JWT token: %s", err)))
		return
	}
	c.JSON(http.StatusOK, OkMessage(token))
}

func Logout(c *gin.Context) {
	//Change JWT as ""(nil)
}

func createTokenJWT(id string) (string, error) {
	// 1 - Start creating a new token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(time.Hour * 24 * 14) // 14 Days
	// 2 - Edit token details
	token.Claims = &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		Subject:   id,
	}
	// 3 - Sign token with secret key in env
	signed, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	// Return signed token
	return signed, nil
}
