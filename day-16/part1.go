package main

import (
	"fmt"
	"math"
)

func part1(data *Data) int {
	cornered := data.m.Copy()
	corn := cornered.getCorners()
	prev := D(data.s, corn, &data.m)

	sum := 0
	for p := data.e; !p.Equal(data.s); p = prev[p] {
		sum += 1000 + p.dist(prev[p])
	}
	return sum
}

func (m Maze) Copy() Maze {
	c := make(Maze, len(m))
	for i := range m {
		c[i] = make([]int, len(m[i]))
		copy(c[i], m[i])
	}
	return c
}

func (m *Maze) walkBFS(s Point) {
	directions := m.Copy()
	directions.markDirs(s, 1)
	m.set(s, 1)
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

func (m *Maze) markDirs(p Point, dir int) {
	if m.value(p) != PATH {
		return
	}

	m.set(p, dir)

	switch dir {
	case 1:
		m.markDirs(Point{p.x - 1, p.y}, 1)
		m.markDirs(Point{p.x, p.y - 1}, 3)
		m.markDirs(Point{p.x, p.y + 1}, 4)
	case 2:
		m.markDirs(Point{p.x + 1, p.y}, 2)
		m.markDirs(Point{p.x, p.y - 1}, 3)
		m.markDirs(Point{p.x, p.y + 1}, 4)
	case 3:
		m.markDirs(Point{p.x, p.y - 1}, 3)
		m.markDirs(Point{p.x - 1, p.y}, 1)
		m.markDirs(Point{p.x + 1, p.y}, 2)
	case 4:
		m.markDirs(Point{p.x, p.y + 1}, 4)
		m.markDirs(Point{p.x - 1, p.y}, 1)
		m.markDirs(Point{p.x + 1, p.y}, 2)
	}
}

func (m *Maze) value(p Point) int {
	return (*m)[p.x][p.y]
}

func (m *Maze) set(p Point, val int) {
	(*m)[p.x][p.y] = val
}

func (m Maze) print() {
	for _, row := range m {
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

func (p Point) GetNexts(m *Maze) []Point {
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

func (m Maze) getCorners() *[]Point {
	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m)-1; j++ {
			if (m)[i-1][j] == WALL && (m)[i+1][j] == WALL && (m)[i][j] != WALL {
				(m)[i][j] = 1
			}
			if (m)[i][j-1] == WALL && (m)[i][j+1] == WALL && (m)[i][j] != WALL {
				(m)[i][j] = 1
			}
		}
	}
	corners := []Point{}
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m); j++ {
			if (m)[i][j] == 0 {
				corners = append(corners, Point{i, j})
			}
		}
	}
	return &corners
}

func D(s Point, corners *[]Point, m *Maze) map[Point]Point {
	dist := make(map[Point]int, len(*corners))
	prev := make(map[Point]Point, len(*corners))
	pq := make(map[Point]bool, len(*corners))
	for _, c := range *corners {
		dist[c] = 1 << 8
		prev[c] = Point{}
		pq[c] = true
	}
	dist[s] = 0
	pq[s] = true

	for len(pq) > 0 {
		u := findMinDist(&pq, &dist)
		delete(pq, u)

		n := u.getNeighbors(&pq, m)

		for _, v := range n {
			alt := dist[u] + 1
			if alt < dist[*v] {
				dist[*v] = alt
				prev[*v] = u
			}
		}
	}

	return prev
}

func (p Point) dist(q Point) int {
	return int(math.Abs(float64(p.x-q.x) + float64(p.y-q.y)))
}

func findMinDist(pq *map[Point]bool, dist *map[Point]int) Point {
	minimum := math.MaxInt
	u := Point{}
	for v := range *pq {
		d := (*dist)[v]
		if minimum >= d {
			minimum = d
			u = v
		}
	}
	return u
}

func (p *Point) getNeighbors(pq *map[Point]bool, m *Maze) []*Point {
	n := []*Point{}

	for i := p.x + 1; i < len(*m)-1; i++ {
		if (*m)[i][p.y] == WALL {
			break
		}
		if _, exists := (*pq)[Point{i, p.y}]; exists {
			n = append(n, &Point{i, p.y})
		}
	}
	for i := p.x - 1; i > 0; i-- {
		if (*m)[i][p.y] == WALL {
			break
		}
		if _, exists := (*pq)[Point{i, p.y}]; exists {
			n = append(n, &Point{i, p.y})
		}
	}
	for i := p.y + 1; i < len((*m)[p.x])-1; i++ {
		if (*m)[p.x][i] == WALL {
			break
		}
		if _, exists := (*pq)[Point{p.x, i}]; exists {
			n = append(n, &Point{p.x, i})
		}
	}
	for i := p.y - 1; i > 0; i-- {
		if (*m)[p.x][i] == WALL {
			break
		}
		if _, exists := (*pq)[Point{p.x, i}]; exists {
			n = append(n, &Point{p.x, i})
		}
	}

	return n
}
