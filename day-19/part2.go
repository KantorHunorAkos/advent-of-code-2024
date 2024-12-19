package main

import "strings"

func part2(towels, arrangements []string) int {
	count := 0

	for _, arr := range arrangements {
		filteredTowels := []string{}
		for _, t := range towels {
			if strings.Contains(arr, t) {
				filteredTowels = append(filteredTowels, t)
			}
		}
		count += countPossible(filteredTowels, arr)
	}

	return count
}
