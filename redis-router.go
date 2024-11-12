package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddApp struct {
	HashKey string `json:"hash_key"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

type DeleteApp struct {
	HashKey string `json:"hash_key"`
	Key     string `json:"key"`
}

func ADD_APP(c *gin.Context) {
	var addApp AddApp
	if err := c.ShouldBindJSON(&addApp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	AddField(addApp.HashKey, addApp.Key, addApp.Value)
	c.JSON(http.StatusAccepted, nil)
}

func DELETE_APP(c *gin.Context) {
	var deleteApp DeleteApp
	if err := c.ShouldBindJSON(&deleteApp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	DeleteField(deleteApp.HashKey, deleteApp.Key)
	c.JSON(http.StatusAccepted, nil)
}
