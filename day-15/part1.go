package main

import "fmt"

func part1(data *PuzzleData) int {
	for _, m := range data.d {
		data.r.tryMove(Direction(m), data.w)
	}
	sum := 0
	for _, v := range data.w.getGPSCoordinates() {
		sum += v
	}
	return sum
}

func (r *Robot) tryMove(d Direction, w Warehouse) {
	switch d {
	case UP:
		if r.canMoveTo(-1, 0, w) || r.tryShiftBox(-1, 0, w) {
			r.x--
		}
	case RIGHT:
		if r.canMoveTo(0, +1, w) || r.tryShiftBox(0, +1, w) {
			r.y++
		}
	case DOWN:
		if r.canMoveTo(+1, 0, w) || r.tryShiftBox(+1, 0, w) {
			r.x++
		}
	case LEFT:
		if r.canMoveTo(0, -1, w) || r.tryShiftBox(0, -1, w) {
			r.y--
		}
	}
}

func (r *Robot) canMoveTo(x, y int, w Warehouse) bool {
	return w[r.x+x][r.y+y] == EMPTY
}

func (r *Robot) tryShiftBox(x, y int, w Warehouse) bool {
	if x != 0 {
		for i := r.x + x; w[i][r.y] != WALL; i += x {
			if w[i][r.y] == EMPTY {
				w[i][r.y] = BOX
				w[r.x+x][r.y] = EMPTY
				return true
			}
		}
	}
	if y != 0 {
		for i := r.y + y; w[r.x][i] != WALL; i += y {
			if w[r.x][i] == EMPTY {
				w[r.x][i] = BOX
				w[r.x][r.y+y] = EMPTY
				return true
			}
		}
	}
	return false
}

func (r *Robot) Print(w Warehouse) {
	for i, row := range w {
		for j, c := range row {
			if i == r.x && j == r.y {
				fmt.Print("@")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
}

func (r *Robot) PrintWide(w Warehouse) {
	for i, row := range w {
		for j, c := range row {
			if i == r.x && j == r.y {
				fmt.Print("@.")
			} else if c == BOX {
				fmt.Print("[]")
			} else {
				fmt.Printf("%c%c", c, c)
			}
		}
		fmt.Println()
	}
}

func (w *Warehouse) getGPSCoordinates() []int {
	coords := []int{}
	for i, row := range *w {
		for j, c := range row {
			if c == BOX {
				coords = append(coords, 100*i+j)
			}
		}
	}
	return coords
}
