package main

import (
	"fmt"
	"sync"
)

func part2(data [][]byte) int {
	startingPoint, err := getStartingPosition(data)
	if err != nil {
		fmt.Printf("%s\n", err)
		return 0
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	sum := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == Obstacle || isStartingPoint(rune(data[i][j])) {
				continue
			}

			wg.Add(1)

			go func() {
				defer wg.Done()

				dataCopy := copyData(data)
				visited := make([][]int, len(data))
				for k := range data {
					visited[k] = make([]int, len(data[i]))
				}

				dataCopy[i][j] = Obstacle
				visited[i][j]++

				direction := data[startingPoint.x][startingPoint.y]
				dataCopy[startingPoint.x][startingPoint.y] = '.'
				position := startingPoint

				for position.inBound(dataCopy) {
					visited[position.x][position.y]++
					if visited[position.x][position.y] > 4 {
						mu.Lock()
						sum++
						mu.Unlock()
						break
					}
					direction, position, dataCopy, _ = moveToNextPositionGuard(direction, position, dataCopy, direction)
				}
			}()
		}
	}

	wg.Wait()

	return sum
}
