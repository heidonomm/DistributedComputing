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
	// // freeSlots := availability.Body
	// bestSlot := availability.Body[0]
	// fmt.Println(bestSlot)

	// responseBytes := reserveSlot(HotelPort, bestSlot)
	// fmt.Println(string(responseBytes))

	// bookingsBytes := getBookings(HotelPort)
	// fmt.Println(string(bookingsBytes))
}
