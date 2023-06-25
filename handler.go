package main

import (
	"encoding/json"
	"log"
)

func handleInitialize(req initializeRequest) {
	resp := initializeResponse{response: response{Id: req.Id}, Result: initializeResult{
		Capabilities: serverCapabilities{
			TextDocumentSync: textDocumentSyncOption{OpenClose: true, Change: 1},
		},
	}}
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
