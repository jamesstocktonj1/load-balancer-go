package main

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	resp, err := http.Get("http://10.5.0.11:3000/ping")
	fmt.Println(resp.Status)

	if err != nil {
		fmt.Println("Error Pinging Database")
	} else {

		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)

		d := map[string]string {}

		json.Unmarshal(data, &d)
		
		fmt.Println(d)
	}
	
	c.JSON(http.StatusOK, map[string]string {})
}


func main() {

	router := gin.Default()
	router.GET("/ping", Ping)
	router.Run(":3000")
}