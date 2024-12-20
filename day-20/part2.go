package main

import (
	"math"
)

func part2(data *Data, minSave int) int {
	data.m.walkBFS(data.s)
	count := 0
	minSave += 2
	for i := 1; i < len(data.m)-1; i++ {
		for j := 1; j < len(data.m[i])-1; j++ {
			p := Point{i, j}
			if data.m.value(p) == WALL {
				continue
			}
			currSave := 0
			for k := 1; k < len(data.m)-1; k++ {
				for l := 1;l < len(data.m[i])-1; l++ {
					d := Point{k, l}
					if p.dist(d) > 20 {
						break
					}
					if data.m.value(d) == WALL {
						continue
					}
					if data.m.value(d) > data.m.value(p) {
						if data.m.value(d)-data.m.value(p) > currSave {
							currSave = data.m.value(d) - data.m.value(p)
						}
					}
				}
			}
			if currSave > minSave {
				count++
			}
		}

	}
	return count
}

func (p Point) dist(q Point) int {
	return int(math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2)))
}
