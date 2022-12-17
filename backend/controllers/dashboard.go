package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		// Don't allow post
		c.JSON(http.StatusForbidden, ErrorMessage("User not logged in."))
		return
	}
}
