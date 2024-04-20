package server

import (
	"github.com/gin-gonic/gin"
	"thsr/m/server/receiver"
)

func Init(r receiver.Router) {
	app := gin.Default()
	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	app.Static("/static", "./static/page")

	r.RegisterAPI(app)

	// listen and serve on 0.0.0.0:8080
	if err := app.Run(); err != nil {
		panic(err)
	}
}
