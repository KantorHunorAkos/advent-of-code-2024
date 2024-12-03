package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part2(data string) int {
	mulRegex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create regular expression with error: %s", err)
		return 0
	}

	sum := 0
	enabled := true

	matches := mulRegex.FindAllString(data, -1)
	for _, match := range matches {

		switch match {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
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
	}
	}

	return sum
}