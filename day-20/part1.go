package main

import (
	"fmt"
)

func part1(data *Data, minSave int) int {
	data.m.walkBFS(data.s)
	count := 0
	minSave += 2

	for i := 1; i < len(data.m)-1; i++ {
		for j := 1; j < len(data.m[i])-1; j++ {
			p := Point{i, j}
			if data.m.value(p) != WALL {
				continue
			}

			if data.m.isCheat(p, Point{-1, 0}, minSave) {
				count++
			}

			if data.m.isCheat(p, Point{0, +1}, minSave) {
				count++
			}

			if data.m.isCheat(p, Point{+1, 0}, minSave) {
				count++
			}

			if data.m.isCheat(p, Point{0, -1}, minSave) {
				count++
			}

		}
	}

	return count
}

func (m *Track) isCheat(p, d Point, minSave int) bool {
	if m.value(Point{p.x + d.x, p.y + d.y}) != WALL &&
		m.value(Point{p.x + d.x, p.y + d.y}) > m.value(Point{p.x - d.x, p.y - d.y}) {
		if m.value(Point{p.x + d.x, p.y + d.y})-m.value(Point{p.x - d.x, p.y - d.y}) >= minSave {
			return true
		}
	}
	return false
}

func (m *Track) walkBFS(s Point) {
	m.set(s, 0)
	queue := []Point{s}
	visited := []Point{s}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		adj := q.GetNexts(m)
		for _, p := range adj {
			if !Contains(visited, p) {
				visited = append(visited, p)
				queue = append(queue, p)
				m.set(p, m.value(q)+1)
			}
		}
	}
}

func (m *Track) value(p Point) int {
	return (*m)[p.x][p.y]
}

func (m *Track) set(p Point, val int) {
	(*m)[p.x][p.y] = val
}

func (m *Track) print() {
	for _, row := range *m {
		for _, v := range row {
			if v == WALL {
				fmt.Printf("####")
			} else {
				fmt.Printf("%3d ", v)
			}
		}
		fmt.Println()
	}
}

func (p0 Point) Equal(p1 Point) bool {
	return p0.x == p1.x && p0.y == p1.y
}

func Contains(v []Point, p Point) bool {
	for _, p0 := range v {
		if p.Equal(p0) {
			return true
		}
	}
	return false
}

func (p Point) GetNexts(m *Track) []Point {
	nexts := []Point{}
	nx := Point{p.x - 1, p.y}
	if m.value(nx) != WALL {
		nexts = append(nexts, nx)
	}
	nx = Point{p.x + 1, p.y}
	if m.value(nx) != WALL {
		nexts = append(nexts, nx)
	}
	nx = Point{p.x, p.y - 1}
	if m.value(nx) != WALL {
		nexts = append(nexts, nx)
	}
	nx = Point{p.x, p.y + 1}
	if m.value(nx) != WALL {
		nexts = append(nexts, nx)
	}
	return nexts
}
