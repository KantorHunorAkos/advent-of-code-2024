package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	UP    = '^'
	RIGHT = '>'
	DOWN  = 'v'
	LEFT  = '<'
)

const (
	BOX            = 'O'
	BOX_LEFT_SIDE  = '['
	BOX_RIGHT_SIDE = ']'
	WALL           = '#'
	ROBOT          = '@'
	EMPTY          = '.'
)

type Robot struct{ x, y int }

type Warehouse [][]rune
type Direction rune
type Directions []rune

type PuzzleData struct {
	w Warehouse
	d Directions
	r Robot
}

func main() {
	data := getInputFromFile("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	solution := part1(data)
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)

	data = getInputFromFile("data.input")
	fmt.Println("Running part2")
	start = time.Now()
	solution = part2(data)
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)
}

func getInputFromFile(filename string) *PuzzleData {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	data := PuzzleData{}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := []rune(scanner.Text())
		if len(line) == 0 {
			break
		}
		for j, c := range line {
			if c == ROBOT {
				data.r = Robot{i, j}
				line[j] = EMPTY
			}
		}
		data.w = append(data.w, line)
	}
	for scanner.Scan() {
		line := []rune(scanner.Text())
		data.d = append(data.d, line...)
	}

	return &data
}
