package main

import "sort"

func part1(data *Data) int {
	sum := 0

	for _, update := range data.updates {
		updateData := UpdateData{data.rules, update}
		if sort.IsSorted(updateData) {
			sum += update[len(update)/2]
		}
	}

	return sum
}
