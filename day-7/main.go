package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Operation struct {
	result   int
	operands []int
}

func main() {
	data := getInputFromFile("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	solution := part1(data)
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)

	fmt.Println("Running part2")
	start = time.Now()
	solution = part2(data)
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)
}

func getInputFromFile(filename string) []Operation {
	data := []Operation{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return data
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ":")
		result, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Coundn't parse %q into result with error: %q\n", tokens[0], err)
			return data
		}
		operands := []int{}
		tokens[1] = strings.TrimSpace(tokens[1])
		for _, token := range strings.Split(tokens[1], " ") {
			operand, err := strconv.Atoi(token)
			if err != nil {
				fmt.Printf("Coundn't parse %q into opernad with error: %q\n", token, err)
				return data
			}
			operands = append(operands, operand)
		}

		data = append(data, Operation{result, operands})
	}

	return data
}
