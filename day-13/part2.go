package main

func part2(data *[]ClawMachine) int {
	sum := 0
	for _, clawMachine := range *data {
		clawMachine = ClawMachine{
			a: clawMachine.a,
			b: clawMachine.b,
			prize: Point{
				clawMachine.prize.x + 10000000000000,
				clawMachine.prize.y + 10000000000000,
			},
		}
		sum += cheapestPrizeCost(clawMachine)
	}
	return sum
}
