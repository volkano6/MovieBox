package controllers

import (
	"net/http"
	"time"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	// Check if user exists
	_, exists := c.Get("user")
	logged := true
	if !exists {
		// User not logged in
		// ; Show login/register button
		logged = false
	}
	var monthlyPopular []Movie
	var allTimeFavorites []Movie
	var latestReviews []Comment
	//userObject := user.(User)
	// User logged in
	// ; Show profile button

	// Get watched by current month and rank based on count per every movie
	sql := `SELECT COUNT(user_watched.movie_id) AS watched_count, (SELECT COUNT(movie_id) FROM user_favorites WHERE user_favorites.movie_id = user_watched.movie_id) AS favorite_count, user_watched.movie_id, movies.title, movies.description, movies.release_date, movies.poster, movies.length, movies.rating FROM user_watched 
	INNER JOIN movies ON user_watched.movie_id = movies.id 
	WHERE MONTH(watched_date) = ?
	GROUP BY user_watched.movie_id, movies.title, movies.release_date, movies.poster, movies.length, movies.rating
	LIMIT 5;`
	rows, err := database.DB.Query(sql, int(time.Now().Month()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.MonthlyWatchedCount, &movie.FavoriteCount, &movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Poster, &movie.Length, &movie.Rating)
		monthlyPopular = append(monthlyPopular, movie)
	}
	var totalWatched int
	sql = `SELECT COUNT(movie_id) FROM user_favorites WHERE user_favorites.movie_id = ?;`
	for i := 0; i < len(monthlyPopular); i++ {
		err = database.DB.QueryRow(sql, monthlyPopular[i].ID).Scan(&totalWatched)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		monthlyPopular[i].WatchedCount = totalWatched
	}
	// Get favorites by all time and rank based on count per every movie
	// Return movie_id, movie tite, release date, poster, length, rating, genre, watched, favorites
	sql = `SELECT COUNT(movie_id) AS favorite_count, (SELECT COUNT(movie_id) FROM user_watched WHERE user_watched.movie_id = user_favorites.movie_id) AS watched_count, movie_id, movies.title, movies.description, movies.release_date, movies.poster, movies.length, movies.rating FROM user_favorites 
	INNER JOIN movies ON movie_id = movies.id 
	GROUP BY movie_id, movies.title, movies.description, movies.release_date, movies.poster, movies.length, movies.rating
	LIMIT 5;`
	rows, err = database.DB.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	var movie2 Movie
	for rows.Next() {
		rows.Scan(&movie2.FavoriteCount, &movie2.WatchedCount, &movie2.ID, &movie2.Title, &movie2.Description, &movie2.ReleaseDate, &movie2.Poster, &movie2.Length, &movie2.Rating)
		allTimeFavorites = append(allTimeFavorites, movie2)
	}
	// Get 5 comments ranked by recent
	// Return user_id, display name, user pp, comment, comment date, movie name, movie id
	sql = `SELECT user_id, users.displayname, users.avatar, comment, comment_date, movies.title, movie_id
	FROM user_comments INNER JOIN users ON user_id = users.id INNER JOIN movies ON movie_id = movies.id
	ORDER BY comment_date DESC
	LIMIT 5;`
	rows, err = database.DB.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	var comment Comment
	for rows.Next() {
		rows.Scan(&comment.UserID, &comment.Username, &comment.UserAvatar, &comment.Comment, &comment.CommentDate, &comment.MovieTitle, &comment.MovieID)
		latestReviews = append(latestReviews, comment)
	}
	c.JSON(http.StatusOK, HomeResponse{
		Status:           "ok",
		Logged:           logged,
		CurrentMonth:     time.Now().Month().String(),
		MonthlyPopular:   monthlyPopular,
		AllTimeFavorites: allTimeFavorites,
		LatestReviews:    latestReviews,
	})
}

func Profile(c *gin.Context) {
	// Check if user exists
	user, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in / invalid token."))
		return
	}
	// User logged in
	c.JSON(http.StatusOK, OkMessage(user))
}
