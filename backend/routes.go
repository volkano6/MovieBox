package backend

import (
	"github.com/C305DatabaseProject/database-project/backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/homepage", controllers.Hompage)
		api.GET("/profile", controllers.Hompage) // JWT

		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
			auth.POST("/logout", controllers.Logout)
		}

		users := api.Group("/user/:id")
		{
			users.GET("/", controllers.GetUser)
			users.GET("/watchlists", controllers.Watchlist)
			users.GET("/watched", controllers.Watched)
			users.GET("/favorites", controllers.Favorites)
			users.GET("/comments", controllers.Comments)
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

GET /api/profile/ (Needs JWT header) *

GET /api/user/:id/ *
GET /api/user/:id/watchlist/ *
GET /api/user/:id/watched/ *
GET /api/user/:id/favorites/ *

PUT /api/profile/


GET /api/movies/:id/
GET /api/movies/:id/comments/

POST /api/movies/
POST /api/comments/
POST /api/rating/

PUT /api/movies/:id/
PUT /api/comments/:id/

DELETE /api/movies/:id/
DELETE /api/comments/:id


GET /api/actors/:id

POST /api/actors/

PUT /api/actors/:id

DELETE /api/actors/:id

*/
