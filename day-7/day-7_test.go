package main

import "testing"

func TestPart1(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part1(data)
	want := 3749

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part2(data)
	want := 11387

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
