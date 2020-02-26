package main

import (
	"encoding/xml"
)

type RequestReserve struct {
	XMLName   xml.Name `xml:"reserve"`
	RequestID int      `xml:"request_id"`
	Username  string   `xml:"username"`
	Password  string   `xml:"password"`
	SlotID    int      `xml:"slot_id"`
}

type Response struct {
	XMLName xml.Name `xml:"response"`
	Code    int      `xml:"code"`
	Body    string   `xml:"body"`
}

type RequestCancel struct {
	XMLName   xml.Name `xml:"cancel"`
	RequestID int      `xml:"request_id"`
	Username  string   `xml:"username"`
	Password  string   `xml:"password"`
	SlotID    int      `xml:"slot_id"`
}

type RequestAvailability struct {
	XMLName   xml.Name `xml:"availability"`
	RequestID int      `xml:"request_id"`
	Username  string   `xml:"username"`
	Password  string   `xml:"password"`
}

type ResponseAvailability struct {
	XMLName xml.Name `xml:"response"`
	Code    int      `xml:"code"`
	Body    []int    `xml:"body>availability>slot_id"`
}

type RequestBookings struct {
	XMLName   xml.Name `xml:"bookings"`
	RequestID int      `xml:"request_id"`
	Username  string   `xml:"username"`
	Password  string   `xml:"password"`
}

type ResponseBookings struct {
	XMLName xml.Name `xml:"response"`
	Code    int      `xml:"code"`
	SlotID  []int    `xml:"body>bookings>slot_id"`
}
