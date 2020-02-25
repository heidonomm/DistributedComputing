package main

import (
	"encoding/xml"
	"time"
)

func getUniqueId() int {
	return int(time.Now().UnixNano())
}

func buildRequestReserve(slotID int) []byte {
	reqBody := &RequestReserve{
		RequestID: getUniqueId(),
		Username:  Username,
		Password:  Password,
		SlotID:    slotID,
	}

	body, err := xml.MarshalIndent(reqBody, "", " ")
	if err != nil {
		panic(err)
	}

	return body
}

func buildRequestCancel(slotID int) []byte {
	reqBody := &RequestCancel{
		RequestID: getUniqueId(),
		Username:  Username,
		Password:  Password,
		SlotID:    slotID,
	}

	body, err := xml.MarshalIndent(reqBody, "", " ")
	if err != nil {
		panic(err)
	}

	return body
}

func buildRequestAvailability() []byte {
	reqBody := &RequestAvailability{
		RequestID: getUniqueId(),
		Username:  Username,
		Password:  Password,
	}

	body, err := xml.MarshalIndent(reqBody, "", " ")
	if err != nil {
		panic(err)
	}

	return body
}

func buildRequestBookings() []byte {
	reqBody := &RequestBookings{
		RequestID: getUniqueId(),
		Username:  Username,
		Password:  Password,
	}

	body, err := xml.MarshalIndent(reqBody, "", " ")
	if err != nil {
		panic(err)
	}

	return body
}
