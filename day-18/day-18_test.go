package main

import "testing"

func TestPart1(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part1(data, 7, 12)
	want := 22

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data.input")
	got := part2(data, 7, 12)
	want := Point{1,6}

	if !got.equal(want) {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
