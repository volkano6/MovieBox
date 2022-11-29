package controllers

import (
	_ "net/http"

	_ "github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {}
func GetMovie(c *gin.Context) {}
func PostMovies(c *gin.Context) {}
func GetMovieComments(c *gin.Context) {}
func PostMovieComments(c *gin.Context) {}
func PostMovieRating(c *gin.Context) {}
func PostMovieFavorite(c *gin.Context) {}
func PostMovieCast(c *gin.Context) {}


