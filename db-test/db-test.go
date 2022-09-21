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
	dataCount = 800
)

func randomString(l int) string {

	b := make([]byte, l)

	for i := 0; i < l; i++ {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func formatResponseTime(times []int64) {

	var timeSum int64 = 0
	timeLow := times[0]
	timeHigh := times[0]

	for _, t := range times {
		timeSum += t

		if t > timeHigh {
			timeHigh = t
		}

		if t < timeLow {
			timeLow = t
		}
	}

	timeAverage := timeSum / int64(len(times))

	fmt.Printf("Minimum Response: %f ms\n", float32(timeLow) / 1000)
	fmt.Printf("Maximum Response: %f ms\n", float32(timeHigh) / 1000)
	fmt.Printf("Average Response: %f ms\n\n", float32(timeAverage) / 1000)
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
	timePeriods := []int64 {}
	for testKey, _ := range testData {

		startTime := time.Now()
		err = n.CreateValue(testKey, "")
		endTime := time.Now()

		timePeriods = append(timePeriods, endTime.Sub(startTime).Microseconds())

		if err != nil {
			errorLog.Println(err)
		}
	}
	formatResponseTime(timePeriods) 


	fmt.Println("Test Get Value...")
	timePeriods = []int64 {}
	for testKey, _ := range testData {

		startTime := time.Now()
		_, err = n.GetValue(testKey)
		endTime := time.Now()

		timePeriods = append(timePeriods, endTime.Sub(startTime).Microseconds())
		
		if err != nil {
			errorLog.Println(err)
		}

		time.Sleep(time.Millisecond)
	}
	formatResponseTime(timePeriods) 


	fmt.Println("Test Set Value")
	timePeriods = []int64 {}
	for testKey, testValue := range testData {

		startTime := time.Now()
		err = n.SetValue(testKey, testValue)
		endTime := time.Now()

		timePeriods = append(timePeriods, endTime.Sub(startTime).Microseconds())

		if err != nil {
			errorLog.Println(err)
		}

		time.Sleep(time.Millisecond)
	}
	formatResponseTime(timePeriods) 


	fmt.Println("Test Verify Value...")
	timePeriods = []int64 {}
	for testKey, testValue := range testData {

		startTime := time.Now()
		respValue, err := n.GetValue(testKey)
		endTime := time.Now()

		timePeriods = append(timePeriods, endTime.Sub(startTime).Microseconds())

		if err != nil {
			errorLog.Println(err)
		} else if respValue != testValue {
			errorLog.Printf("Error: mismatch in values, expected %s but got %s\n", testValue, respValue)
		}

		time.Sleep(time.Millisecond)
	}
	formatResponseTime(timePeriods) 
}