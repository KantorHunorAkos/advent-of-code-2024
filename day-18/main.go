package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

func main() {
	data := getInputFromFile("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	solution1 := part1(data, 71, 1024)
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution1)

	fmt.Println("Running part2")
	start = time.Now()
	solution2 := part2(data, 71, 1024)
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution2)
}

func getInputFromFile(filename string) []Point {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := []Point{}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		y, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Coundn't parse number with error: %q\n", err)
			return nil
		}
		x, err := strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Printf("Coundn't parse number with error: %q\n", err)
			return nil
		}
		data = append(data, Point{x, y})
	}

	return data
}
