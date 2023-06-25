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
	// always return "2.0"
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

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
type initializeRequest struct {
	request
	Params struct{} `json:"params"`
}

type initializeResponse struct {
	response
	Result initializeResult `json:"result"`
}

type initializeResult struct {
	Capabilities serverCapabilities `json:"capabilities"`
}

type serverCapabilities struct {
	CompletionProvider struct{}               `json:"completionProvider"`
	TextDocumentSync   textDocumentSyncOption `json:"textDocumentSync"`
}

type textDocumentSyncOption struct {
	OpenClose bool `json:"openClose"`
	Change    int  `json:"change"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion
type completionRequest struct {
	request
	Params struct {
		TextDocument struct {
			Uri string `json:"uri"`
		} `json:"textDocument"`
		Position position `json:"position"`
	} `json:"params"`
}

type completionResponse struct {
	response
	Result []completionItem `json:"result"`
}

type completionItem struct {
	Label string `json:"label"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didClose
type didCloseNotification struct {
	notification
	Params struct {
		TextDocument struct {
			Uri string `json:"uri"`
		} `json:"textDocument"`
	} `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didChange
type didChangeNotification struct {
	notification
	Params struct {
		TextDocument struct {
			Version int    `json:"version"`
			Uri     string `json:"uri"`
		} `json:"textDocument"`
		ContentChanges []struct {
			Text string `json:"text"`
		} `json:"contentChanges"`
	} `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didOpen
type didOpenNotification struct {
	notification
	Params struct {
		TextDocument struct {
			Uri        string `json:"uri"`
			LanguageId string `json:"languageId"`
			Version    int    `json:"version"`
			Text       string `json:"text"`
		} `json:"textDocument"`
	} `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_publishDiagnostics
type publishDiagnosticsNotification struct {
	notification
	Params publishDiagnosticsParams `json:"params"`
}

type publishDiagnosticsParams struct {
	Uri         string       `json:"uri"`
	Diagnostics []diagnostic `json:"diagnostics"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnostic
type diagnostic struct {
	Range   range_ `json:"range"`
	Message string `json:"message"`
}

type range_ struct {
	Start position `json:"start"`
	End   position `json:"end"`
}

type position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
