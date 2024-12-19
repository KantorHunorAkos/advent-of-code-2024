package main

import (
	"math"
)

func part2(data []Point, size, time int) Point {
	m := getMatrix(data[:time], size)
	path := getEscapePath(Point{1,1}, Point{size,size}, m, size*size-len(data))
	for t := time + 1; t < len(data); t++ {
		next := Point{data[t].x+1,data[t].y+1}
		if contains(path, next) {
			m = getMatrix(data[:t+1], size)
			path = getEscapePath(Point{1,1}, Point{size,size}, m, size*size-len(data))
			if path == nil {
				return data[t]
			}
		}
	}
	return Point{}
}

func getEscapePath(s,e Point, m [][]int, size int) []Point {
	dist, prev := D(s, m, size)
	if dist[e] == math.MaxInt-1 {
		return nil
	}
	p := e
	list := []Point{}
	for !p.equal(s) {
		list = append(list, p)
		p = prev[p]
	}
	return list
}

func (p Point) equal(q Point) bool {
	return p.x == q.x && p.y == q.y
}

func contains(list []Point, p Point) bool {
	for _, e := range list {
		if p.equal(e) {
			return true
		}
	}
	return false
}