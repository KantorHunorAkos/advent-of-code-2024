package main

import (
	"slices"
)

func part1(data *AntennaMap) int {
	uniqueAntinodes := []Antinode{}

	for _, antennas := range data.antennas {
		for index, antenna := range antennas {
			for _, ant := range antennas[index+1:] {
				antinode := Antinode{ant.x*2 - antenna.x, ant.y*2 - antenna.y}
				if antinode.inRange(data) && !slices.Contains(uniqueAntinodes, antinode) {
					uniqueAntinodes = append(uniqueAntinodes, antinode)
				}
				antinode = Antinode{antenna.x*2 - ant.x, antenna.y*2 - ant.y}
				if antinode.inRange(data) && !slices.Contains(uniqueAntinodes, antinode) {
					uniqueAntinodes = append(uniqueAntinodes, antinode)
				}
			}
		}
	}

	return len(uniqueAntinodes)
}
