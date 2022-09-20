package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var data map[string]string

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{})
}

func DumpData(c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

func PostData(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	_, present := data[key]
	if present {
		c.Status(http.StatusBadRequest)
		return
	}

	data[key] = value
	c.JSON(http.StatusCreated, map[string]string{"value": value})
}

func DeleteData(c *gin.Context) {
	//key := c.Param("key")

}

func GetData(c *gin.Context) {
	key := c.Param("key")

	value, present := data[key]
	if !present {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"value": value})
}

func PutData(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	_, present := data[key]
	if !present {
		c.Status(http.StatusNotFound)
		return
	}

	data[key] = value
	c.JSON(http.StatusOK, map[string]string{"value": value})
}

func main() {

	data = make(map[string]string)

	router := gin.Default()
	router.GET("/ping", Ping)
	router.GET("/dump", DumpData)
	router.POST("/data/:key", PostData)
	router.DELETE("/data/:key", DeleteData)
	router.GET("/data/:key", GetData)
	router.PUT("/data/:key", PutData)
	router.Run(":3000")
}