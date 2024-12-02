package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	data := getInputFromFile()

	fmt.Println("Running part1")
	start := time.Now()
	solution := part1(data)
	duration := time.Until(start).Abs()
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", duration.String(), solution)

	fmt.Println("Running part2")
	start = time.Now()
	solution = part2(data)
	duration = time.Until(start).Abs()
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %d\n", duration.String(), solution)
}

func getInputFromFile() any {
	file, err := os.Open("data.input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file '%s'\n", err)
		return true
	}
	defer file.Close()


	
	return nil
}
