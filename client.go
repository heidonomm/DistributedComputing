package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func reserveSlot(port int, slotID int) []byte {
	return RequestSingle(port, slotID, buildRequestReserve)
}

func cancelSlot(port int, slotID int) []byte {
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

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/xml")

	var resp *http.Response
	resp, _ = client.Do(req)
	for resp.StatusCode >= 400 {
		resp, _ = client.Do(req)
	}

	bodyRes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// bodyRes := repeatedExecutionWithReqObj(req)

	var resUrl string
	if err := xml.Unmarshal(bodyRes, &resUrl); err != nil {
		panic(err)
	}

	getQueryUrl := fmt.Sprintf("%s?username=%s&password=%s", resUrl, Username, Password)
	fmt.Println(getQueryUrl)
	//get the response

	return repeatedExecution(getQueryUrl)
}

func RequestMultiple(port int, getAll funcGetAll) []byte {

	body := getAll()
	url := fmt.Sprintf("http://jewel.cs.man.ac.uk:%d/queue/enqueue", port)

	// client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/xml")

	var resp *http.Response
	resp, _ = client.Do(req)
	for resp.StatusCode >= 400 {
		resp, _ = client.Do(req)
	}

	bodyRes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// bodyRes := repeatedExecutionWithReqObj(req)

	//parse the url to make the request to (url where server response is hidden)
	var resUrl string
	if err := xml.Unmarshal(bodyRes, &resUrl); err != nil {
		panic(err)
	}
	getQueryUrl := fmt.Sprintf("%s?username=%s&password=%s", resUrl, Username, Password)
	fmt.Println(getQueryUrl)
	//get the response

	return repeatedExecution(getQueryUrl)
}

func repeatedExecution(queryUrl string) []byte {

	var respMult *http.Response

	respMult, _ = http.Get(queryUrl)

	counter := 1
	var respBody []byte

	for respMult.StatusCode != 200 {
		respMult, _ = http.Get(queryUrl)

		time.Sleep(time.Second)
		fmt.Printf("\rTime failed: %d, statusCode: %d,  \n", counter, respMult.StatusCode)
		counter++
	}

	b, err := ioutil.ReadAll(respMult.Body)
	if err != nil {
		panic(err)
	}
	respBody = b
	fmt.Println(string(b))

	return respBody
}

func repeatedExecutionWithReqObj(req *http.Request) []byte {
	client := &http.Client{}

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/xml")

	var resp *http.Response
	resp, _ = client.Do(req)
	for resp.StatusCode >= 400 {
		resp, _ = client.Do(req)
	}

	bodyRes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return bodyRes
}
