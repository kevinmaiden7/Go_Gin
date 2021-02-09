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

	app.GET("/opt", func(c *gin.Context) {
		c.String(http.StatusOK, quote.Opt())
	})

	app.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "index OK (200)",
		})
	})

	app.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	app.GET("/concat", func(c *gin.Context) {
		num1 := c.Query("a")
		num2 := c.Query("b")
		num3 := c.Query("c")
		c.String(http.StatusOK, num1+num2+num3)
	})

	app.Run() // by default listen and serve on 0.0.0.0:8080
	// app.Run(":3000") for a hard coded port
}
