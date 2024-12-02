package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	numbers := getNumbersFromFile("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	part1(numbers)
	duration := time.Until(start).Abs()
	fmt.Printf("part1 ran succesfully\ntime: %s\n", duration.String())

	fmt.Println("Running part2")
	start = time.Now()
	part2(numbers)
	duration = time.Until(start).Abs()
	fmt.Printf("part2 ran succesfully\ntime: %s\n", duration.String())
}

func getNumbersFromFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file '%s'\n", err)
		return nil
	}
	defer file.Close()

	numbers := [][]int{}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		numberLine := []int{}

		for _, token := range tokens {
			number, err := strconv.Atoi(token)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't convert token %s to number with error %s\n", token, err)
				return nil
			}
			numberLine = append(numberLine, number)
		}

		numbers = append(numbers, numberLine)
	}

	return numbers
}