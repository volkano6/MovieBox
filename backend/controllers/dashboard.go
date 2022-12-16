package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		// Don't allow post
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
	if user.(User).Type != 1 {
		// Don't allow non-admins
		c.JSON(http.StatusForbidden, ErrorMessage("Insufficient permissions."))
		return
	}
}
