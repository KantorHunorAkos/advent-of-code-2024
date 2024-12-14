package main

func part1(data []Robot, mapSize *MapSize) int {
	midx := mapSize.x / 2
	midy := mapSize.y / 2
	s := 100
	quadrants := [4]int{}

	for _, r := range data {
		x := (r.p.x + (r.v.x+mapSize.x)*s) % mapSize.x
		y := (r.p.y + (r.v.y+mapSize.y)*s) % mapSize.y

		if x < midx && y < midy {
			quadrants[0]++
		}
		if x > midx && y < midy {
			quadrants[1]++
		}
		if x < midx && y > midy {
			quadrants[2]++
		}
		if x > midx && y > midy {
			quadrants[3]++
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}
