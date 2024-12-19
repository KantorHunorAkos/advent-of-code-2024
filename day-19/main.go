package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	towels, arrangements := getInputFromFile("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	solution := part1(towels, arrangements)
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)

	fmt.Println("Running part2")
	start = time.Now()
	solution = part2(towels, arrangements)
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)
}

func getInputFromFile(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	towels := strings.Split(scanner.Text(), ", ")

	scanner.Scan()
	arrangements := []string{}
	for scanner.Scan() {
		arrangements = append(arrangements, scanner.Text())
	}

	return towels, arrangements
}
