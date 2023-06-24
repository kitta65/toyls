package main

import (
	"encoding/json"
	"log"
)

func handleCompletion(req completionRequest) {
	resp := completionResponse{
		response: response{Id: req.Id},
		Result: []completionItem{
			{Label: "TypeScript"},
			{Label: "JavaScript"},
		},
	}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	respond(b)
}

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
