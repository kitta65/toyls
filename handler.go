package main

import (
	"encoding/json"
	"log"
)

const (
	jsonrpcVer = "2.0"
)

func handleInitialize(req initializeRequest) {
	resp := initializeResponse{response: response{Id: req.Id}}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	respond(b)
}

func handleShutdown(req request) {
	resp := response{Id: req.Id, Result: nil}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	respond(b)
}
