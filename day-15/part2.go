package main

func part2(data *PuzzleData) int {
	w := data.w.getWideWarehouse()
	data.r.y = data.r.y * 2
	for _, m := range data.d {
		data.r.tryMoveWide(Direction(m), w)
	}
	sum := 0
	for _, v := range w.getWideGPSCoordinates() {
		sum += v
	}
	return sum
}

func (w *Warehouse) getWideGPSCoordinates() []int {
	coords := []int{}
	for i, row := range *w {
		for j, c := range row {
			if c == BOX_LEFT_SIDE {
				coords = append(coords, 100*i+j)
			}
		}
	}
	return coords
}

func (w *Warehouse) getWideWarehouse() Warehouse {
	ww := make(Warehouse, len(*w))
	for i := range ww {
		ww[i] = make([]rune, len((*w)[i])*2)
	}
	for i, row := range *w {
		wj := 0
		for _, c := range row {
			if c == BOX {
				ww[i][wj] = BOX_LEFT_SIDE
				wj++
				ww[i][wj] = BOX_RIGHT_SIDE
				wj++
			} else {
				ww[i][wj] = c
				wj++
				ww[i][wj] = c
				wj++
			}
		}
	}
	return ww
}

func (r *Robot) tryMoveWide(d Direction, w Warehouse) {
	switch d {
	case UP:
		if r.canMoveTo(-1, 0, w) || r.tryShiftBoxWide(-1, 0, w) {
			r.x--
		}
	case RIGHT:
		if r.canMoveTo(0, +1, w) || r.tryShiftBoxWide(0, +1, w) {
			r.y++
		}
	case DOWN:
		if r.canMoveTo(+1, 0, w) || r.tryShiftBoxWide(+1, 0, w) {
			r.x++
		}
	case LEFT:
		if r.canMoveTo(0, -1, w) || r.tryShiftBoxWide(0, -1, w) {
			r.y--
		}
	}
}

func (r *Robot) tryShiftBoxWide(x, y int, w Warehouse) bool {
	if x != 0 {
		if r.canShiftWideBoxesVertically(x, w) {
			return r.shiftWideBoxesVertically(x, w)
		}
	}
	if y != 0 {
		for i := r.y + y; w[r.x][i] != WALL; i += y {
			if w[r.x][i] == EMPTY {
				for j := i; w[r.x][j-y] != EMPTY; j -= y {
					w[r.x][j] = w[r.x][j-y]
				}
				w[r.x][r.y+y] = EMPTY
				return true
			}
		}
	}
	return false
}

func (r *Robot) canShiftWideBoxesVertically(x int, w Warehouse) bool {
	if w[r.x+x][r.y] == WALL {
		return false
	}
	r0 := Robot{r.x + x, r.y}
	r1 := Robot{r.x + x, r.y}
	if w[r0.x][r.y] == BOX_LEFT_SIDE {
		r1.y++
	} else if w[r0.x][r.y] == BOX_RIGHT_SIDE {
		r1.y--
	}
	if w[r0.x][r1.y] == EMPTY && w[r1.x][r1.y] == EMPTY {
		return true
	}
	return r0.canShiftWideBoxesVertically(x, w) && r1.canShiftWideBoxesVertically(x, w)
}

func (r *Robot) shiftWideBoxesVertically(x int, w Warehouse) bool {
	if w[r.x+x][r.y] == WALL {
		return false
	}
	r0 := Robot{r.x + x, r.y}
	r1 := Robot{r.x + x, r.y}
	if w[r0.x][r.y] == BOX_LEFT_SIDE {
		r1.y++
	} else if w[r0.x][r.y] == BOX_RIGHT_SIDE {
		r1.y--
	}
	if w[r0.x][r1.y] == EMPTY && w[r1.x][r1.y] == EMPTY {
		return true
	}
	
	canMove := r0.shiftWideBoxesVertically(x, w) && r1.shiftWideBoxesVertically(x, w)
	if canMove {
		w[r0.x+x][r0.y] = w[r0.x][r0.y]
		w[r1.x+x][r1.y] = w[r1.x][r1.y]
		w[r0.x][r0.y] = EMPTY
		w[r1.x][r1.y] = EMPTY
	}
	return canMove
}
