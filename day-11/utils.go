package main

import (
	"fmt"
	"math"
	"strconv"
)

func _(numberOfBlinks int, data []string) int {
	stones := make([]string, len(data))
	copy(stones, data)
	for i := 0; i < numberOfBlinks; i++ {
		for index := 0; index < len(stones); index++ {
			if len(stones[index])%2 == 0 {
				stoneOne, err := strconv.Atoi(stones[index][:len(stones[index])/2])
				if err != nil {
					fmt.Printf("Couldn't parse %q as a digit with error: %q", stones[index], err)
					return 0
				}
				stoneTwo, err := strconv.Atoi(stones[index][len(stones[index])/2:])
				if err != nil {
					fmt.Printf("Couldn't parse %q as a digit with error: %q", stones[index], err)
					return 0
				}
				newStones := []string{strconv.Itoa(stoneOne), strconv.Itoa(stoneTwo)}
				if index+1 < len(stones) {
					newStones = append(newStones, stones[index+1:]...)
				}
				stones = append(stones[:index], newStones...)
				index++
			} else {
				number, err := strconv.Atoi(stones[index])
				if err != nil {
					fmt.Printf("Couldn't parse %q as a digit with error: %q", stones[index], err)
					return 0
				}
				if number == 0 {
					stones[index] = "1"
				} else {
					stones[index] = strconv.Itoa(number * 2024)
				}
			}
		}
	}
	return len(stones)
}

func countStones(numberOfBlinks int, data []int) int {
	stones := make(map[int]int, len(data))
	for _, stone := range data {
		stones[stone] = 1
	}
	for i := 0; i < numberOfBlinks; i++ {
		newStones := make(map[int]int, len(stones))
		for stone, nr := range stones {
			stoneOne, stoneTwo := blink(stone)
			newStones[stoneOne] = newStones[stoneOne] + nr
			if stoneTwo != -1 {
				newStones[stoneTwo] = newStones[stoneTwo] + nr
			}
		}
		stones = newStones
	}
	sum := 0
	for _, stoneCount := range stones {
		sum += stoneCount
	}
	return sum
}

func blink(stone int) (int, int) {
	size := size(stone)

	if size == 0 {
		return 1, -1
	}

	if size%2 == 1 {
		return stone * 2024, -1
	}

	exp := int(math.Pow10(size / 2))

	return stone / exp, stone % exp
}

func size(i int) int {
	s := 0
	for i != 0 {
		s++
		i /= 10
	}
	return s
}
