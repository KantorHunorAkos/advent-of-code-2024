package main

import (
	"fmt"
)

func part1(data [][]byte) int {
	dataCopy := copyData(data)
	startingPoint, err := getStartingPosition(dataCopy)
	if err != nil {
		fmt.Printf("%s\n", err)
		return 0
	}

	direction := dataCopy[startingPoint.x][startingPoint.y]
	position := startingPoint

	for position.inBound(dataCopy) {
		direction, position, dataCopy, _ = moveToNextPositionGuard(direction, position, dataCopy, Mark)
	}

	sum := 0
	for _, row := range dataCopy {
		for _, r := range row {
			if r == Mark {
				sum++
			}
		}
	}
	return sum
}
