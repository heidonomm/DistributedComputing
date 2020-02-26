package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HotelPort = 3010
const BandPort = 3020

func main() {
	response := getFreeSlots(3020)
	var url string
	err := xml.Unmarshal(response, &url)
	if err != nil {
		fmt.Println("esimene")
		panic(err)
	}

	query := fmt.Sprintf("%s?username=%s&password=%s", url, Username, Password)
	fmt.Println(query)
	time.Sleep(20000 * time.Millisecond)
	resp, err := http.Get(query)
	if err != nil {
		fmt.Println("teine")
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("kolmas")
		panic(err)
	}

	var availability ResponseAvailability
	xml.Unmarshal(body, &availability)
	fmt.Println(availability)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}
