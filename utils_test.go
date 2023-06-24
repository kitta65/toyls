package main

import (
	"strings"
	"testing"
)

func Test_parseHeader(t *testing.T) {
	data := []struct {
		name     string
		input    string
		expected header
	}{
		{
			"simple",
			"Content-Length: 300\r\n" +
				"\r\n" +
				"{}",
			header{
				ContentLength: 300,
				ContentType:   "",
			},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			actual := parseHeader(strings.NewReader(d.input))
			if d.expected != actual {
				t.Errorf("expected: %v\nactual: %v", d.expected, actual)
			}
		})
	}
}
