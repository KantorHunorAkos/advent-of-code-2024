package main

import "testing"

func TestPart1(t *testing.T) {
	ids1, ids2 := getIds("test_data.input")
	got := part1(ids1, ids2)
	want := 11

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	ids1, ids2 := getIds("test_data.input")
	got := part2(ids1, ids2)
	want := 31

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
