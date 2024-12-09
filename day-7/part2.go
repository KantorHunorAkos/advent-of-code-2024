package main

import (
	"fmt"
	"strconv"
	"sync"
)

func part2(data []Operation) int {
	sum := 0
	var wg sync.WaitGroup
	var mut sync.Mutex

	for _, op := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isSolvablePart2(op) {
				mut.Lock()
				sum += op.result
				mut.Unlock()
			}
		}()
	}

	wg.Wait()
	return sum
}

func isSolvablePart2(op Operation) bool {
	return solvePart2(op.operands[1:], op.result, op.operands[0])
}

func solvePart2(opers []int, exp, acc int) bool {
	if len(opers) == 0 {
		return exp == acc
	}
	a := acc * opers[0]
	b := acc + opers[0]
	concatenatedString := strconv.Itoa(acc) + strconv.Itoa(opers[0])
	c, err := strconv.Atoi(concatenatedString)
	if err != nil {
		panic(fmt.Errorf("coundn't parse %q into result with error: %q", concatenatedString, err))
	}

	return solvePart2(opers[1:], exp, a) || solvePart2(opers[1:], exp, b) || solvePart2(opers[1:], exp, c)
}
