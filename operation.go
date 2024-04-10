package main

import (
	"encoding/json"
	"io"
)

type Operation struct {
	Value1   float32 `json:"value1"`
	Value2   float32 `json:"value2"`
	Operator string  `json:"oprt"`
}

func Calculation(body io.ReadCloser) float32 {
	var data Operation
	var result float32
	json.NewDecoder(body).Decode(&data)

	if data.Operator == "addition" {
		result = addition(data.Value1, data.Value2)
		return result
	}
	if data.Operator == "multiply" {
		result = multiply(data.Value1, data.Value2)
		return result
	}
	if data.Operator == "subtraction" {
		result = subtraction(data.Value1, data.Value2)
		return result
	}
	if data.Operator == "divide" {
		result = divide(data.Value1, data.Value2)
		return result
	}
	return result
}
