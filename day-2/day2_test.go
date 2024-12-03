package main

import "testing"

func TestPart1(t *testing.T) {
	data := getNumbersFromFile("test_data.input")
	got := part1(data)
	want := 2

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := getNumbersFromFile("test_data.input")
	got := part2(data)
	want := 4

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
