package main

import "sync"

func part1(data []Operation) int {
	sum := 0
	var wg sync.WaitGroup
	var mut sync.Mutex

	for _, op := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isSolvable(op) {
				mut.Lock()
				sum += op.result
				mut.Unlock()
			}
		}()
	}
	wg.Wait()

	return sum
}

func isSolvable(op Operation) bool {
	return solve(op.operands[1:], op.result, op.operands[0])
}

func solve(opers []int, exp, acc int) bool {
	if len(opers) == 0 {
		return exp == acc
	}
	a := acc * opers[0]
	b := acc + opers[0]

	return solve(opers[1:], exp, a) || solve(opers[1:], exp, b)
}
