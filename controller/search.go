package controller

import "github.com/gin-gonic/gin"

func SearchByQuery(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
