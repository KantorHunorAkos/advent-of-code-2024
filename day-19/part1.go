package main

import "strings"

func part1(towels, arrangements []string) int {
	possible := 0

	for _, arr := range arrangements {
		filteredTowels := []string{}
		for _, t := range towels {
			if strings.Contains(arr, t) {
				filteredTowels = append(filteredTowels, t)
			}
		}
		if countPossible(filteredTowels, arr) > 0 {
			possible++
		}
	}

	return possible
}

func countPossible(towels []string, a string) int {
	aLen := len(a)
	dp := make([]int, aLen+1)
	dp[0] = 1
	for i := 0; i < aLen; i++ {
		if dp[i] == 0 {
			continue
		}
		for _, towel := range towels {
			endIdx := len(towel) + i
			if endIdx <= aLen && towel == a[i:endIdx] {
				dp[endIdx] += dp[i]
			}
		}
	}
	return dp[aLen]
}
