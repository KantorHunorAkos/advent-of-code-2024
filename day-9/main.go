package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	data := getInputFromFile("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	solution := part1(data)
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", time.Since(start), solution)

	data = getInputFromFile("data.input")

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

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	data := []int{}
	id := 0
	free := false
	for _, r := range line {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Printf("couldn't parse %q as digit with error: %q", r, err)
			return nil
		}

		for i := 0; i < digit; i++ {
			if free {
				data = append(data, -1)
			} else {
				data = append(data, id)
			}
		}

		if !free {
			id++
		}

		free = !free
	}

	return data
}
