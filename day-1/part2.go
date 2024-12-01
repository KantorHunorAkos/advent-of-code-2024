package main

import (
	"fmt"
)

func part2(ids1 []int, ids2 []int) {
	map1 := make(map[int]int)

	for _, value := range ids1 {
		map1[value] = map1[value] + 1
	}

	map2 := make(map[int]int)

	for _, value := range ids2 {
		map2[value] = map2[value] + 1
	}

	sum := 0

	for k, v := range map1 {
		sum += v * k * map2[k]
	}

	fmt.Println(sum)
}
