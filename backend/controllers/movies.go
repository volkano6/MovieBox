package controllers

import (
	"net/http"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	// Get all movie data from the database with genres, favorite and watched count
	sql := `SELECT id, title, description, release_date, poster, rating, length, movie_genres.genre_name,
	(SELECT COUNT(movie_id) FROM user_favorites WHERE user_favorites.movie_id = movies.id) AS favorite_count, 
	(SELECT COUNT(movie_id) FROM user_watched WHERE user_watched.movie_id = movies.id) AS watched_count 
	FROM movies	INNER JOIN movie_genres ON id = movie_genres.movie_id;`
	rows, err := database.DB.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	var movies []Movie
	var movie Movie
	var genre string
	for rows.Next() {
		rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Poster, &movie.Rating, &movie.Length, &genre, &movie.FavoriteCount, &movie.WatchedCount)
		if len(movies) != 0 && movies[len(movies)-1].ID == movie.ID {
			movies[len(movies)-1].Genres = append(movies[len(movies)-1].Genres, genre)
			continue
		}
		movies = append(movies, movie)
	}
	if movies == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("No movie found in the database."))
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
