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
	var ver2 jsonrpc = "2.0"
	if j != "" && j != ver2 {
		log.Fatal("unexpected version")
	}
	return []byte(ver2), nil
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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion
type completionRequest struct {
	request
	Params struct {
		textDocument string
		Position     struct {
			Line      int `json:"line"`
			Character int `json:"character"`
		} `json:"position"`
	} `json:"params"`
}

type completionResponse struct {
	response
	Result []completionItem `json:"result"`
}

type completionItem struct {
	Label string `json:"label"`
}
