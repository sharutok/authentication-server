package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HEALTH_CHECK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
