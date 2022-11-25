package controllers

import (
	"net/http"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID          int
	Displayname string
	Email       string
	DateOfBirth string
	Avatar      string
	Bio         string
	Location    string
	Twitter     string
	Instagram   string
	Type        int
}

type Movie struct {
	ID          int
	Title       string
	ReleaseDate string
	Poster      string
	Rating      float32
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	// Check if user with given id exists
	sql := `SELECT id, displayname, email, dateofbirth, avatar, bio, location, social_twitter, social_instagram, type
		FROM users WHERE id = ?`
	database.DB.QueryRow(sql, id).Scan(&user.ID, &user.Displayname, &user.Email, &user.DateOfBirth, &user.Avatar, &user.Bio, &user.Location, &user.Twitter, &user.Instagram, &user.Type)
	if user.ID == 0 {
		// User not found
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User not found.",
		})
		return
	}
	// If it exists, retrieve user data
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   user,
	})
}

func Watchlist(c *gin.Context) {

}

func Watched(c *gin.Context) {
	id := c.Param("id")
	// Check watched table with given user id
	sql := `SELECT id, title, release_date, poster, rating
		FROM movies LEFT JOIN user_watched
		ON movies.id = user_watched.movie_id
		WHERE user_watched.user_id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {

	}
	// Return movie details with the movie ids
	var movies []Movie
	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Poster, &movie.Rating)
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   movies,
	})
}

func Favorites(c *gin.Context) {

}

func Comments(c *gin.Context) {

}
