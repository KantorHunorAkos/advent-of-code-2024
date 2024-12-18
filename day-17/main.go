package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	adv = iota // A div 2^combo -> A
	bxl        // B XOR literal -> B
	bst        // combo % 8 -> b
	jnz        // A==0 ? nothing : pc = literal
	bxc        // B XOR  C -> B
	out        // combo % 8 -> stdout
	bdv        // A div 2^combo -> B
	cdv        // A div 2^combo -> C
)

type Comp struct {
	A, B, C      int
	pc           int
	prog, output []int
}

func main() {
	data := getInputFromFile("data.input")

	fmt.Println("Running part1")
	start := time.Now()
	solution := part1(data)
	fmt.Printf("Part one ran succesfully\n Time: %s\n Solution: %v\n", time.Since(start), solution)

	fmt.Println("Running part2")
	start = time.Now()
	solution2 := part2(data)
	fmt.Printf("Part two ran succesfully\n Time: %s\n Solution: %v\n", time.Since(start), solution2)
}

func getInputFromFile(filename string) *Comp {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while opening the input file(%s) '%s'\n", filename, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	comp := Comp{
		pc:     0,
		output: []int{},
	}

	scanner.Scan()
	A, err := strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
	if err != nil {
		fmt.Printf("Coundn't parse number with error: %q\n", err)
		return nil
	}
	comp.A = A

	scanner.Scan()
	B, err := strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
	if err != nil {
		fmt.Printf("Coundn't parse number with error: %q\n", err)
		return nil
	}
	comp.B = B

	scanner.Scan()
	C, err := strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
	if err != nil {
		fmt.Printf("Coundn't parse number with error: %q\n", err)
		return nil
	}
	comp.C = C

	scanner.Scan()
	scanner.Scan()
	nums := strings.Split(strings.Split(scanner.Text(), ": ")[1], ",")
	comp.prog = make([]int, len(nums))
	for i, n := range nums {
		p, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("Coundn't parse number with error: %q\n", err)
			return nil
		}
		comp.prog[i] = p
	}

	return &comp
}
