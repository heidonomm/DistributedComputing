package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReserveSlot(port int, slotID int) []byte {
	return RequestSingle(port, slotID, buildRequestReserve)
}

func CancelSlot(port int, slotID int) []byte {
	return RequestSingle(port, slotID, buildRequestCancel)
}

func getFreeSlots(port int) []byte {
	return RequestMultiple(port, buildRequestAvailability)
}

func getBookings(port int) []byte {
	return RequestMultiple(port, buildRequestBookings)
}

type funcGetAll func() []byte
type funcGetSingle func(slotID int) []byte

func RequestSingle(port int, slotID int, putRequest funcGetSingle) []byte {

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
	resp.Body.Close()

	return bodyRes
}

func RequestMultiple(port int, getAll funcGetAll) []byte {

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
	resp.Body.Close()

	return bodyRes
}
