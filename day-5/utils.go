package main

import "slices"

func comparePages(a, b PageData) int {
	if slices.Contains(a.rulesBefore, b.page) {
		return -1
	}

	if slices.Contains(a.rulesAfter, b.page) {
		return 1
	}

	return 0
}
