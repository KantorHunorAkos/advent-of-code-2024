package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

func getInputFromFile(filename string) *GuideData {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]int{}
	tailheads := []Coord{}
	rowNum := 0

	for scanner.Scan() {
		row := []int{}
		for colNum, char := range scanner.Text() {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Printf("Couldn't parse %q as a digit with error: %q", char, err)
				return nil
			}

			row = append(row, digit)

			if digit == 0 {
				tailheads = append(tailheads, Coord{rowNum, colNum})
			}
		}
		grid = append(grid, row)
		rowNum++
	}

	return &GuideData{grid, tailheads}
}
