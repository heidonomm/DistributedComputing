package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const HotelPort = 3010
const BandPort = 3020

type funcGetAll func() []byte
type funcPutRequest func(slotID int) []byte

func RequestSlot(port int, slotID int, putRequest funcPutRequest) {

	body := putRequest(slotID)
	url := fmt.Sprintf("http://jewel.cs.man.ac.uk:%d/queue/enqueue", port)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/xml")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bodyRes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(bodyRes))
}

func RequestFreeSlots(port int, getAll funcGetAll) []byte {

	body := getAll()
	url := fmt.Sprintf("http://jewel.cs.man.ac.uk:%d/queue/enqueue", port)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/xml")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bodyRes, _ := ioutil.ReadAll(resp.Body)
	return bodyRes
}
