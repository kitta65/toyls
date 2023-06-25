package main

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#headerPart
type header struct {
	ContentLength int
	ContentType   string
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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#notificationMessage
type notification struct {
	message
	Method string `json:"method"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseMessage
type response struct {
	message
	Id     interface{} `json:"id"` // integer | string
	Result interface{} `json:"result"`
}
