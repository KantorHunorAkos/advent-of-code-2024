package main

import (
	"fmt"
	"slices"
)

func part2(comp *Comp) int {
	regA := 1
	output := []int{}
	for ; !slices.Equal(output, comp.prog); regA = regA << 3 {
		A := regA
		output = []int{}
		for A != 0 {
			B := (A & 0b111) ^ 0b11
			C := A >> B
			output = append(output, (B ^ C) & 0b111)
			A = A >> 3
		}

		fmt.Println(output)
	}

	return regA
}
