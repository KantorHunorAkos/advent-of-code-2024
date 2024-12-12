package main

import "testing"

func TestPart1(t *testing.T) {
	for _, testData := range getTestTable() {
		data := getInputFromFile(testData.filename)
		got := part1(data)

		if got != testData.part1Answer {
			t.Errorf("Got: %d want: %d", got, testData.part1Answer)
		}
	}
}

func TestPart2(t *testing.T) {
	for _, testData := range getTestTable() {
		data := getInputFromFile(testData.filename)
		got := part2(data)

		if got != testData.part2Answer {
			t.Errorf("Got: %d want: %d", got, testData.part2Answer)
		}
	}
}

type TestData struct {
	filename                 string
	part1Answer, part2Answer int
}

func getTestTable() []TestData {
	return []TestData{
		{"test_data_1.input", 140, 80},
		{"test_data_2.input", 1930, 1206},
		{"test_data_3.input", 772, 436},
		{"test_data_4.input", 2566, 946},
		{"test_data_5.input", 692, 236},
		{"test_data_6.input", 1184, 368},
	}
}
