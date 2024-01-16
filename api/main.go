package main

import (
	"desafio-rate-limiter-go-expert-pos/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	limiter := pkg.Initialize()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		pkg.Middleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			c.Next()
		}), limiter.Storage).ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.Run(":8080")
}
