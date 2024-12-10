package main

import "slices"

func part2(data *AntennaMap) int {
	uniqueAntinodes := []Antinode{}

	for _, antennas := range data.antennas {
		for index, antenna := range antennas {
			for _, ant := range antennas[index+1:] {
				dx, dy := ant.x-antenna.x, ant.y-antenna.y
				antinode := Antinode(antenna)
				for antinode.inRange(data) {
					if !slices.Contains(uniqueAntinodes, antinode) {
						uniqueAntinodes = append(uniqueAntinodes, antinode)
					}
					antinode = Antinode{antinode.x - dx, antinode.y - dy}
				}
				antinode = Antinode(antenna)
				for antinode.inRange(data) {
					if !slices.Contains(uniqueAntinodes, antinode) {
						uniqueAntinodes = append(uniqueAntinodes, antinode)
					}
					antinode = Antinode{antinode.x + dx, antinode.y + dy}
				}
			}
		}
	}

	return len(uniqueAntinodes)
}
