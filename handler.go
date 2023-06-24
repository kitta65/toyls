package main

import (
  "encoding/json"
  "log"
)

func handleInitialize(req initializeRequest) {
	resp := initializeResponse{response: response{Id: req.Id}}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	respond(b)
}

