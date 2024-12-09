package main

import (
	"bytes"
	"fmt"
)

func isStartingPoint(r rune) bool {
	if r == Up || r == Right || r == Down || r == Left {
		return true
	}

	return false
}

func getStartingPosition(mat [][]byte) (Point, error) {
	startingPoint := Point{-1, -1}
loop:
	for i, row := range mat {
		j := bytes.IndexFunc(row, isStartingPoint)
		if j != -1 {
			startingPoint = Point{i, j}
			break loop
		}
	}

	if startingPoint.x < 0 || startingPoint.y < 0 {
		return startingPoint, fmt.Errorf("couldn't find the starting point")
	}

	return startingPoint, nil
}

func (p Point) inBound(mat [][]byte) bool {
	return p.x >= 0 && p.x < len(mat) && p.y >= 0 && p.y < len(mat[0])
}

func copyData(mat [][]byte) [][]byte {
	dataCopy := make([][]byte, len(mat))
	for i, row := range mat {
		dataCopy[i] = make([]byte, len(row))
		copy(dataCopy[i], row)
	}
	return dataCopy
}

func moveToNextPositionGuard(direction byte, position Point, data [][]byte, mark byte) (byte, Point, [][]byte, byte) {
	// #? --> move back
	if data[position.x][position.y] == Obstacle {
		switch direction {
		case Up:
			position.x++
			direction = Right
		case Right:
			position.y--
			direction = Down
		case Left:
			position.y++
			direction = Up
		case Down:
			position.x--
			direction = Left
		}
	}
	//mark
	data[position.x][position.y] = mark
	//move
	switch direction {
	case Up:
		position.x--
	case Right:
		position.y++
	case Left:
		position.y--
	case Down:
		position.x++
	}

	return direction, position, data, mark
}
