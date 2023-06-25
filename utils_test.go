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
			"only Content-Lenght",
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

func Test_search(t *testing.T) {
	data := []struct {
		name     string
		idx      int
		text     string
		expected position
	}{
		{
			"single line",
			9,
			"0123456789",
			position{0, 9},
		},
		{
			"multiple lines",
			10,
			"012345678\n" +
				"0123456789",
			position{1, 0},
		},
		{
			"empty text",
			0,
			"",
			position{0, 0},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			actual, err := idx2pos(d.idx, d.text)
			if err != nil {
				t.Error(err)
			} else if d.expected != actual {
				t.Errorf("expected: %v\nactual: %v", d.expected, actual)
			}
		})
	}
}
