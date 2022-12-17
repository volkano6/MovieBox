package controllers

import (
	"fmt"
	"net/http"
	"net/mail"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

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
	// Validation check
	// Username
	if !regexp.MustCompile(`\s`).MatchString(userRegisterObj.Username) {
		c.JSON(http.StatusInternalServerError, ErrorMessage("Username should not contain a space."))
		return
	}
	// Email
	_, err = mail.ParseAddress(userRegisterObj.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage("Invalid email."))
		return
	}
	// Password
	if len(userRegisterObj.Password) < 6 {
		c.JSON(http.StatusInternalServerError, ErrorMessage("Password cannot be shorter than 6 characters."))
		return
	}
	if len(userRegisterObj.Password) > 32 {
		c.JSON(http.StatusInternalServerError, ErrorMessage("Password cannot be longer than 32 characters."))
		return
	}
	// Date of birth
	_, err = time.Parse("2006-01-02", userRegisterObj.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage("Invalid date of birth format. Use YYYY/MM/DD."))
		return
	}
	sql := `INSERT INTO users (username, password, displayname, email, dateofbirth) VALUES (?, ?, ?, ?, ?);`
	_, err = database.DB.Exec(sql, userRegisterObj.Username, string(hashedPassword), strings.ToLower(userRegisterObj.Username), userRegisterObj.Email, userRegisterObj.DateOfBirth)
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
