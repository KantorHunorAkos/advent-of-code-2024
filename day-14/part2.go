package main

import (
	"fmt"
)

func part2(data []Robot, mapSize *MapSize) int {
	size := len(data)

	for s:=1; ; s++{
		x, y, avgX,avgY := 0,0,0,0
		for i, r := range data {
			r.p.x = (r.p.x + (r.v.x + mapSize.x)) % mapSize.x
			r.p.y = (r.p.y + (r.v.y + mapSize.y)) % mapSize.y
			data[i] = r

			avgX += r.p.x
			avgY += r.p.y
		}

		avgX /= size
		avgY /= size

		for _, r := range data {
			x += (r.p.x-avgX) * (r.p.x-avgX)
			y += (r.p.y-avgY) * (r.p.y-avgY)
		}

		x /= size
		y /= size

		if x < 650 && y < 650 {
			printPic(makePic(mapSize, data))
			return s
		}
	}
}

func makePic(m *MapSize, robots []Robot) [][]int {
	pic := make([][]int, m.x)
	for i := 0; i < m.x; i++ {
		pic[i] = make([]int, m.y)
	}

	for _, r := range robots {
		pic[r.p.x][r.p.y] = 1
	}

	return pic
}

func printPic(pic [][]int) {
	for _, row := range pic {
		for _, px := range row {
			if px == 1 {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
