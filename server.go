package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richaidos/SupportSection/controller"
	d "github.com/richaidos/SupportSection/database"
	g "github.com/richaidos/SupportSection/globals"
)

func main() {
	r := gin.Default()

	//Инициализация глобальных переменных
	g.Init()

	//Инициализация базы данных
	d.Init()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api")
	{
		v1.GET("search", controller.SearchByQuery)
	}

	r.Run(g.ISettings.Config.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
