package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/C305DatabaseProject/database-project/backend/controllers"
	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth(c *gin.Context) {
	// Get auth cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.Next()
		return
	}
	// Validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if token == nil {
		c.Next()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Next()
			return
		}
		// Get user from DB
		var user controllers.User
		sql := `SELECT id, username, displayname, email, dateofbirth, avatar, bio, location, social_twitter, social_instagram, type
			FROM users WHERE id = ?;`
		database.DB.QueryRow(sql, claims["sub"]).Scan(
			&user.ID, &user.Username, &user.Displayname, &user.Email, &user.DateOfBirth, &user.Avatar,
			&user.Bio, &user.Location, &user.Twitter, &user.Instagram, &user.Type)
		if user.ID == 0 {
			c.Next()
			return
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.Next()
		return
	}
}
