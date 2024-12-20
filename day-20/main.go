package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

const (
	WALL = math.MinInt
	PATH = 0
)

type Point struct{ x, y int }
type Track [][]int

type Data struct {
	m Track
	s Point
	e Point
}

func main() {
	data := getInputFromFile("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	solution := part1(data, 100)
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)

	fmt.Println("Running part2")
	start = time.Now()
	solution = part2(data, 100)
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)
}

func getInputFromFile(filename string) *Data {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	data := Data{}
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		row := make([]int, len(line))
		for j, c := range line {
			switch c {
			case '#':
				row[j] = WALL
			case '.':
				row[j] = PATH
			case 'E':
				data.e.x = i
				data.e.y = j
				row[j] = PATH
			case 'S':
				data.s.x = i
				data.s.y = j
				row[j] = PATH
			}
		}
		data.m = append(data.m, row)
	}

	return &data
}
