package main

import (
	"fmt"
	"log"
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#headerPart
type header struct {
	ContentLength int
	ContentType   string
}

func (h header) toString() string {
	length := fmt.Sprintf("Content-Length: %v\r\n", h.ContentLength)
	type_ := h.ContentType
	if type_ == "" {
		type_ = "application/vscode-jsonrpc; charset=utf-8" // default
	}
	type_ = fmt.Sprintf("Content-Type: %v\r\n", type_)
	return length + type_
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#abstractMessage
type message struct {
	Jsonrpc string `json:"jsonrpc"`
}

type jsonrpc string

func (j jsonrpc) MarshalJSON() ([]byte, error) {
	if j != "" && j != "2.0" {
		log.Fatal("unexpected version")
	}
	return []byte(j), nil
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#requestMessage
type request struct {
	message
	Id     interface{} `json:"id"` // integer | string
	Method string      `json:"method"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseMessage
type response struct {
	message
	Id     interface{} `json:"id"` // integer | string
	Result interface{} `json:"result"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
type initializeRequest struct {
	request
	Params struct{} `json:"params"`
}

type initializeResponse struct {
	response
	Result struct {
		Capabilities struct {
			CompletionProvider struct{} `json:"completionProvider"`
		} `json:"capabilities"`
	} `json:"result"`
}
