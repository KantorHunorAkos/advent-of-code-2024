package main

import "testing"

func TestPart1(t *testing.T) {
	for _, testData := range getTestTable() {
		data := getInputFromFile(testData.filename)
		got := part1(data)

		if got != testData.answer {
			t.Errorf("Got: %d want: %d", got, testData.answer)
		}
	}
}

func TestPart2(t *testing.T) {
	data := getInputFromFile("test_data_3.input")
	got := part2(data)
	want := 9021

	if got != want {
		t.Errorf("Got: %d want: %d", got, want)
	}
}

type testData struct {
	filename string
	answer   int
}

func getTestTable() []testData {
	return []testData{
		{"test_data_1.input", 10092},
		{"test_data_2.input", 2028},
	}
}
