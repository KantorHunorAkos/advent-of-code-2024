package main

import "testing"

func TestPart1(t *testing.T) {
	data := getInputFromFile("test_data_part1.input")
	got := part1(data)
	want := 161

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data_part2.input")
	got := part2(data)
	want := 48

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
