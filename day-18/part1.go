package main

import (
	"fmt"
	"math"
)

const (
	FREE      = 0
	CORRUPTED = 1
)

func part1(data []Point, size, time int) int {
	m := getMatrix(data[:time], size)
	dist, _ := D(Point{1, 1}, m, size*size-len(data))
	return dist[Point{size, size}]
}

func getMatrix(points []Point, size int) [][]int {
	m := make([][]int, size+2)
	for i := range m {
		m[i] = make([]int, size+2)
	}
	for i := range m {
		m[0][i] = CORRUPTED
		m[size+1][i] = CORRUPTED
		m[i][0] = CORRUPTED
		m[i][size+1] = CORRUPTED
	}
	for _, p := range points {
		m[p.x+1][p.y+1] = CORRUPTED
	}
	return m
}

func print(m [][]int) {
	for i := range m {
		for j := range m {
			if m[i][j] == FREE {
				fmt.Print(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}

func D(s Point, m [][]int, size int) (map[Point]int, map[Point]Point) {
	dist := make(map[Point]int, size)
	prev := make(map[Point]Point, size)
	pq := make(map[Point]bool, size)
	for i := range m {
		for j := range m {
			if m[i][j] == FREE {
				p0 := Point{i, j}
				dist[p0] = math.MaxInt - 1
				prev[p0] = Point{}
				pq[p0] = true
			}
		}
	}
	dist[s] = 0

	for len(pq) > 0 {
		u := findMinDist(&pq, &dist)
		delete(pq, u)

		n := u.getNeighbors(&pq, m)

		for _, v := range n {
			alt := dist[u] + 1
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	return dist, prev
}

func findMinDist(pq *map[Point]bool, dist *map[Point]int) Point {
	minimum := math.MaxInt
	u := Point{}
	for v := range *pq {
		d := (*dist)[v]
		if minimum > d {
			minimum = d
			u = v
		}
	}
	return u
}

func (p *Point) getNeighbors(pq *map[Point]bool, m [][]int) []Point {
	n := []Point{}

	if m[p.x+1][p.y] == FREE {
		if _, exists := (*pq)[Point{p.x + 1, p.y}]; exists {
			n = append(n, Point{p.x + 1, p.y})
		}

	}
	if m[p.x-1][p.y] == FREE {
		if _, exists := (*pq)[Point{p.x - 1, p.y}]; exists {
			n = append(n, Point{p.x - 1, p.y})
		}

	}
	if m[p.x][p.y+1] == FREE {
		if _, exists := (*pq)[Point{p.x, p.y + 1}]; exists {
			n = append(n, Point{p.x, p.y + 1})
		}
	}
	if m[p.x][p.y-1] == FREE {
		if _, exists := (*pq)[Point{p.x, p.y - 1}]; exists {
			n = append(n, Point{p.x, p.y - 1})
		}
	}

	return n
}
