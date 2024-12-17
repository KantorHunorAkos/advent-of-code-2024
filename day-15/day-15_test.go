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
		{"test_data_1.input", 10092, 9021},
		{"test_data_2.input", 2028, 1751},
		{"test_data_3.input", 908, 618},
		{"test_data_4.input", 1505, 1211},
		{"test_data_5.input", 1206, 1213},
		{"test_data_6.input", 1816, 1430},
		{"test_data_7.input", 3012, 2827},
	}
}
