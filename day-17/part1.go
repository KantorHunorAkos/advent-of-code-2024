package main

import "fmt"

func part1(comp *Comp) []int {
	for comp.pc < int(len(comp.prog)) {
		switch comp.prog[comp.pc] {
		case adv:
			comp.adv(comp.prog[comp.pc+1])
		case bxl:
			comp.bxl(comp.prog[comp.pc+1])
		case bst:
			comp.bst(comp.prog[comp.pc+1])
		case jnz:
			comp.jnz(comp.prog[comp.pc+1])
		case bxc:
			comp.bxc(comp.prog[comp.pc+1])
		case out:
			comp.out(comp.prog[comp.pc+1])
		case bdv:
			comp.bdv(comp.prog[comp.pc+1])
		case cdv:
			comp.cdv(comp.prog[comp.pc+1])
		default:
			panic(fmt.Sprintf("Unknown opcode: %v", comp.prog[comp.pc]))
		}
	}

	return comp.output
}

func (c *Comp) adv(input int) {
	c.A = c.A >> c.combo(input)
	c.pc += 2
}

func (c *Comp) bxl(input int) {
	c.B = c.B ^ input
	c.pc += 2
}

func (c *Comp) bst(input int) {
	c.B = c.combo(input) & 7
	c.pc += 2
}

func (c *Comp) jnz(input int) {
	if c.A == 0 {
		c.pc += 2
		return
	}
	c.pc = input
}

func (c *Comp) bxc(input int) {
	_ = input
	c.B = c.B ^ c.C
	c.pc += 2
}

func (c *Comp) out(input int) {
	c.output = append(c.output, c.combo(input)&7)
	c.pc += 2
}

func (c *Comp) bdv(input int) {
	c.B = c.A >> c.combo(input)
	c.pc += 2
}

func (c *Comp) cdv(input int) {
	c.C = c.A >> c.combo(input)
	c.pc += 2
}

func (c *Comp) combo(in int) int {
	switch in {
	case 0, 1, 2, 3:
		return in
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	case 7:
		panic(fmt.Errorf("combo operand 7 is reserved and will not appear in valid programs"))
	default:
		panic(fmt.Errorf("unknown combo operand: %v", in))
	}
}
