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
	ids1, ids2 := getIds("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	solution := part1(ids1, ids2)
	duration := time.Until(start).Abs()
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %d\n", duration.String(), solution)

	fmt.Println("Running part2")
	start = time.Now()
	solution = part2(ids1, ids2)
	duration = time.Until(start).Abs()
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %d\n", duration.String(), solution)
}

func getIds(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file '%s'\n", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ids1 := []int{}
	ids2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		id1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't convert string %s to number", nums[0])
			return nil, nil
		}
		ids1 = append(ids1, id1)
		id2, err := strconv.Atoi(nums[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't convert string %s to number", nums[3])
			return nil, nil
		}
		ids2 = append(ids2, id2)
	}

	return ids1, ids2
}
