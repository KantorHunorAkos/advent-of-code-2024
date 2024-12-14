package main

func part1(data *[]ClawMachine) int {
	sum := 0
	for _, clawMachine := range *data {
		sum += cheapestPrizeCost(clawMachine)
	}
	return sum
}
