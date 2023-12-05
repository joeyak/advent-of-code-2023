package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("ERROR: could not get input: %v\n", err)
		os.Exit(1)
	}

	data := string(input)

	part1Answer, err := solvePart1(data)
	if err != nil {
		fmt.Printf("ERROR: could not solve part 1: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 Answer: %d\n", part1Answer)

	part2Answer, err := solvePart2(data)
	if err != nil {
		fmt.Printf("ERROR: could not solve part 2: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 Answer: %d\n", part2Answer)
}

func solvePart1(data string) (int, error) {
	result := 0

	for i, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}

		var first, last string
		for _, r := range line {
			c := string(r)
			_, err := strconv.Atoi(c)
			if err != nil {
				continue
			}

			if first == "" {
				first = c
				continue
			}

			last = c
		}

		if last == "" {
			last = first
		}

		number, err := strconv.Atoi(first + last)
		if err != nil {
			return -1, fmt.Errorf("could not convert line %d %#v to number: %w", i, line, err)
		}

		result += number
	}

	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0

	replacements := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for i, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}

		var first, last, past string
		setNumbers := func(c string) {
			if first == "" {
				first = c
				return
			}

			last = c
		}

		for _, r := range line {
			c := string(r)
			_, err := strconv.Atoi(c)
			if err != nil {
				past += c

				for old, new := range replacements {
					if strings.HasSuffix(past, old) {
						setNumbers(new)
						break
					}
				}

				continue
			}

			past = ""
			setNumbers(c)

		}

		if last == "" {
			last = first
		}

		number, err := strconv.Atoi(first + last)
		if err != nil {
			return -1, fmt.Errorf("could not convert line %d %#v to number: %w", i, line, err)
		}

		result += number
	}

	return result, nil
}
