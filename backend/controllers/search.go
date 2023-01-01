package controllers

import (
	"fmt"
	"net/http"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	search := c.Query("search")
	if search == "" {
		c.JSON(http.StatusOK, ErrorMessage("No query found."))
		return
	}
	logged := true
	perm := "Default"
	// Check if user exists
	user, exists := c.Get("user")
	if !exists {
		logged = false
	} else {
		sql := `SELECT permission FROM user_admins WHERE user_id = ?;`
		err := database.DB.QueryRow(sql, user.(User).ID).Scan(&perm)
		if err != nil {
			perm = "Default"
		}
	}
	sql := `SELECT id, title, description, release_date, poster, rating, length, movie_genres.genre_name,
	(SELECT COUNT(movie_id) FROM user_favorites WHERE user_favorites.movie_id = movies.id) AS favorite_count, 
	(SELECT COUNT(movie_id) FROM user_watched WHERE user_watched.movie_id = movies.id) AS watched_count 
	FROM movies	INNER JOIN movie_genres ON id = movie_genres.movie_id
	WHERE title LIKE ?;`
	rows, err := database.DB.Query(sql, fmt.Sprint("%"+search+"%"))
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
	var users []User
	sql = `SELECT id, username, displayname, email, dateofbirth, avatar, bio, location, social_twitter, social_instagram
		FROM users WHERE username = ? OR displayname = ?;`
	rows, err = database.DB.Query(sql, search, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Username, &user.Displayname, &user.Email, &user.DateOfBirth, &user.Avatar, &user.Bio, &user.Location, &user.Twitter, &user.Instagram)
		users = append(users, user)
	}
	c.JSON(http.StatusOK, SearchResponse{
		Status: "ok",
		Logged: logged,
		Perm:   perm,
		Movies: movies,
		Users:  users,
	})
}
