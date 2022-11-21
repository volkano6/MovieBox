package backend

import (
	"github.com/C305DatabaseProject/database-project/backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", controllers.Health)
	}
}

/*

Planned database API endpoints.

GET /auth/token/

POST /auth/logout/
POST /auth/login/
POST /auth/register/

GET /api/homepage/

GET /api/profile/ (Needs JWT header)

GET /api/users/:id/
GET /api/users/:id/watchlists/
GET /api/users/:id/watched/
GET /api/users/:id/favorites/ TODO: Implement favorites table.

PUT /api/profile/


GET /api/movies/:id/
GET /api/movies/:id/comments/

POST /api/movies/
POST /api/comments/

PUT /api/movies/:id/
PUT /api/comments/:id/

DELETE /api/movies/:id/
DELETE /api/comments/:id


GET /api/actors/:id

POST /api/actors/

PUT /api/actors/:id

DELETE /api/actors/:id

*/
