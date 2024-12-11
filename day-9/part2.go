package main

func part2(data []int) int {
	freeIndex := 0
	for i := len(data) - 1; i > 0; i-- {
		if data[i] == -1 {
			continue
		}

		freeIndex = findFreeIndex(0, data)
		freeSize := findSize(-1, freeIndex, data)
		usedIndex := findIndex(data[i], data)
		usedSize := findSize(data[usedIndex], usedIndex, data)
		for freeSize < usedSize && freeIndex < usedIndex {
			j := freeIndex + 1
			for data[j] == -1 {
				j++
			}
			freeIndex = findFreeIndex(j, data)
			freeSize = findSize(-1, freeIndex, data)
		}
		i = usedIndex

		if freeIndex > i {
			continue
		}

		for j := 0; j < usedSize; j++ {
			data[freeIndex+j] = data[usedIndex+j]
			data[usedIndex+j] = -1
		}
	}

	checksum := 0
	for index, id := range data {
		if id > 0 {
			checksum += index * id
		}
	}

	return checksum
}

func findFreeIndex(start int, data []int) int {
	index := 0

	for j := start; j < len(data); j++ {
		if data[j] == -1 {
			index = j
			break
		}
	}

	return index
}

func findIndex(id int, data []int) int {
	index := 0

	for j := 0; j < len(data); j++ {
		if data[j] == id {
			index = j
			break
		}
	}

	return index
}

func findSize(id, start int, data []int) int {
	size := 0

	for j := start; j < len(data) && data[j] == id; j++ {
		size++
	}

	return size
}
