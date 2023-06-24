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

	// handle request
	r := bufio.NewReader(os.Stdin)
	var b []byte
	for {
		h := parseHeader(r)
		if cap(b) < h.ContentLength {
			b = make([]byte, h.ContentLength)
		} else {
			b = b[:h.ContentLength]
		}

		_, err := io.ReadFull(r, b)
		var req request
		if err := json.Unmarshal(b, &req); err != nil {
			log.Fatal(err)
		}

		switch method := req.Method; method {
		case "initialize":
			log.Println(string(b))
			var req initializeRequest
			if err := json.Unmarshal(b, &req); err != nil {
				log.Fatal(err)
			}
			handleInitialize(req)
		default:
			// method not impremented
			log.Println(string(b))
		}
		if err == io.EOF {
			break // connection closed
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
