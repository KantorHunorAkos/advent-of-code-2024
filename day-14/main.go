package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Position struct{ x, y int }
type Velocity struct{ x, y int }
type Robot struct {
	p Position
	v Velocity
}

type MapSize struct{ x, y int }

func main() {
	data := getInputFromFile("data.input")

	fmt.Println("Running part1")
	m := &MapSize{101, 103}
	start := time.Now()
	solution := part1(data, m)
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)

	fmt.Println("Running part2")
	start = time.Now()
	solution = part2(data, m)
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)
}

func getInputFromFile(filename string) []Robot {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := []Robot{}

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		p, v := tokens[0][2:], tokens[1][2:]
		tokens = strings.Split(p, ",")
		x, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Couldn't parse %q as a digit with error: %q", x, err)
			return nil
		}
		y, err := strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Printf("Couldn't parse %q as a digit with error: %q", x, err)
			return nil
		}
		pos := Position{x, y}
		tokens = strings.Split(v, ",")
		x, err = strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Couldn't parse %q as a digit with error: %q", x, err)
			return nil
		}
		y, err = strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Printf("Couldn't parse %q as a digit with error: %q", x, err)
			return nil
		}
		velo := Velocity{x, y}

		robots = append(robots, Robot{pos, velo})
	}

	return robots
}
