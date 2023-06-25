package main

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func parseHeader(reader io.Reader) header {
	r := bufio.NewReader(reader)
	var h header
	for {
		s, err := r.ReadString('\n')
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
