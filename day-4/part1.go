package main

func part1(data []string) int {
	sum := 0

	for i, row := range data {
		for j, c := range row {
			if c == 'X' {
				if j+3 < len(data[i]) {
					if data[i][j+1] == 'M' && data[i][j+2] == 'A' && data[i][j+3] == 'S' {
						sum++
					}
				}
				if j-3 >= 0 {
					if data[i][j-1] == 'M' && data[i][j-2] == 'A' && data[i][j-3] == 'S' {
						sum++
					}
				}
				if i+3 < len(data) {
					if data[i+1][j] == 'M' && data[i+2][j] == 'A' && data[i+3][j] == 'S' {
						sum++
					}
				}
				if i-3 >= 0 {
					if data[i-1][j] == 'M' && data[i-2][j] == 'A' && data[i-3][j] == 'S' {
						sum++
					}
				}
				if i-3 >= 0 && j-3 >= 0 {
					if data[i-1][j-1] == 'M' && data[i-2][j-2] == 'A' && data[i-3][j-3] == 'S' {
						sum++
					}
				}
				if i-3 >= 0 && j+3 < len(data) {
					if data[i-1][j+1] == 'M' && data[i-2][j+2] == 'A' && data[i-3][j+3] == 'S' {
						sum++
					}
				}
				if i+3 < len(data) && j-3 >= 0 {
					if data[i+1][j-1] == 'M' && data[i+2][j-2] == 'A' && data[i+3][j-3] == 'S' {
						sum++
					}
				}
				if i+3 < len(data) && j+3 < len(data[i]) {
					if data[i+1][j+1] == 'M' && data[i+2][j+2] == 'A' && data[i+3][j+3] == 'S' {
						sum++
					}
				}
			}
		}
	}

	return sum
}
