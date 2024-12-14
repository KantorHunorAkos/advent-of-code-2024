package main

func cheapestPrizeCost(c ClawMachine) int {
	// Cramer
	det := (c.a.x*c.b.y - c.a.y*c.b.x)
	A := (c.prize.x*c.b.y - c.prize.y*c.b.x) / det
	B := (c.a.x*c.prize.y - c.a.y*c.prize.x) / det

	if c.a.x*A+c.b.x*B == c.prize.x && c.a.y*A+c.b.y*B == c.prize.y {

		return 3*A + B
	}
	return 0
}