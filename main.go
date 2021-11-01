package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}

func main() {
	r := gin.Default()

	r.Handle(http.MethodGet, "/", Handler)

	r.Run(":8000")
}
