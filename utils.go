package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func respond(b []byte) {
	h := header{
		ContentLength: len(b),
	}.toString()
	j := string(b)
	log.Println("SERVER:", j)
	fmt.Printf("%v\r\n%v", h, j)
}

func validate(text string) []diagnostic {
	var diagnostics = []diagnostic{}
	ranges := search(`(?i)voldemort`, text)
	for _, r := range ranges {
		diagnostics = append(diagnostics, diagnostic{
			Range:   r,
			Message: "Do not call his name!",
		})
	}
	return diagnostics
}

func search(exp string, text string) []range_ {
	re := regexp.MustCompile(exp)
	matches := re.FindAllIndex([]byte(text), -1)
	var ranges []range_
	for _, m := range matches {
		start, err := idx2pos(m[0], text)
		if err != nil {
			log.Fatal(err)
		}
		end, err := idx2pos(m[1], text)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, range_{start, end})
	}
	return ranges
}

func idx2pos(idx int, text string) (position, error) {
	lines := strings.Split(text, "\n")
	curr := 0
	for i, line := range lines {
		if len(line)+curr < idx {
			curr += len(line) + 1
			continue
		}
		return position{i, idx - curr}, nil
	}
	return position{}, fmt.Errorf(
		"cannot convert index into position. idx: %v, text length: %v",
		idx,
		len(text),
	)
}

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
