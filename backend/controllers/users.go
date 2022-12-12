package controllers

import (
	"net/http"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	// Check if user with given id exists
	sql := `SELECT id, username, displayname, email, dateofbirth, avatar, bio, location, social_twitter, social_instagram, type
		FROM users WHERE id = ?`
	database.DB.QueryRow(sql, id).Scan(&user.ID, &user.Username, &user.Displayname, &user.Email, &user.DateOfBirth, &user.Avatar, &user.Bio, &user.Location, &user.Twitter, &user.Instagram, &user.Type)
	if user.ID == 0 {
		// User not found
		c.JSON(http.StatusNotFound, ErrorMessage("User not found."))
		return
	}
	// If it exists, retrieve user data
	c.JSON(http.StatusOK, OkMessage(user))
}

func Watchlist(c *gin.Context) {
	id := c.Param("id")
	// Check Watchlist table with given user id
	sql := `SELECT id, title, description, release_date, poster, rating, length, genre_name
	FROM movies LEFT JOIN user_watchlist
	ON movies.id = user_watchlist.movie_id
	JOIN movie_genres ON movies.id = movie_genres.movie_id
	WHERE user_watchlist.user_id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies []Movie
	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Poster, &movie.Rating, &movie.Length, &movie.Genre)
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, OkMessage(movies))
}

func Watched(c *gin.Context) {
	id := c.Param("id")
	// Check watched table with given user id
	sql := `SELECT id, title, release_date, poster, rating, length, genre_name
	FROM movies LEFT JOIN user_watched
	ON movies.id = user_watched.movie_id
	JOIN movie_genres ON movies.id = movie_genres.movie_id
	WHERE user_watched.user_id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies []Movie
	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Poster, &movie.Rating, &movie.Length, &movie.Genre)
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, OkMessage(movies))
}

func Favorites(c *gin.Context) {
	id := c.Param("id")
	// Check Favorites table with given user id
	sql := `SELECT id, title, release_date, poster, rating, length, genre_name
		FROM movies JOIN user_favorites
		ON movies.id = user_favorites.movie_id
		JOIN movie_genres ON movies.id = movie_genres.movie_id
		WHERE user_favorites.user_id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies []Movie
	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Poster, &movie.Rating, &movie.Length, &movie.Genre)
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, OkMessage(movies))
}

func Comments(c *gin.Context) {
	id := c.Param("id")
	// Check comments with user id
	sql := `SELECT users.id, users.username, movie_id , comment, comment_date, title
	FROM project.user_comments LEFT JOIN project.users
	ON user_comments.user_id = project.users.id
	JOIN movies ON movie_id = movies.id
	WHERE project.users.id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return Comments details with the User ids
	var comments []Comment
	var comment Comment
	for rows.Next() {
		rows.Scan(&comment.UserID, &comment.Username, &comment.MovieID, &comment.Comment, &comment.CommentDate, &comment.Title)
		comments = append(comments, comment)
	}
	if comments == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("User not have any comment."))
		return
	}
	c.JSON(http.StatusOK, OkMessage(comments))
}
