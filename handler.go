package main

import (
	"encoding/json"
	"log"
)

func handleCompletion(req completionRequest) {
	resp := completionResponse{
		response: response{Id: req.Id},
		Result: []completionItem{
			{Label: "Harry", Detail: "name of a famous wizard"},
		},
	}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	respond(b)
}

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

func handleDidClose(notif didCloseNotification) {
	uri := notif.Params.TextDocument.Uri

	// empty diagnostics
	params := publishDiagnosticsParams{Uri: uri, Diagnostics: []diagnostic{}}
	resp := publishDiagnosticsNotification{
		notification: notification{Method: "textDocument/publishDiagnostics"},
		Params:       params,
	}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	respond(b)
}

func handleDidChange(notif didChangeNotification) {
	uri := notif.Params.TextDocument.Uri
	texts[uri] = notif.Params.ContentChanges[0].Text

	params := publishDiagnosticsParams{Uri: uri, Diagnostics: validate(texts[uri])}
	resp := publishDiagnosticsNotification{
		notification: notification{Method: "textDocument/publishDiagnostics"},
		Params:       params,
	}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	respond(b)
}

func handleDidOpen(notif didOpenNotification) {
	uri := notif.Params.TextDocument.Uri
	texts[uri] = notif.Params.TextDocument.Text

	params := publishDiagnosticsParams{Uri: uri, Diagnostics: validate(texts[uri])}
	resp := publishDiagnosticsNotification{
		notification: notification{Method: "textDocument/publishDiagnostics"},
		Params:       params,
	}
	b, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	respond(b)
}
