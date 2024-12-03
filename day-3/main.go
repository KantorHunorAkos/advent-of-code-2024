package main

import (
	"bufio"
	"fmt"
	"os"
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

func getInputFromFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file '%s'\n", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := ""
	for scanner.Scan() {
		data += scanner.Text()
	}

	return data
}
