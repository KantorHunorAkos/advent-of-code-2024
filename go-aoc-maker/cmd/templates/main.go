package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Open("data.input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file '%s'\n", err)
		return
	}
	defer file.Close()

	fmt.Println("Running part1")
	start := time.Now()
	part1(nil)
	duration := time.Until(start).Abs()
	fmt.Printf("part1 ran succesfully\ntime: %s\n", duration.String())

	fmt.Println("Running part2")
	start = time.Now()
	part2(nil)
	duration = time.Until(start).Abs()
	fmt.Printf("part2 ran succesfully\ntime: %s\n", duration.String())
}
