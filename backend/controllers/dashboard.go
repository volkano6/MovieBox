package controllers

import (
	"net/http"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		// Don't allow post
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	var id int
	var permission string
	sql := `SELECT user_id, permission FROM user_admins
	WHERE user_id = ?;`
	err := database.DB.QueryRow(sql, user.(User).ID).Scan(&id, &permission)
	if err != nil || id == 0 || permission != "Admin" {
		c.JSON(http.StatusForbidden, ErrorMessage("Insufficient permissions."))
		return
	}
	// Retrieve movies
	sql = `SELECT id, title, description, release_date, poster, rating, length, movie_genres.genre_name,
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
	// Retrieve users
	sql = `SELECT id, username, displayname, email, dateofbirth, avatar, bio, location, social_twitter, social_instagram FROM users`
	rows, err = database.DB.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Username, &user.Displayname, &user.Email, &user.DateOfBirth, &user.Avatar, &user.Bio, &user.Location, &user.Twitter, &user.Instagram)
		users = append(users, user)
	}
	c.JSON(http.StatusOK, DashboardResponse{
		Status: "ok",
		Logged: true,
		Movies: movies,
		Users:  users,
	})
}
