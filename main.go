package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

var texts = map[string]string{}

func main() {
	// configure log
	f, err := os.Create("/tmp/toyls.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // this may return error so it is a little optimistic
	log.SetOutput(f)

	// handle message
	r := bufio.NewReader(os.Stdin)
	var b []byte
	for {
		h := parseHeader(r)
		if cap(b) < h.ContentLength {
			b = make([]byte, h.ContentLength)
		} else {
			b = b[:h.ContentLength]
		}

		// read json
		_, err := io.ReadFull(r, b)
		var reqOrNotif request
		if err := json.Unmarshal(b, &reqOrNotif); err != nil {
			log.Fatal(err)
		}
		log.Println("CLIENT:", string(b))

		// check method
		method := reqOrNotif.Method
		switch method {
		case "initialize":
			var req initializeRequest
			if err := json.Unmarshal(b, &req); err != nil {
				log.Fatal(err)
			}
			handleInitialize(req)
		case "shutdown":
			handleShutdown(reqOrNotif)
		case "textDocument/completion":
			var req completionRequest
			if err := json.Unmarshal(b, &req); err != nil {
				log.Fatal(err)
			}
			handleCompletion(req)
		case "textDocument/didClose":
			var notif didCloseNotification
			if err := json.Unmarshal(b, &notif); err != nil {
				log.Fatal(err)
			}
			handleDidClose(notif)
		case "textDocument/didChange":
			var notif didChangeNotification
			if err := json.Unmarshal(b, &notif); err != nil {
				log.Fatal(err)
			}
			handleDidChange(notif)
		case "textDocument/didOpen":
			var notif didOpenNotification
			if err := json.Unmarshal(b, &notif); err != nil {
				log.Fatal(err)
			}
			handleDidOpen(notif)
		}

		// check status
		if err == io.EOF {
			break // connection may be closed
		} else if err != nil {
			log.Fatal(err)
		}
		if method == "exit" {
			break
		}
	}
}
