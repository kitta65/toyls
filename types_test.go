package main

import (
	"encoding/json"
	"testing"
)

func Test_message(t *testing.T) {
	msg := message{}

	json_, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(json_, &msg)
	if err != nil {
		t.Error(err)
	}

	if (msg != message{Jsonrpc: "2.0"}) {
		t.Errorf("msg.Jsonrpc must be 2.0. actual: %v", msg.Jsonrpc)
	}
}
