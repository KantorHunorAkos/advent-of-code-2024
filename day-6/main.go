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

func getInputFromFile(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := [][]byte{}

	for scanner.Scan() {
		text := scanner.Text()
		bytes := []byte(text)
		data = append(data, bytes)
	}

	return data
}
