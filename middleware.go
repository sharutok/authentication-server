package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Applicationlist []string

var tokenAssign Token

func CheckIfAppNameExist(c *gin.Context) {
	rdb := ConnectToRedis()

	res, err := rdb.HGetAll(ctx, "authentication_server_app_sub").Result()
	if err != nil {
		log.Fatalf("Could not retrieve hash: %v", err)
	}

	if strings.HasPrefix(c.GetHeader("Authorization"), "Bearer") && os.Getenv("AUTH_TOKEN") == strings.Split(c.GetHeader("Authorization"), " ")[1] {
		app := c.Param("app")
		if _, exists := res[app]; exists {
			value := strings.Split(res[app], ":")
			tokenAssign.hmacSampleSecretAccessToken = []byte(value[0])
			tokenAssign.hmacSampleSecretRefreshToken = []byte(value[1])
			c.Next()
		} else {
			log.Println("No match")
			c.JSON(404, gin.H{"error": "App not found"})
			c.Abort()
		}
	} else {
		c.JSON(404, gin.H{"error": "no proper token"})
		c.Abort()
	}
}
