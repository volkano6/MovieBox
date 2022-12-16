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
	for rows.Next() {
		var movie Movie
		var genre string
		rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Poster, &movie.Rating, &movie.Length, &genre, &movie.FavoriteCount, &movie.WatchedCount)
		if len(movies) != 0 && movies[len(movies)-1].ID == movie.ID {
			movies[len(movies)-1].Genres = append(movies[len(movies)-1].Genres, genre)
			continue
		}
		movie.Genres = append(movie.Genres, genre)
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, OkMessage(movies))
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")
	logged := true
	// Check if user exists
	_, exists := c.Get("user")
	if !exists {
		logged = false
	}
	// Movie Info
	sql := `SELECT id, title, description, release_date, poster, rating, length, movie_genres.genre_name,
	(SELECT COUNT(movie_id) FROM user_favorites WHERE user_favorites.movie_id = movies.id) AS favorite_count, 
	(SELECT COUNT(movie_id) FROM user_watched WHERE user_watched.movie_id = movies.id) AS watched_count 
	FROM movies	INNER JOIN movie_genres ON id = movie_genres.movie_id
	WHERE id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	var movie Movie
	for rows.Next() {
		var genre string
		rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Poster, &movie.Rating, &movie.Length, &genre, &movie.FavoriteCount, &movie.WatchedCount)
		movie.Genres = append(movie.Genres, genre)
	}
	// Comments
	sql = `SELECT user_id, users.displayname, users.avatar, comment, comment_date, movies.title, movie_id
	FROM user_comments INNER JOIN users ON user_id = users.id INNER JOIN movies ON movie_id = movies.id
	WHERE movies.id = ?;`
	rows, err = database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	var comments []Comment
	var comment Comment
	for rows.Next() {
		rows.Scan(&comment.UserID, &comment.Username, &comment.UserAvatar, &comment.Comment, &comment.CommentDate, &comment.MovieTitle, &comment.MovieID)
		comments = append(comments, comment)
	}
	c.JSON(http.StatusOK, MovieResponse{
		Status:   "ok",
		Logged:   logged,
		Movie:    movie,
		Comments: comments,
	})
}

func PostMovies(c *gin.Context)        {}
func GetMovieComments(c *gin.Context)  {}
func PostMovieComments(c *gin.Context) {}
func PostMovieRating(c *gin.Context)   {}
func PostMovieFavorite(c *gin.Context) {}
func PostMovieCast(c *gin.Context)     {}
