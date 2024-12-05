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

func getInputFromFile(filename string) *Data {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := OrderingRules{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		rule := strings.Split(line, "|")
		before, err := strconv.Atoi(rule[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't convert first string (%s) into a number with error: %s", rule[0], err)
			return nil
		}
		after, err := strconv.Atoi(rule[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't convert second string (%s) into a number with error: %s", rule[1], err)
			return nil
		}

		rules = append(rules, OrderingRule{
			x: before,
			y: after,
		})
	}

	updates := Updates{}
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		numbers := Update{}
		for _, token := range tokens {
			number, err := strconv.Atoi(token)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't convert second string (%s) into a number with error: %s", token, err)
				return nil
			}
			numbers = append(numbers, number)
		}
		updates = append(updates, numbers)
	}

	return &Data{rules, updates}
}
