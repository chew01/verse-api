package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func check(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ok")
}

func main() {
	router := gin.Default()
	router.GET("/", check)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
