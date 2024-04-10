package main

type HTTPMessage struct {
	Code        int    `json:"code"`
	Description string `json:"desc"`
}

var messages = map[int]HTTPMessage{
	404: {Code: 404, Description: "Not found"},
	200: {Code: 200, Description: "OK"},
}
