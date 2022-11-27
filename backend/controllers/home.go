package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	// Check if user exists
	user, exists := c.Get("user")
	if !exists {
		// User not logged in
		// ; Show login/register button
		c.JSON(http.StatusOK, OkMessage("Not logged in."))
		return
	}
	// User logged in
	// ; Show profile button
	c.JSON(http.StatusOK, OkMessage(user))
}

func Profile(c *gin.Context) {
	// Check if user exists
	user, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in / invalid token."))
		return
	}
	// User logged in
	c.JSON(http.StatusOK, OkMessage(user))
}
