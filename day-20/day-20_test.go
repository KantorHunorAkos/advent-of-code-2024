package main

import "testing"

func TestPart1(t *testing.T) {
	data := getInputFromFile("test_data.input")
	for _, test := range getTestTablePart1() {

		got := part1(data, test.minSave)
		want := test.answer

		if got != want {
			t.Errorf("Got: %d want: %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data.input")
	for _, test := range getTestTablePart2() {

		got := part2(data, test.minSave)
		want := test.answer

		if got != want {
			t.Errorf("Got: %d want: %d", got, want)
		}
	}
}

func getTestTablePart1() []struct {
	minSave,
	answer int
} {
	return []struct {
		minSave,
		answer int
	}{
		{64, 1},
		{40, 2},
		{38, 3},
		{36, 4},
		{20, 5},
		{12, 8},
		{10, 10},
		{8, 14},
		{6, 16},
		{4, 30},
		{2, 44},
	}
}

func getTestTablePart2() []struct {
	minSave,
	answer int
} {
	return []struct {
		minSave,
		answer int
	}{
		// {50, 285},
		// {52, 253},
		// {54, 222},
		// {56, 193},
		// {58, 154},
		// {60, 129},
		// {62, 106},
		// {64, 86},
		// {66, 67},
		// {68, 55},
		// {70, 41},
		// {72, 29},
		// {74, 7},
		{76, 3},
	}
}
