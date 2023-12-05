package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	debug := false

	part1Answer, err := solvePart1(debug)
	if err != nil {
		fmt.Printf("ERROR: could not solve part 1: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 Answer: %d\n", part1Answer)

	part2Answer, err := solvePart2(debug)
	if err != nil {
		fmt.Printf("ERROR: could not solve part 2: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 Answer: %d\n", part2Answer)
}

func solvePart1(debug bool) (int, error) {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		return -1, fmt.Errorf("could not get input: %w", err)
	}

	data := strings.ReplaceAll(string(input), "\r\n", "\n")

	maxes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	result := 0
	for gameNum, line := range strings.Split(data, "\n") {
		gameNum += 1
		if debug {
			fmt.Print(gameNum)
		}
		success := true
		for _, set := range strings.Split(strings.Split(line, ":")[1], ";") {
			for _, cube := range strings.Split(set, ",") {
				parts := strings.Split(strings.TrimSpace(cube), " ")
				num, _ := strconv.Atoi(parts[0])
				if num > maxes[parts[1]] {
					success = false
					break
				}
			}
			if !success {
				break
			}
		}

		if success {
			result += gameNum
		}
		if debug {
			fmt.Println(" ", success)
		}
	}

	return result, nil
}

func solvePart2(debug bool) (int, error) {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		return -1, fmt.Errorf("could not get input: %w", err)
	}

	data := strings.ReplaceAll(string(input), "\r\n", "\n")

	result := 0
	for gameNum, line := range strings.Split(data, "\n") {
		gameNum += 1
		if debug {
			fmt.Print(gameNum)
		}
		maxes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, set := range strings.Split(strings.Split(line, ":")[1], ";") {
			for _, cube := range strings.Split(set, ",") {
				parts := strings.Split(strings.TrimSpace(cube), " ")
				num, _ := strconv.Atoi(parts[0])
				if num > maxes[parts[1]] {
					maxes[parts[1]] = num
				}
			}
		}

		gameResult := 1
		for _, num := range maxes {
			gameResult *= num
		}

		if debug {
			fmt.Println(" ", maxes, " - ", gameResult)
		}
		result += gameResult
	}

	return result, nil
}
