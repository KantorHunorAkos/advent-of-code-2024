package main

import "slices"

type Plot struct {
	area, perimeter int
}

type Point struct {
	x, y int
}

func part1(data []string) int {
	garden := getGarden(data)
	regions := []Plot{}

	for row := 1; row < len(garden)-1; row++ {
		for col := 1; col < len(garden[row])-1; col++ {
			if garden[row][col] != '*' {
				visited := []Point{}
				regions = append(regions, walk(garden[row][col], row, col, garden, &visited))
			}
		}
	}

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.area * region.perimeter
	}

	return totalPrice
}

func getGarden(data []string) [][]byte {
	garden := make([][]byte, len(data)+2)
	for i := 0; i < len(data)+2; i++ {
		garden[i] = make([]byte, len(data[0])+2)
	}
	for i := 0; i < len(data[0])+2; i++ {
		garden[0][i] = '*'
		garden[len(garden)-1][i] = '*'
	}
	for i := 1; i < len(garden)-1; i++ {
		garden[i][0] = '*'
		for j := 1; j < len(garden[0])-1; j++ {
			garden[i][j] = data[i-1][j-1]
		}
		garden[i][len(garden[i])-1] = '*'
	}
	return garden
}

func walk(label byte, x, y int, garden [][]byte, visited *[]Point) Plot {
	if slices.Contains(*visited, Point{x, y}) {
		return Plot{0, 0}
	}
	if garden[x][y] != label {
		return Plot{0, 1}
	}

	*visited = append(*visited, Point{x, y})
	garden[x][y] = '*'

	up := walk(label, x-1, y, garden, visited)
	down := walk(label, x+1, y, garden, visited)
	left := walk(label, x, y-1, garden, visited)
	right := walk(label, x, y+1, garden, visited)

	return Plot{
		area:      up.area + down.area + left.area + right.area + 1,
		perimeter: up.perimeter + down.perimeter + left.perimeter + right.perimeter,
	}
}
