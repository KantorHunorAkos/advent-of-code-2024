package main

import "slices"

type Plot2 struct {
	area, corner int
}

func part2(data []string) int {
	garden := getGarden(data)
	regions := []Plot2{}

	for row := 1; row < len(garden)-1; row++ {
		for col := 1; col < len(garden[row])-1; col++ {
			if garden[row][col] != '*' {
				visited := []Point{}
				regions = append(regions, walk2(garden[row][col], row, col, garden, &visited))
				for _, v := range visited {
					garden[v.x][v.y] = '*'
				}
			}
		}
	}

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.area * region.corner
	}

	return totalPrice
}

func walk2(label byte, x, y int, garden [][]byte, visited *[]Point) Plot2 {
	if !inGarden(x, y, garden) {
		return Plot2{0, 0}
	}
	if slices.Contains(*visited, Point{x, y}) || garden[x][y] != label {
		return Plot2{0, 0}
	}

	*visited = append(*visited, Point{x, y})

	up := walk2(label, x-1, y, garden, visited)
	down := walk2(label, x+1, y, garden, visited)
	left := walk2(label, x, y-1, garden, visited)
	right := walk2(label, x, y+1, garden, visited)

	corner := countCorners(x, y, garden, label)
	return Plot2{
		area:   up.area + down.area + left.area + right.area + 1,
		corner: up.corner + down.corner + left.corner + right.corner + corner,
	}
}

func countCorners(x, y int, garden [][]byte, lable byte) int {
	if !inGarden(x, y, garden) {
		return 0
	}
	corner := 0
	if garden[x-1][y] != lable && garden[x][y+1] != lable {
		corner++
	}
	if garden[x-1][y] != lable && garden[x][y-1] != lable {
		corner++
	}
	if garden[x+1][y] != lable && garden[x][y+1] != lable {
		corner++
	}
	if garden[x+1][y] != lable && garden[x][y-1] != lable {
		corner++
	}
	if garden[x-1][y] == lable && garden[x][y+1] == lable && garden[x-1][y+1] != lable {
		corner++
	}
	if garden[x+1][y] == lable && garden[x][y-1] == lable && garden[x+1][y-1] != lable {
		corner++
	}
	if garden[x+1][y] == lable && garden[x][y+1] == lable && garden[x+1][y+1] != lable {
		corner++
	}
	if garden[x-1][y] == lable && garden[x][y-1] == lable && garden[x-1][y-1] != lable {
		corner++
	}
	return corner
}
func inGarden(x, y int, garden [][]byte) bool {
	if x-1 < 0 || y-1 < 0 || x+1 >= len(garden) || y+1 >= len(garden[0]) {
		return false
	}
	return true
}
