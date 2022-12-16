package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	logged := true
	// Check if user exists
	_, exists := c.Get("user")
	if !exists {
		logged = false
	}
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
	c.JSON(http.StatusOK, OkResponse{
		Status: "ok",
		Logged: logged,
		Data:   movies,
	})
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

func PostMovieWatched(c *gin.Context) {
	id := c.Param("id")
	user, exists := c.Get("user")
	if !exists {
		// Don't allow post
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	sql := `INSERT INTO user_watched (movie_id, user_id, watched_date) VALUES (?, ?, ?);`
	_, err := database.DB.Exec(sql, id, user.(User).ID, time.Now().UTC())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage("Successfully added to watched."))
}

func PostMovieWatchlist(c *gin.Context) {
	id := c.Param("id")
	user, exists := c.Get("user")
	if !exists {
		// Don't allow post
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	sql := `INSERT INTO user_watchlist (movie_id, user_id, added_date) VALUES (?, ?, ?);`
	_, err := database.DB.Exec(sql, id, user.(User).ID, time.Now().UTC())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage("Successfully added to watchlist."))
}

func PostMovieFavorite(c *gin.Context) {
	id := c.Param("id")
	user, exists := c.Get("user")
	if !exists {
		// Don't allow post
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	sql := `INSERT INTO user_favorites (movie_id, user_id, favorited_date) VALUES (?, ?, ?);`
	_, err := database.DB.Exec(sql, id, user.(User).ID, time.Now().UTC())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage("Successfully added to favorites."))
}

func PostMovieRating(c *gin.Context) {
	id := c.Param("id")
	rating := c.Query("star")
	fmt.Println(rating)
	user, exists := c.Get("user")
	if rating == "" {
		c.JSON(http.StatusOK, ErrorMessage("Invalid rating."))
		return
	}
	if !exists {
		// Don't allow post
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	sql := `INSERT INTO user_ratings (movie_id, user_id, rating) VALUES (?, ?, ?);`
	_, err := database.DB.Exec(sql, id, user.(User).ID, rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	sql = `UPDATE movies SET rating = (SELECT AVG(rating) FROM user_ratings WHERE movie_id = id) WHERE id = ?;`
	_, err = database.DB.Exec(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage("Successfully added rating."))
}

func PostMovieComment(c *gin.Context) {
	id := c.Param("id")
	user, exists := c.Get("user")
	var comment CommentPost
	err := c.BindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if !exists {
		// Don't allow post
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	sql := `INSERT INTO user_comments (movie_id, user_id, comment, comment_date) VALUES (?, ?, ?, ?);`
	_, err = database.DB.Exec(sql, id, user.(User).ID, comment.Comment, time.Now().UTC())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage("Successfully added comment."))
}
