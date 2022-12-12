package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/C305DatabaseProject/database-project/backend/database"
)

func Homepage(c *gin.Context) {
	// Check if user exists
	user, exists := c.Get("user")
	if !exists {
		// User not logged in
		// ; Show login/register button
		c.JSON(http.StatusOK, OkMessage("Not logged in."))
		return
	}
	// User logged in
	// ; Show profile button
	c.JSON(http.StatusOK, OkMessage(user))
}

func Profile(c *gin.Context) {
	var profileResponse ProfileResponse
	// Check if user exists
	user, exists := c.Get("user")
	profileResponse.Logged = true
	if !exists {
		// User not logged in
		profileResponse.Logged = false
	}
	userObject := user.(User)
	id := userObject.ID
	profileResponse.User = userObject
	sql := `SELECT id, title, description, release_date, poster, rating, length, genre_name
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
	var movies_watched []Movie
	var movie_watched Movie
	for rows.Next() {
		rows.Scan(&movie_watched.ID, &movie_watched.Title, &movie_watched.Description, &movie_watched.ReleaseDate, &movie_watched.Poster, &movie_watched.Rating, &movie_watched.Length, &movie_watched.Genre)
		movies_watched = append(movies_watched, movie_watched)
	}
	profileResponse.UserWatched = movies_watched

	sql = `SELECT id, title, description, release_date, poster, rating, length, genre_name
	FROM movies LEFT JOIN user_watchlist
	ON movies.id = user_watchlist.movie_id
	JOIN movie_genres ON movies.id = movie_genres.movie_id
	WHERE user_watchlist.user_id = ?;`
	rows, err = database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies_watchlist []Movie
	var movie_watchlist Movie
	for rows.Next() {
		rows.Scan(&movie_watchlist.ID, &movie_watchlist.Title, &movie_watchlist.Description, &movie_watchlist.ReleaseDate, &movie_watchlist.Poster, &movie_watchlist.Rating, &movie_watchlist.Length, &movie_watchlist.Genre)
		movies_watchlist = append(movies_watchlist, movie_watchlist)
	}
	profileResponse.UserWatchlist = movies_watchlist

	sql = `SELECT users.id, users.username, movie_id , comment, comment_date, title
	FROM project.user_comments LEFT JOIN project.users
	ON user_comments.user_id = project.users.id
	JOIN movies ON movie_id = movies.id
	WHERE project.users.id = ?;`
	rows, err = database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
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
	profileResponse.UserComments = comments

	sql = `SELECT id, title, description, release_date, poster, rating, length, genre_name
		FROM movies JOIN user_favorites
		ON movies.id = user_favorites.movie_id
		JOIN movie_genres ON movies.id = movie_genres.movie_id
		WHERE user_favorites.user_id = ?;`
	
	rows, err = database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies_favorite []Movie
	var movie_favorite Movie
	for rows.Next() {
		rows.Scan(&movie_favorite.ID, &movie_favorite.Title, &movie_favorite.Description, &movie_favorite.ReleaseDate, &movie_favorite.Poster, &movie_favorite.Rating, &movie_favorite.Length, &movie_favorite.Genre)
		movies_favorite = append(movies_favorite, movie_favorite)
	}
	profileResponse.UserFavorites = movies_favorite 
	profileResponse.Status = "ok"

	// User logged in
	c.JSON(http.StatusOK, profileResponse)
}
