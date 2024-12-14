package main

import "testing"

func TestPart1(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part1(data)
	want := 480

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part2(data)
	want := 875318608908

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
