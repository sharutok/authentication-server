package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GET_ACCESS_TOKEN(c *gin.Context) {
	var jsonBody RequestBody
	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tokens := GenerateTokens(jsonBody.Obj)
	c.JSON(http.StatusOK, gin.H{
		"accessToken":  tokens["accessToken"],
		"refreshToken": tokens["refreshToken"],
		"created_at":   time.Now(),
	})
}

func GET_REFRESH_TOKEN(c *gin.Context) {
	var jsonBody struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newAccessToken, err := RefreshAccessToken(jsonBody.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to refresh access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"newAccessToken": newAccessToken})
}

func VALIDATE_ACCESS_TOKEN(c *gin.Context) {
	var jsonBody struct {
		AccessToken string `json:"accessToken"`
	}
	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	obj, err := ValidateAccessToken(jsonBody.AccessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access token is valid", "obj": obj})
}
