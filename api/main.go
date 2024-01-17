package main

import (
	"net/http"

	"github.com/booscaaa/desafio-rate-limiter-go-expert-pos/ratelimiter"

	"github.com/gin-gonic/gin"
)

func main() {
	ratelimiter.Initialize()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		ratelimiter.Middleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.Run(":8080")
}
