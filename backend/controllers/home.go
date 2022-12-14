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
	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.MonthlyWatchedCount, &movie.FavoriteCount, &movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Poster, &movie.Length, &movie.Rating)
		monthlyPopular = append(monthlyPopular, movie)
	}
	for i := 0; i < len(monthlyPopular); i++ {
		// Get total watched count for each movie
		var totalWatched int
		sql = `SELECT COUNT(movie_id) FROM user_watched WHERE user_watched.movie_id = ?;`
		err = database.DB.QueryRow(sql, monthlyPopular[i].ID).Scan(&totalWatched)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		monthlyPopular[i].WatchedCount = totalWatched
		// Get genres for each movie
		var genres []string
		var genre string
		sql = `SELECT genre_name FROM movie_genres WHERE movie_id = ?;`
		rows, err = database.DB.Query(sql, monthlyPopular[i].ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		for rows.Next() {
			rows.Scan(&genre)
			genres = append(genres, genre)
		}
		monthlyPopular[i].Genres = genres
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
	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.FavoriteCount, &movie.WatchedCount, &movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Poster, &movie.Length, &movie.Rating)
		allTimeFavorites = append(allTimeFavorites, movie)
	}
	for i := 0; i < len(allTimeFavorites); i++ {
		// Get genres for each movie
		var genres []string
		var genre string
		sql = `SELECT genre_name FROM movie_genres WHERE movie_id = ?;`
		rows, err = database.DB.Query(sql, allTimeFavorites[i].ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		for rows.Next() {
			rows.Scan(&genre)
			genres = append(genres, genre)
		}
		allTimeFavorites[i].Genres = genres
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
	var profileResponse ProfileResponse
	// Check if user exists
	user, exists := c.Get("user")
	profileResponse.Logged = true
	if !exists {
		// User not logged in
		c.JSON(http.StatusUnauthorized, ErrorMessage("User not logged in."))
		return
	}
	userObject := user.(User)
	id := userObject.ID
	profileResponse.User = userObject

	// movies_watched
	sql := `SELECT id, title, description, release_date, poster, rating, length
	FROM movies LEFT JOIN user_watched
	ON movies.id = user_watched.movie_id
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
		rows.Scan(&movie_watched.ID, &movie_watched.Title, &movie_watched.Description, &movie_watched.ReleaseDate, &movie_watched.Poster, &movie_watched.Rating, &movie_watched.Length)
		movies_watched = append(movies_watched, movie_watched)
	}
	for i := 0; i < len(movies_watched); i++ {
		var genres []string
		var genre string
		sql = `SELECT genre_name FROM movie_genres WHERE movie_id = ?;`
		rows, err = database.DB.Query(sql, movies_watched[i].ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		for rows.Next() {
			rows.Scan(&genre)
			genres = append(genres, genre)
		}
		movies_watched[i].Genres = genres

		var totalFavorites int
		sql = `SELECT COUNT(movie_id) FROM user_favorites WHERE movie_id = ?;`
		err = database.DB.QueryRow(sql, movies_watched[i].ID).Scan(&totalFavorites)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		movies_watched[i].FavoriteCount = totalFavorites

		var totalWatched int
		sql = `SELECT COUNT(movie_id) FROM user_watched WHERE movie_id = ?;`
		err = database.DB.QueryRow(sql, movies_watched[i].ID).Scan(&totalWatched)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		movies_watched[i].WatchedCount = totalWatched
	}
	profileResponse.UserWatched = movies_watched

	//movies_watchlist
	sql = `SELECT id, title, description, release_date, poster, rating, length
	FROM movies LEFT JOIN user_watchlist
	ON movies.id = user_watchlist.movie_id
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
		rows.Scan(&movie_watchlist.ID, &movie_watchlist.Title, &movie_watchlist.Description, &movie_watchlist.ReleaseDate, &movie_watchlist.Poster, &movie_watchlist.Rating, &movie_watchlist.Length)
		movies_watchlist = append(movies_watchlist, movie_watchlist)
	}
	for i := 0; i < len(movies_watchlist); i++ {
		var genres []string
		var genre string
		sql = `SELECT genre_name FROM movie_genres WHERE movie_id = ?;`
		rows, err = database.DB.Query(sql, movies_watchlist[i].ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		for rows.Next() {
			rows.Scan(&genre)
			genres = append(genres, genre)
		}
		movies_watchlist[i].Genres = genres

		var totalFavorites int
		sql = `SELECT COUNT(movie_id) FROM user_favorites WHERE movie_id = ?;`
		err = database.DB.QueryRow(sql, movies_watchlist[i].ID).Scan(&totalFavorites)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		movies_watchlist[i].FavoriteCount = totalFavorites

		var totalWatched int
		sql = `SELECT COUNT(movie_id) FROM user_watched WHERE movie_id = ?;`
		err = database.DB.QueryRow(sql, movies_watchlist[i].ID).Scan(&totalWatched)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		movies_watchlist[i].WatchedCount = totalWatched
	}
	profileResponse.UserWatchlist = movies_watchlist

	//comments
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
		rows.Scan(&comment.UserID, &comment.Username, &comment.MovieID, &comment.Comment, &comment.CommentDate, &comment.MovieTitle)
		comments = append(comments, comment)
	}
	if comments == nil {
		c.JSON(http.StatusNotFound, ErrorMessage("User not have any comment."))
		return
	}
	profileResponse.UserComments = comments

	//movies_favorite
	sql = `SELECT id, title, description, release_date, poster, rating, length
		FROM movies JOIN user_favorites
		ON movies.id = user_favorites.movie_id
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
		rows.Scan(&movie_favorite.ID, &movie_favorite.Title, &movie_favorite.Description, &movie_favorite.ReleaseDate, &movie_favorite.Poster, &movie_favorite.Rating, &movie_favorite.Length)
		movies_favorite = append(movies_favorite, movie_favorite)
	}
	for i := 0; i < len(movies_favorite); i++ {
		var genres []string
		var genre string
		sql = `SELECT genre_name FROM movie_genres WHERE movie_id = ?;`
		rows, err = database.DB.Query(sql, movies_favorite[i].ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		for rows.Next() {
			rows.Scan(&genre)
			genres = append(genres, genre)
		}
		movies_favorite[i].Genres = genres

		var totalFavorites int
		sql = `SELECT COUNT(movie_id) FROM user_favorites WHERE movie_id = ?;`
		err = database.DB.QueryRow(sql, movies_favorite[i].ID).Scan(&totalFavorites)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		movies_favorite[i].FavoriteCount = totalFavorites

		var totalWatched int
		sql = `SELECT COUNT(movie_id) FROM user_watched WHERE movie_id = ?;`
		err = database.DB.QueryRow(sql, movies_favorite[i].ID).Scan(&totalWatched)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
			return
		}
		movies_favorite[i].WatchedCount = totalWatched
	}
	profileResponse.UserFavorites = movies_favorite
	profileResponse.Status = "ok"

	// User logged in
	c.JSON(http.StatusOK, profileResponse)
}
