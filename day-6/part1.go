package main

import (
	"bytes"
	"fmt"
)

const (
	Up       = '^'
	Right    = '>'
	Left     = '<'
	Down     = 'v'
	Obstacle = '#'
	Mark     = 'X'
)

type Point struct {
	x int
	y int
}

func part1(data [][]byte) int {
	startingPoint := Point{-1, -1}
loop:
	for i, row := range data {
		j := bytes.IndexFunc(row, isStartingPoint)
		if j != -1 {
			startingPoint = Point{i, j}
			break loop
		}
	}

	if startingPoint.x < 0 || startingPoint.y < 0 {
		fmt.Println("Couldn't find the starting point")
		return 0
	}

	direction := data[startingPoint.x][startingPoint.y]
	position := startingPoint

	for position.inBound(data) {
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
		data[position.x][position.y] = Mark
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
	}

	sum := 0
	for _, row := range data {
		for _, r := range row {
			if r == Mark {
				sum++
			}
		}
	}
	return sum
}

func isStartingPoint(r rune) bool {
	if r == Up || r == Right || r == Down || r == Left {
		return true
	}

	return false
}

func (p Point) inBound(mat [][]byte) bool {
	return p.x >= 0 && p.x < len(mat) && p.y >= 0 && p.y < len(mat[0])
}
