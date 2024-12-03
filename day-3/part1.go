package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(data string) int {
	mulRegex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create regular expression `mul\\(\\d{1,3},\\d{1,3}\\)` with error: %s", err)
		return 0
	}

	sum := 0

	matches := mulRegex.FindAllString(data, -1)
	for _, match := range matches {
		tokens := strings.Split(match[3:], ",")

		firstNumber, err := strconv.Atoi(tokens[0][1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't convert %s to number with error: %s", tokens[0][1:], err)
			return 0
		}

		secondNumber, err := strconv.Atoi(tokens[1][:len(tokens[1])-1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't convert %s to number with error: %s", tokens[1][:len(tokens[1])-1], err)
			return 0
		}

		sum += firstNumber * secondNumber
	}

	return sum
}
