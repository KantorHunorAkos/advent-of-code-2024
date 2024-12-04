package main

func part2(data []string) int {
	sum := 0

	for i, row := range data {
		for j, c := range row {
			if c == 'A' {
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(data) && j+1 < len(row) {
					if (data[i-1][j-1] == 'M' && data[i+1][j+1] == 'S' && data[i-1][j+1] == 'M' && data[i+1][j-1] == 'S') ||
						(data[i-1][j-1] == 'M' && data[i+1][j+1] == 'S' && data[i+1][j-1] == 'M' && data[i-1][j+1] == 'S') ||
						(data[i+1][j+1] == 'M' && data[i-1][j-1] == 'S' && data[i-1][j+1] == 'M' && data[i+1][j-1] == 'S') ||
						(data[i+1][j+1] == 'M' && data[i-1][j-1] == 'S' && data[i+1][j-1] == 'M' && data[i-1][j+1] == 'S') {
						sum++
					}
				}
			}
		}
	}

	return sum
}
