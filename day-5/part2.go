package main

import (
	"slices"
)

func part2(updates *Updates) int {
	sum := 0

	for _, update := range *updates {
		if !slices.IsSortedFunc(update, comparePages) {
			slices.SortFunc(update, comparePages)
			sum += update[len(update)/2].page
		}
	}

	return sum
}
