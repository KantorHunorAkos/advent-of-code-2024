package main

func solve(data *GuideData, isPart1 bool) int {
	score := 0
	for _, tailhead := range data.tailheads {
		grid := make([][]int, len(data.grid))
		for i := range data.grid {
			grid[i] = make([]int, len(data.grid[i]))
			copy(grid[i], data.grid[i])
		}
		score += walk(tailhead, grid, isPart1)
	}
	return score
}

func walk(start Coord, grid [][]int, isPart1 bool) int {
	if grid[start.x][start.y] == 9 {
		if isPart1 {
			grid[start.x][start.y] = 10
		}
		return 1
	}

	acc := 0
	if start.x-1 >= 0 {
		if grid[start.x-1][start.y]-grid[start.x][start.y] == 1 {
			acc += walk(Coord{start.x - 1, start.y}, grid, isPart1)
		}
	}

	if start.x+1 < len(grid) {
		if grid[start.x+1][start.y]-grid[start.x][start.y] == 1 {
			acc += walk(Coord{start.x + 1, start.y}, grid, isPart1)
		}
	}

	if start.y-1 >= 0 {
		if grid[start.x][start.y-1]-grid[start.x][start.y] == 1 {
			acc += walk(Coord{start.x, start.y - 1}, grid, isPart1)
		}
	}

	if start.y+1 < len(grid[start.x]) {
		if grid[start.x][start.y+1]-grid[start.x][start.y] == 1 {
			acc += walk(Coord{start.x, start.y + 1}, grid, isPart1)
		}
	}

	return acc
}
