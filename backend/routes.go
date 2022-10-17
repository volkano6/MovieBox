package backend

import (
	"github.com/C305DatabaseProject/database-project/backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")

		v1.GET("/health", controllers.Health)
	}
}
