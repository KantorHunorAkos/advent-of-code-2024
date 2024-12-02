package main

import (
	"fmt"
)

func part1(data [][]int) {
	sum := 0

	for _, row := range data {
		if isSafe(row) {
			sum++
		}
	}

	fmt.Println(sum)
}

func isSafe(row []int) bool {
	if row[0] < row[1] {
		return isSafeAscending(row)
	}

	return isSafeDecending(row)
}

func isSafeAscending(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		diff := row[i+1] - row[i]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafeDecending(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		diff := row[i] - row[i+1]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}