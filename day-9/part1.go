package main

func part1(data []int) int {
	for i := len(data) - 1; i > 0; i-- {
		if data[i] == -1 {
			continue
		}

		firstFree := 0
		for j := 0; j < len(data); j++ {
			if data[j] == -1 {
				firstFree = j
				break
			}
		}

		if firstFree > i {
			data = data[:firstFree]
			break
		}

		data[firstFree] = data[i]
		data[i] = -1
	}

	checksum := 0
	for index, id := range data {
		checksum += index * id
	}

	return checksum
}
