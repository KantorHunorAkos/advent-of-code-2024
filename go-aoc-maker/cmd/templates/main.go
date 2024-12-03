package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	data := getInputFromFile("data.input")

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

func getInputFromFile(filename string) any {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return true
	}
	defer file.Close()


	
	return nil
}
