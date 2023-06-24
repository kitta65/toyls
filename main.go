package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log" // NOTE try slog when v1.21 is released
	"os"
)

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
		var req request
		if err := json.Unmarshal(b, &req); err != nil {
			log.Fatal(err)
		}
		log.Println("CLIENT:", string(b))

		// respond to message
		method := req.Method
		switch method {
		case "initialize":
			var req initializeRequest
			if err := json.Unmarshal(b, &req); err != nil {
				log.Fatal(err)
			}
			handleInitialize(req)
		case "shutdown":
			handleShutdown(req)
			// TODO return error for all requests except for exit
		case "textDocument/completion":
			var req completionRequest
			if err := json.Unmarshal(b, &req); err != nil {
				log.Fatal(err)
			}
			handleCompletion(req)
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
