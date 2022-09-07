package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string {})
}


func main() {

	router := gin.Default()
	router.GET("/ping", Ping)
	router.Run(":3000")
}