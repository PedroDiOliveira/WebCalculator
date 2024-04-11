package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

////////////////////////////////////////////////////////////////////////
//Struct router que direciona o usuario a partir dos paths no servidor//
////////////////////////////////////////////////////////////////////////

type Router struct{}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if path != "/calculator" {
		resp := messages[404]
		send, _ := json.Marshal(resp)
		res.Write(send)
		return
	}

	if path == "/calculator" && req.Method == "GET" {
		res.Write([]byte("DEU BOM!"))
		return
	}

	if path == "/calculator" && req.Method == "POST" {
		result := Calculation(req.Body)
		send, _ := json.Marshal(result)
		res.Write(send)
		return
	}

}

/////////////////////////////////////
//Função de rosteamento do servidor//
/////////////////////////////////////

func RunService() {
	s := http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      Router{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe()) //log fatal para formatar um possivel erro na função
}
