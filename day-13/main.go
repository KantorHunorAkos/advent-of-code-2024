package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Button struct {
	x, y int
}

type Point struct {
	x, y int
}

type ClawMachine struct {
	a, b  Button
	prize Point
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

func getInputFromFile(filename string) *[]ClawMachine {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := []ClawMachine{}
	ax, ay := 0, 0
	bx, by := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		tokens := strings.Split(line, ":")
		ident, coords := strings.TrimSpace(tokens[0]), strings.Split(tokens[1], ", ")
		x, y := strings.TrimSpace(coords[0])[2:], strings.TrimSpace(coords[1])[2:]
		xNum, err := strconv.Atoi(x)
		if err != nil {
			fmt.Printf("Couldn't parse %q as a digit with error: %q", x, err)
			return nil
		}
		yNum, err := strconv.Atoi(y)
		if err != nil {
			fmt.Printf("Couldn't parse %q as a digit with error: %q", y, err)
			return nil
		}
		if ident == "Button A" {
			ax, ay = xNum, yNum
		}
		if ident == "Button B" {
			bx, by = xNum, yNum
		}
		if ident == "Prize" {
			data = append(data, ClawMachine{a: Button{ax, ay}, b: Button{bx, by}, prize: Point{xNum, yNum}})
		}
	}

	return &data
}
