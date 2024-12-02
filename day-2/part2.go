package main

import (
	"fmt"
	"slices"
)

func part2(data [][]int) {
	sum := 0

	for _, row := range data {
		if isSafeWithProblemDampener(row, 1) {
			sum++
		}
	}

	fmt.Println(sum)
}

func isSafeWithProblemDampener(row []int, toleranceLevel uint) bool {
	direction := 0
	for i := 0; i < len(row)-1; i++ {
		if row[i] - row[i+1] < 0 {
			direction++
		} else {
			direction--
		}
	}

	if direction >= 0 {
		return isSafeAscendingWithProblemDampener(row, toleranceLevel)
	}
	
	return isSafeDecendingWithProblemDampener(row, toleranceLevel)
}

func isSafeAscendingWithProblemDampener(row []int, toleranceLevel uint) bool {
	for i := 0; i < len(row)-1; i++ {
		diff := row[i+1] - row[i]
		if diff < 1 || diff > 3 {
			return checkWithOneRemoved(row, i, toleranceLevel)
		}
	}
	return true
}

func isSafeDecendingWithProblemDampener(row []int, toleranceLevel uint) bool {
	for i := 0; i < len(row)-1; i++ {
		diff := row[i] - row[i+1]
		if diff < 1 || diff > 3 {
			return checkWithOneRemoved(row, i, toleranceLevel)
		}
	}
	return true
}

func checkWithOneRemoved(row []int, index int, toleranceLevel uint) bool {

	if toleranceLevel == 0 {
		return false
	}

	removed := slices.Clone(row)
	removed = slices.Delete(removed, index, index+1)
	if isSafeWithProblemDampener(removed, toleranceLevel-1) {
		return true
	}

	removed = slices.Clone(row)
	removed = slices.Delete(removed, index+1, index+2)
	return isSafeWithProblemDampener(removed, toleranceLevel-1)
}