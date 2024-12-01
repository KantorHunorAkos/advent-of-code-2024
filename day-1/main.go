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
	file, err := os.Open("data.input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file '%s'\n", err)
		return
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
			return
		}
		ids1 = append(ids1, id1)
		id2, err := strconv.Atoi(nums[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't convert string %s to number", nums[3])
			return
		}
		ids2 = append(ids2, id2)
	}

	fmt.Println("Running part1")
	start := time.Now()
	part1(ids1, ids2)
	duration := time.Until(start).Abs()
	fmt.Printf("part1 ran succesfully\ntime: %s\n", duration.String())

	fmt.Println("Running part2")
	start = time.Now()
	part2(ids1, ids2)
	duration = time.Until(start).Abs()
	fmt.Printf("part2 ran succesfully\ntime: %s\n", duration.String())
}
