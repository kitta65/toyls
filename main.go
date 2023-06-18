package main

import (
	"bufio"
	"io"
	"log" // NOTE try slog when v1.21 is released
	"os"
	"strconv"
	"strings"
)

func main() {
	// configure log
	f, err := os.Create("/tmp/toyls.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // this may return error so it is a little optimistic
	log.SetOutput(f)

	// handle messages
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
		log.Println(string(b))
		if err == io.EOF {
			break // connection closed
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func parseHeader(reader io.Reader) header {
	r := bufio.NewReader(reader)
	var h header
	for {
		s, err := r.ReadString('\n')
		log.Println(s)
		if s == "\r\n" { // empty line
			return h
		}
		kv := strings.SplitN(s, ":", 2)
		k := kv[0]
		v := strings.TrimSpace(kv[1])
		switch k {
		case "Content-Length":
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			h.ContentLength = i
		case "Content-Type":
			h.ContentType = v
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return h
}
