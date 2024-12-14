package main

import "testing"

func TestPart1(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part1(data, &MapSize{11, 7})
	want := 12

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part2(data, &MapSize{11, 7})
	want := 0

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
