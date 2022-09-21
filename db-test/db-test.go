package main

import (
	"db-test/node"
	"log"
	"fmt"
	"os"
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	dataCount = 5000
)

func randomString(l int) string {

	b := make([]byte, l)

	for i := 0; i < l; i++ {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func main() {
	n := node.Node{}
	n.Address = "127.0.0.1"
	n.Port = "3000"

	rand.Seed(time.Now().UnixNano())

	logFile, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	errorLog := log.New(logFile, "", log.Ltime)



	fmt.Println("Starting Test")
	
	
	fmt.Println("Creating Test Data")
	testData := map[string]string {}
	for i:=0; i<dataCount; i++ {
		testKey := randomString(8)
		testValue := randomString(12)
		testData[testKey] = testValue
	}
	fmt.Printf("Data Size: %d\n\n", len(testData))

	
	fmt.Println("Test Create Value...")
	startTime := time.Now()
	for testKey, _ := range testData {

		err = n.CreateValue(testKey, "")
		if err != nil {
			errorLog.Println(err)
		}
	}
	endTime := time.Now()
	t := endTime.Sub(startTime)
	fmt.Printf("Create Value Response Time: %f ms\n\n", float64(t.Microseconds()) / (dataCount * 1000))


	fmt.Println("Test Get Value...")
	startTime = time.Now()
	for testKey, _ := range testData {
		_, err = n.GetValue(testKey)
		if err != nil {
			errorLog.Println(err)
		}

		time.Sleep(time.Millisecond)
	}
	endTime = time.Now()
	t = endTime.Sub(startTime)
	fmt.Printf("Get Value Response Time: %f ms\n\n", float64(t.Microseconds()) / (dataCount * 1000))


	startTime = time.Now()
	fmt.Println("Test Set Value")
	for testKey, testValue := range testData {
		err = n.SetValue(testKey, testValue)
		if err != nil {
			errorLog.Println(err)
		}

		time.Sleep(time.Millisecond)
	}
	endTime = time.Now()
	t = endTime.Sub(startTime)
	fmt.Printf("Set Value Response Time: %f ms\n\n", float64(t.Microseconds()) / (dataCount * 1000))


	startTime = time.Now()
	fmt.Println("Test Verify Value...")
	for testKey, testValue := range testData {
		respValue, err := n.GetValue(testKey)
		if err != nil {
			errorLog.Println(err)
		} else if respValue != testValue {
			errorLog.Printf("Error: mismatch in values, expected %s but got %s\n", testValue, respValue)
		}

		time.Sleep(time.Millisecond)
	}
	endTime = time.Now()
	t = endTime.Sub(startTime)
	fmt.Printf("Get Value Response Time: %f ms\n", float64(t.Microseconds()) / (dataCount * 1000))
}