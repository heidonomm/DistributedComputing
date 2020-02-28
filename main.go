package main

import (
	"encoding/xml"
	"fmt"
)

const HotelPort = 3010
const BandPort = 3020

const (
	waitforServer   = iota
	reqAvailability = iota
	reqSlot         = iota
	reqBookings     = iota
)

func main() {
	var availability ResponseAvailability
	availabilityBytes := getFreeSlots(HotelPort)
	xml.Unmarshal(availabilityBytes, &availability)
	fmt.Println(availability)
	// freeSlots := availability.Body
	bestSlot := availability.Body[0]
	fmt.Println(bestSlot)

	responseBytes := reserveSlot(HotelPort, bestSlot)
	fmt.Println(string(responseBytes))

	bookingsBytes := getBookings(HotelPort)
	fmt.Println(string(bookingsBytes))
}

// func main() {
// 	response := getBookings(3020)
// 	var url string
// 	err := xml.Unmarshal(response, &url)
// 	if err != nil {
// 		fmt.Println("esimene")
// 		panic(err)
// 	}

// 	query := fmt.Sprintf("%s?username=%s&password=%s", url, Username, Password)
// 	fmt.Println(query)
// 	time.Sleep(20000 * time.Millisecond)
// 	resp, err := http.Get(query)
// 	if err != nil {
// 		fmt.Println("teine")
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("kolmas")
// 		panic(err)
// 	}

// 	var availability ResponseBookings
// 	xml.Unmarshal(body, &availability)
// 	fmt.Println(availability)
// 	fmt.Println(resp.StatusCode)
// 	fmt.Println(string(body))
// }

// response := getBookings(3020)
// var url string
// err := xml.Unmarshal(response, &url)
// if err != nil {
// 	fmt.Println("esimene")
// 	panic(err)
// }
// state := make(chan int)
// select {
// 	tocker
// }
// for {
// 	select <-ticker {
// 		switch <-state {

// 		}
// 	}

// }

// ticker := time.NewTicker(500 * time.Millisecond)
// done := make(chan bool)

// go func() {
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case t := <-ticker.C:
// 			fmt.Println("Tick at", t)
// 		}
// 	}
// }()

// time.Sleep(1600 * time.Millisecond)
// ticker.Stop()
// done <- true
// fmt.Println("Ticker stopped")
