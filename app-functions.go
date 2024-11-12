package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Obj string `json:"obj"`
	jwt.StandardClaims
}

type RequestBody struct {
	Obj string `json:"obj"`
}

// var (
// 	hmacSampleSecretAccessToken  = []byte("my_secret_access_key")
// 	hmacSampleSecretRefreshToken = []byte("my_secret_refresh_key")
// )

func GenerateTokens(obj string) map[string]string {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Obj: obj,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	accessTokenString, err := accessToken.SignedString(tokenAssign.hmacSampleSecretAccessToken)
	if err != nil {
		log.Println("Error creating access token:", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Obj: obj,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	refreshTokenString, err := refreshToken.SignedString(tokenAssign.hmacSampleSecretRefreshToken)
	if err != nil {
		log.Println("Error creating refresh token:", err)
	}

	return map[string]string{"accessToken": accessTokenString, "refreshToken": refreshTokenString}
}

func RefreshAccessToken(refreshToken string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return tokenAssign.hmacSampleSecretRefreshToken, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid or expired refresh token")
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Obj: claims.Obj,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	accessTokenString, err := newAccessToken.SignedString(tokenAssign.hmacSampleSecretAccessToken)
	if err != nil {
		return "", fmt.Errorf("failed to create new access token")
	}

	return accessTokenString, nil
}

func ValidateAccessToken(accessToken string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return tokenAssign.hmacSampleSecretAccessToken, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid or expired access token")
	}
	return claims.Obj, nil
}
