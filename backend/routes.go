package backend

import (
	"github.com/C305DatabaseProject/database-project/backend/controllers"
	"github.com/C305DatabaseProject/database-project/backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/homepage", middleware.CheckAuth, controllers.Homepage)
		api.GET("/profile", middleware.CheckAuth, controllers.Profile)

		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
			auth.POST("/logout", controllers.Logout)
		}
		users := api.Group("/user/:id")
		{
			users.GET("/", middleware.CheckAuth, controllers.GetUser)
			users.GET("/watchlist", controllers.Watchlist)
			users.GET("/watched", controllers.Watched)
			users.GET("/favorites", controllers.Favorites)
			users.GET("/comments", controllers.Comments)
		}
		movies := api.Group("/movies")
		{
			//   /api/movies/...
			movies.GET("/", middleware.CheckAuth, controllers.GetMovies)
			movies.POST("/", controllers.PostMovies)
			id := movies.Group("/:id")
			{
				//   /api/movies/:id/...
				id.GET("/", middleware.CheckAuth, controllers.GetMovie)
				id.GET("/comments", controllers.GetMovieComments)
				id.POST("/comments", controllers.PostMovieComments)
				id.POST("/ratings", controllers.PostMovieRating)
				id.POST("/favorite", controllers.PostMovieFavorite)
				id.POST("/cast", controllers.PostMovieCast)

			}

		}
	}
}

/*

Planned database API endpoints.

GET /auth/token/

POST /auth/logout/ *
POST /auth/login/ *
POST /auth/register/ *

GET /api/homepage/ *

GET /api/profile/ (Needs JWT header) * - Sonra, Arda

PUT /api/profile/ (Needs JWT header) - Sonra, Arda

GET /api/user/:id/ *
GET /api/user/:id/watchlists/ * - Arda
GET /api/user/:id/watched/ *
GET /api/user/:id/favorites/ *
GET /api/user/:id/comments/ *

POST /api/user/:id/watchlists/ 	Request: watchlist_name - Arda

POST /api/user/:id/watchlists/:id/movie/ 	Request: watchlist_id, movie_id - Arda

DELETE /api/user/:id/watchlists/:id/movie/ 	Request: watchlist_id, movie_id - Arda

GET /api/movies/:id/	*		Data: id, title, description, release_date, poster, rating, length, trailer
GET /api/movies/:id/comments/ *	Data: users.id, users.displayname, comment, comment_date

POST /api/movies/		*		Request: title, description, release_date, poster, rating, length, trailer
POST /api/movies/:id/comments/*	Request: movie_id, user_id, comment, comment_date
POST /api/movies/:id/rating/*	Request: movie_id, user_id, rating
POST /api/movies/:id/favorite/*	Request: movie_id, user_id, favorited_date
POST /api/movies/:id/cast/ *	Request: movie_id, actor_id [SAME WITH POST /api/actors/:id/cast]

PUT /api/movies/:id/			Request: title, description, release_date, poster, rating, length, trailer
PUT /api/movies/:id/comments/	Request: user_id, comment
PUT /api/movies/:id/rating/		Request: movie_id, user_id, rating
PUT /api/movies/:id/cast/		Request: movie_id, actor_id [SAME WITH PUT /api/actors/:id/cast]

DELETE /api/movies/:id/				Delete
DELETE /api/movies/:id/comments/	Delete, Request: movie_id, user_id
DELETE /api/movies/:id/rating/		Delete, Request: movie_id, user_id
DELETE /api/movies/:id/favorite/	Delete, Request: movie_id, user_id

GET /api/actors/:id/			Data: id, fullname, dateofbirth, gender, bio, location, films[actor_filmography.movie_id, actor_filmography.actor_id]

POST /api/actors/				Request: fullname, dateofbirth, gender, bio, location
POST /api/actors/:id/cast/		Request: movie_id, actor_id [SAME WITH POST /api/movies/:id/cast]

PUT /api/actors/:id				Request: fullname, dateofbirth, gender, bio, location
PUT /api/actors/:id/cast/		Request: movie_id, actor_id [SAME WITH PUT /api/movies/:id/cast]

DELETE /api/actors/:id/			Delete


*/
