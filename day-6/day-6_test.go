package main

import "testing"

func TestPart1(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part1(data)
	want := 41

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part2(data)
	want := 0

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
