package main

import (
	"fmt"
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
	Jsonrpc jsonrpc `json:"jsonrpc"`
}

type jsonrpc string

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#requestMessage
type request struct {
	message
	Id     interface{} `json:"id"` // integer | string
	Method string      `json:"method"`
}
