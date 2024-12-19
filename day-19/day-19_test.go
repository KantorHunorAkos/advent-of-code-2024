package main

import "testing"

func TestPart1(t *testing.T) {
	towels, arrangements := getInputFromFile("test_data.input")
	got := part1(towels, arrangements)
	want := 6

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	towels, arrangements := getInputFromFile("test_data.input")
	got := part2(towels, arrangements)
	want := 16

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
