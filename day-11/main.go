package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func getInputFromFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	data := []int{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, token := range strings.Split(scanner.Text(), " ") {
		number, err := strconv.Atoi(token)
		if err != nil {
			fmt.Printf("Couldn't parse %q as a digit with error: %q", token, err)
			return nil
		}
		data = append(data, number)
	}
	return data
}
