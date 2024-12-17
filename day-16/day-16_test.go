package main

import "testing"

func TestPart1(t *testing.T) {
	for _, testData := range getTestTable() {
		data := getInputFromFile(testData.filename)
		got := part1(data)

		if got != testData.answerPart1 {
			t.Errorf("Got: %d want: %d", got, testData.answerPart1)
		}
	}
}

func TestPart2(t *testing.T) {
	for _, testData := range getTestTable() {
		data := getInputFromFile(testData.filename)
		got := part2(data)

		if got != testData.answerPart2 {
			t.Errorf("Got: %d want: %d", got, testData.answerPart2)
		}
	}
}

type testData struct {
	filename    string
	answerPart1 int
	answerPart2 int
}

func getTestTable() []testData {
	return []testData{
		{"test_data_1.input", 7036, 45},
		{"test_data_2.input", 11048, 64},
	}
}
