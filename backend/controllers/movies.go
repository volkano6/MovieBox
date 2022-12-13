package controllers

import (
	"net/http"
	"fmt"
	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	// Check comments with user id
	sql := `SELECT id, title, release_date, poster, rating, length FROM movies;`
	rows, err := database.DB.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	fmt.Println(rows)
	// Return Comments details with the User ids
	var movies []Movie
	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Poster, &movie.Rating, &movie.Length)
		movies = append(movies, movie)
	}
	if movies == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("User not have any movie."))
		return
	}
	c.JSON(http.StatusOK, OkMessage(movies))
}

func GetMovie(c *gin.Context)          {}
func PostMovies(c *gin.Context)        {}
func GetMovieComments(c *gin.Context)  {}
func PostMovieComments(c *gin.Context) {}
func PostMovieRating(c *gin.Context)   {}
func PostMovieFavorite(c *gin.Context) {}
func PostMovieCast(c *gin.Context)     {}
