package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Antinode struct {
	x, y int
}

type Antenna struct {
	x, y int
}

type AntennaMap struct {
	antennas     map[rune][]Antenna
	sizeX, sizeY int
}

func (a Antinode) inRange(data *AntennaMap) bool {
	return a.x >= 0 && a.y >= 0 && a.x <= data.sizeX && a.y <= data.sizeY
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

func getInputFromFile(filename string) *AntennaMap {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := AntennaMap{
		antennas: make(map[rune][]Antenna),
		sizeX:    0,
		sizeY:    0,
	}

	for i := 0; scanner.Scan(); i++ {
		for j, c := range scanner.Text() {
			if c != '.' {
				data.antennas[c] = append(data.antennas[c], Antenna{i, j})
			}
			data.sizeY = j
		}
		data.sizeX = i
	}

	return &data
}
