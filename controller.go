package main

import (
	"encoding/json"
	"io"
)

//////////////////////////////////////
//Struct de estruturação para o json//
//////////////////////////////////////

type Operation struct {
	Value1   float32 `json:"value1"`
	Value2   float32 `json:"value2"`
	Operator string  `json:"oprt"`
}

/////////////////////////////////////////////////
//mapa de eliminação de if para busca de função//
/////////////////////////////////////////////////

var mapa = map[string]func(float32, float32)float32{
	"addition": addition,
	"multiply": multiply,
	"subtraction":subtraction,
	"divide":divide,
}

///////////////////////////
//Operacoes com os inputs//
///////////////////////////

func multiply(a float32, b float32) float32 {
	return a * b
}

func subtraction(a float32, b float32) float32 {
	return a - b
}

func addition(a float32, b float32) float32 {
	return a + b
}

func divide(a float32, b float32) float32 {
	return a / b
}

////////////////////////////////////////////////////////
//funcao para procura e utilizacao da operacao correta//
////////////////////////////////////////////////////////

func Calculation(body io.ReadCloser) float32 {
	var data Operation
	var result float32
	json.NewDecoder(body).Decode(&data)
	result = mapa[data.Operator] (data.Value1, data.Value2)
	return result

}



