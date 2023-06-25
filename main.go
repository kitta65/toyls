package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
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
		var reqOrNotif request
		if err := json.Unmarshal(b, &reqOrNotif); err != nil {
			log.Fatal(err)
		}
		log.Println("CLIENT:", string(b))
	}
}
