package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rsc.io/quote"
)

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, quote.Hello())
	})

	app.GET("/go", func(c *gin.Context) {
		c.String(http.StatusOK, quote.Go())
	})

	app.GET("/glass", func(c *gin.Context) {
		c.String(http.StatusOK, quote.Glass())
	})

	app.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "index OK",
		})
	})

	app.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	app.Run() // by default listen and serve on 0.0.0.0:8080
	// app.Run(":3000") for a hard coded port
}
