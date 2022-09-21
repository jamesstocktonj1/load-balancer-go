package main

import (
	"load-balancer/node"
	"github.com/gin-gonic/gin"
	"net/http"
)

var DatabaseNodes = []node.Node {
	node.Node{ "10.5.0.11", "3000", "Offline", []string{} },
	node.Node{ "10.5.0.12", "3000", "Offline", []string{} },
}

func Ping(c *gin.Context) {
	var statusArray = map[string]string {}

	for _, n := range DatabaseNodes {
		n.Ping()
		statusArray[n.Address + ":" + n.Port] = n.Status
	}
	c.JSON(http.StatusOK, statusArray)
}

func DumpData(c *gin.Context) {
	var dataArray = map[string]([]string) {}

	for _, n := range DatabaseNodes {

		err := n.UpdateKeys()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		
		dataArray[n.Address + ":" + n.Port] = n.Keys
	}
	c.JSON(http.StatusOK, dataArray)
}

func NextNode() int {
	var lowestIndex = 0

	for i, n := range DatabaseNodes {

		if len(n.Keys) < len(DatabaseNodes[lowestIndex].Keys) {
			lowestIndex = i
		}
	}
	return lowestIndex
}

func KeyIndex(key string) int {

	for i, n := range DatabaseNodes {

		if n.ContainesKey(key) {
			return i
		}
	}
	return -1
}


func PostData(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	isPresent := KeyIndex(key)
	if isPresent != -1 {
		c.Status(http.StatusBadRequest)
		return
	}

	newNode := NextNode()
	err := DatabaseNodes[newNode].CreateValue(key, value)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"value": value})
}


func DeleteData(c *gin.Context) {
	//key := c.Param("key")

}

func GetData(c *gin.Context) {
	key := c.Param("key")

	nodeIndex := KeyIndex(key)
	if nodeIndex == -1 {
		c.Status(http.StatusNotFound)
		return
	}

	value, err := DatabaseNodes[nodeIndex].GetValue(key)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"value": value})
}

func PutData(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")

	nodeIndex := KeyIndex(key)
	if nodeIndex == -1 {
		c.Status(http.StatusNotFound)
		return
	}

	err := DatabaseNodes[nodeIndex].SetValue(key, value)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"value": value})
}


func main() {
	
	router := gin.Default()
	router.GET("/ping", Ping)
	router.GET("/dump", DumpData)
	router.POST("/data/:key", PostData)
	//router.DELETE("/data/:key", DeleteData)
	router.GET("/data/:key", GetData)
	router.PUT("/data/:key", PutData)
	router.Run(":3000")
}