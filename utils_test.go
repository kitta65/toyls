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
				t.Errorf("expected: %v, actual: %v", d.expected, actual)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	text := `voldemort
Voldemort`
	actual := validate(text)
	expected := []diagnostic{
		diagnostic{Range: range_{position{0, 0}, position{0, 9}}},
		diagnostic{Range: range_{position{1, 0}, position{1, 9}}},
	}
	if len(actual) != len(expected) {
		t.Errorf("expected: %v errors, actual: %v errors", len(expected), len(actual))
	}
	for i := 0; i < len(actual); i++ {
		if expected[i].Range != actual[i].Range {
			t.Errorf("expected: %v, actual: %v", expected[i].Range, actual[i].Range)
		}
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
			3,
			"0123456789",
			position{0, 3},
		},
		{
			"multiple lines",
			10,
			"012345678\n" +
				"0123456789",
			position{1, 0},
		},
		{
			"end of line",
			9,
			"012345678\n",
			position{0, 9},
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
				t.Errorf("expected: %v, actual: %v", d.expected, actual)
			}
		})
	}
}
