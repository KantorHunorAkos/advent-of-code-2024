package main

import (
	"slices"
	"testing"
)

func TestPart1(t *testing.T) {
	data := getInputFromFile("test_data_1.input")
	got := part1(data)
	want := []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}

	if slices.Compare(got, want) != 0 {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data_2.input")
	got := part2(data)
	want := 117440

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}
