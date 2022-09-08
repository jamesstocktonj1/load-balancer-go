package main

import (
	"net/http"
	"sync"
	"github.com/gin-gonic/gin"
)

type node struct {
	address		string
	port		string
	keys		[]string
}

var nodes = []node {
	node{ "10.5.0.11", "3000", []string{} },
	node{ "10.5.0.12", "3000", []string{} },
}



func (n node) ContainesKey(key string) bool {

	for _, k := range n.keys {

		if k == key {
			return true
		}
	}
	return false
}

func KeyIndex(key string) int {

	for i, n := range nodes {

		if n.ContainesKey(key) {
			return i
		}
	}
	return -1
}


func Ping(c *gin.Context) {

	var statusArray = map[string]string {}
	var pingSync sync.WaitGroup

	for _, n := range nodes {

		pingSync.Add(1)
		go func(n node) {

			resp, err := http.Get("http://" + n.address + ":" + n.port + "/ping")

			if err != nil {
				statusArray[n.address + ":" + n.port] = "Offline"
			} else if resp.StatusCode == 200 {
				statusArray[n.address + ":" + n.port] = "Online"
			} else {
				statusArray[n.address + ":" + n.port] = "Error"
			}
			pingSync.Done()
		}(n)
	}
	pingSync.Wait()

	c.JSON(http.StatusOK, statusArray)
}

func DumpData(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}


func main() {

	router := gin.Default()
	router.GET("/ping", Ping)
	router.Run(":3000")
}