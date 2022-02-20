package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richaidos/SupportSection/configuration"
)

var iSettings *configuration.InitSettings

func ConfigMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("iSettings", iSettings)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	iSettings = configuration.GetInitSettings()
	r.Use(ConfigMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(iSettings.Config.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
