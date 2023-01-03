package controllers

import (
	"net/http"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	var profileResponse ProfileResponse
	// Check if user with given id exists
	sql := `SELECT id, username, displayname, email, dateofbirth, avatar, bio, location, social_twitter, social_instagram
		FROM users WHERE id = ?`
	database.DB.QueryRow(sql, id).Scan(&user.ID, &user.Username, &user.Displayname, &user.Email, &user.DateOfBirth, &user.Avatar, &user.Bio, &user.Location, &user.Twitter, &user.Instagram)
	if user.ID == 0 {
		// User not found
		c.JSON(http.StatusOK, ErrorMessage("User not found."))
		return
	}
	// Check if user exists
	localUser, exists := c.Get("user")
	profileResponse.Logged = true
	perm := "Default"
	if !exists {
		// User not logged in
		profileResponse.Logged = false
	} else {
		profileResponse.LoggedID = localUser.(User).ID
		sql := `SELECT permission FROM user_admins WHERE user_id = ?;`
		err := database.DB.QueryRow(sql, localUser.(User).ID).Scan(&perm)
		if err != nil {
			perm = "Default"
		}
	}
	profileResponse.User = user
	// movies_watched
	sql = `SELECT id, title, description, release_date, poster, rating, length, movie_genres.genre_name,
	(SELECT COUNT(movie_id) FROM user_favorites WHERE user_favorites.movie_id = movies.id) AS favorite_count, 
	(SELECT COUNT(movie_id) FROM user_watched WHERE user_watched.movie_id = movies.id) AS watched_count FROM movies 
	INNER JOIN user_watched	ON movies.id = user_watched.movie_id
	INNER JOIN movie_genres ON movies.id = movie_genres.movie_id
	WHERE user_watched.user_id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies_watched []Movie
	for rows.Next() {
		var movie_watched Movie
		var genre string
		rows.Scan(&movie_watched.ID, &movie_watched.Title, &movie_watched.Description, &movie_watched.ReleaseDate, &movie_watched.Poster, &movie_watched.Rating, &movie_watched.Length, &genre, &movie_watched.FavoriteCount, &movie_watched.WatchedCount)
		if len(movies_watched) != 0 && movies_watched[len(movies_watched)-1].ID == movie_watched.ID {
			movies_watched[len(movies_watched)-1].Genres = append(movies_watched[len(movies_watched)-1].Genres, genre)
			continue
		}
		movie_watched.Genres = append(movie_watched.Genres, genre)
		movies_watched = append(movies_watched, movie_watched)
	}
	profileResponse.UserWatched = movies_watched
	//movies_watchlist
	sql = `SELECT id, title, description, release_date, poster, rating, LENGTH, movie_genres.genre_name,
	(SELECT COUNT(movie_id) FROM user_favorites WHERE user_favorites.movie_id = movies.id) AS favorite_count, 
	(SELECT COUNT(movie_id) FROM user_watched WHERE user_watched.movie_id = movies.id) AS watched_count FROM movies 
	INNER JOIN user_watchlist ON movies.id = user_watchlist.movie_id
	INNER JOIN movie_genres ON movies.id = movie_genres.movie_id
	WHERE user_watchlist.user_id = ?;`
	rows, err = database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies_watchlist []Movie
	for rows.Next() {
		var movie_watchlist Movie
		var genre string
		rows.Scan(&movie_watchlist.ID, &movie_watchlist.Title, &movie_watchlist.Description, &movie_watchlist.ReleaseDate, &movie_watchlist.Poster, &movie_watchlist.Rating, &movie_watchlist.Length, &genre, &movie_watchlist.FavoriteCount, &movie_watchlist.WatchedCount)
		if len(movies_watchlist) != 0 && movies_watchlist[len(movies_watchlist)-1].ID == movie_watchlist.ID {
			movies_watchlist[len(movies_watchlist)-1].Genres = append(movies_watchlist[len(movies_watchlist)-1].Genres, genre)
			continue
		}
		movie_watchlist.Genres = append(movie_watchlist.Genres, genre)
		movies_watchlist = append(movies_watchlist, movie_watchlist)
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
	profileResponse.UserComments = comments
	//movies_favorite
	sql = `SELECT id, title, description, release_date, poster, rating, LENGTH, movie_genres.genre_name,
	(SELECT COUNT(movie_id) FROM user_favorites WHERE user_favorites.movie_id = movies.id) AS favorite_count, 
	(SELECT COUNT(movie_id) FROM user_watched WHERE user_watched.movie_id = movies.id) AS watched_count FROM movies 
	INNER JOIN user_favorites ON movies.id = user_favorites.movie_id
	INNER JOIN movie_genres ON movies.id = movie_genres.movie_id
	WHERE user_favorites.user_id = ?;`
	rows, err = database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies_favorite []Movie
	for rows.Next() {
		var movie_favorite Movie
		var genre string
		rows.Scan(&movie_favorite.ID, &movie_favorite.Title, &movie_favorite.Description, &movie_favorite.ReleaseDate, &movie_favorite.Poster, &movie_favorite.Rating, &movie_favorite.Length, &genre, &movie_favorite.FavoriteCount, &movie_favorite.WatchedCount)
		if len(movies_favorite) != 0 && movies_favorite[len(movies_favorite)-1].ID == movie_favorite.ID {
			movies_favorite[len(movies_favorite)-1].Genres = append(movies_favorite[len(movies_favorite)-1].Genres, genre)
			continue
		}
		movie_favorite.Genres = append(movie_favorite.Genres, genre)
		movies_favorite = append(movies_favorite, movie_favorite)
	}
	profileResponse.UserFavorites = movies_favorite
	profileResponse.Status = "ok"
	profileResponse.Perm = perm
	c.JSON(http.StatusOK, profileResponse)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user, exists := c.Get("user")
	var updatedUser UserUpdate
	err := c.BindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if !exists {
		// Don't allow update
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	var adminID int
	var permission string
	sql := `SELECT user_id, permission FROM user_admins
	WHERE user_id = ?;`
	err = database.DB.QueryRow(sql, user.(User).ID).Scan(&adminID, &permission)
	if err != nil || adminID == 0 || permission != "Admin" {
		c.JSON(http.StatusForbidden, ErrorMessage("Insufficient permissions."))
		return
	}
	sql = `UPDATE users SET displayname = ?, email = ?, dateofbirth = ?, avatar = ?, bio = ?, location = ?, social_twitter = ?, social_instagram = ?
	WHERE id = ?;`
	_, err = database.DB.Exec(sql, updatedUser.Displayname, updatedUser.Email, updatedUser.DateOfBirth, updatedUser.Avatar, updatedUser.Bio, updatedUser.Location, updatedUser.Twitter, updatedUser.Instagram, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkResponse{
		Status: "ok",
		Logged: true,
		Data:   "Successfully updated user.",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user, exists := c.Get("user")
	if !exists {
		// Don't allow update
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	var adminID int
	var permission string
	sql := `SELECT user_id, permission FROM user_admins
	WHERE user_id = ?;`
	err := database.DB.QueryRow(sql, user.(User).ID).Scan(&adminID, &permission)
	if err != nil || adminID == 0 || permission != "Admin" {
		c.JSON(http.StatusForbidden, ErrorMessage("Insufficient permissions."))
		return
	}
	sql = `DELETE FROM users WHERE id = ?;`
	_, err = database.DB.Exec(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkResponse{
		Status: "ok",
		Logged: true,
		Data:   "Successfully deleted user.",
	})
}

func UpdateProfile(c *gin.Context) {
	user, exists := c.Get("user")
	var updatedUser UserUpdate
	err := c.BindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if !exists {
		// Don't allow update
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	sql := `UPDATE users SET displayname = ?, email = ?, dateofbirth = ?, avatar = ?, bio = ?, location = ?, social_twitter = ?, social_instagram = ?
	WHERE id = ?;`
	_, err = database.DB.Exec(sql, updatedUser.Displayname, updatedUser.Email, updatedUser.DateOfBirth, updatedUser.Avatar, updatedUser.Bio, updatedUser.Location, updatedUser.Twitter, updatedUser.Instagram, user.(User).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkResponse{
		Status: "ok",
		Logged: true,
		Data:   "Successfully updated user.",
	})
}
