package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	data map[string]string
)


func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string {})
}


func PostData(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	data[key] = value
	c.JSON(http.StatusOK, map[string]string { "value": data[key] })
}

func DeleteData(c *gin.Context) {
	//key := c.Param("key")

}

func GetData(c *gin.Context) {
	key := c.Param("key")

	c.JSON(http.StatusOK, map[string]string { "value": data[key] })
}

func PutData(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	data[key] = value
	c.JSON(http.StatusOK, map[string]string { "value": data[key] })
}

func DumpData(c *gin.Context) {
	c.JSON(http.StatusOK, data)
}


func main() {

	data = make(map[string]string)

	router := gin.Default()
	router.GET("/ping", Ping)
	router.POST("/data/:key", PostData)
	router.DELETE("/data/:key", DeleteData)
	router.GET("/data/:key", GetData)
	router.PUT("/data/:key", PutData)
	router.GET("/dump", DumpData)
	router.Run(":3000")
}