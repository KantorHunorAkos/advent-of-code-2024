package main

import (
	"fmt"
	"os"
	"regexp"
)

func part2(data string) int {
	doRegex, err := regexp.Compile(`do\(\)`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create regular expression `do\\(\\)` with error: %s", err)
		return 0
	}

	dontRegex, err := regexp.Compile(`don't\(\)`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create regular expression `don't\\(\\)` with error: %s", err)
		return 0
	}

	donts := dontRegex.FindStringIndex(data)
	for donts != nil {
		dos := doRegex.FindStringIndex(data[donts[1]:])
		if dos != nil {
			data = data[:donts[0]] + data[donts[1]:][dos[1]:]
		} else {
			data = data[:donts[0]]
		}
		donts = dontRegex.FindStringIndex(data)
	}

	return part1(data)
}
