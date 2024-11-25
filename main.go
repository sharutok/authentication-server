package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Token struct {
	hmacSampleSecretAccessToken  []byte
	hmacSampleSecretRefreshToken []byte
}

type AuthToken string

func main() {
	router := gin.Default()
	godotenv.Load(".env")

	router.POST("/get-token/:app", CheckIfAppNameExist, GET_ACCESS_TOKEN)
	router.POST("/refresh-token/:app", CheckIfAppNameExist, GET_REFRESH_TOKEN)
	router.POST("/validate-access-token/:app", CheckIfAppNameExist, VALIDATE_ACCESS_TOKEN)

	router.GET("/health-check", HEALTH_CHECK)

	router.POST("/add-app", ADD_APP)
	router.POST("/delete-app", DELETE_APP)

	router.Run("localhost:8000")
}
