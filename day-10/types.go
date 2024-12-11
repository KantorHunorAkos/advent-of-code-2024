package main

type Coord struct{ x, y int }

type GuideData struct {
	grid      [][]int
	tailheads []Coord
}
