package main

import (
	"math"
	"sort"
)

func part1(ids1 []int, ids2 []int) int {
	sort.Ints(ids1)
	sort.Ints(ids2)

	lenght := len(ids1)
	sum := 0

	for i := 0; i < lenght; i++ {
		sum += int(math.Abs(float64(ids1[i] - ids2[i])))
	}

	return sum
}
