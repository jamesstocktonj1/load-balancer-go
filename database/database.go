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

	_, present := data[key]
	if present {
		c.Status(http.StatusBadRequest)
	} else {
		data[key] = value
		c.JSON(http.StatusCreated, map[string]string { "value": value })
	}
}

func DeleteData(c *gin.Context) {
	//key := c.Param("key")

}

func GetData(c *gin.Context) {
	key := c.Param("key")

	value, present := data[key]
	if !present {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, map[string]string { "value": value })
	}
}

func PutData(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	_, present := data[key]
	if !present {
		c.Status(http.StatusNotFound)
	} else {
		data[key] = value
		c.JSON(http.StatusOK, map[string]string { "value": value })
	}
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