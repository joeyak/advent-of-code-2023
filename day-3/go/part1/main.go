package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	answer, err := solve()
	if err != nil {
		fmt.Printf("ERROR: could not solve: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Answer: %d\n", answer)
}

func solve() (int, error) {
	flag.Parse()
	input, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		return -1, fmt.Errorf("could not get input: %w", err)
	}

	result := 0
	data := strings.ReplaceAll(string(input), "\r\n", "\n")

	var symbols []Pos
	var numbers []Number
	for y, line := range strings.Split(data, "\n") {
		number := Number{}
		for x, r := range line {
			if unicode.IsDigit(r) {
				if number.Value == "" {
					number = Number{Pos: Pos{X: x, Y: y}}
				}
				number.Value += string(r)
				continue
			}

			if number.Value != "" {
				numbers = append(numbers, number)
			}

			if r != '.' {
				symbols = append(symbols, Pos{X: x, Y: y})
			}
			number = Number{}
		}
		if number.Value != "" {
			numbers = append(numbers, number)
		}
	}

	for _, n := range numbers {
		fmt.Printf("Number: %+v - ", n)
		isPartNum := false
		for _, s := range symbols {
			for mod := range n.Value {
				diffX := s.X - (n.X + mod)
				diffY := s.Y - n.Y
				if diffX <= 1 && diffX >= -1 && diffY <= 1 && diffY >= -1 {
					isPartNum = true
					break
				}
			}
			if isPartNum {
				break
			}
		}
		fmt.Println(isPartNum)
		if isPartNum {
			num, _ := strconv.Atoi(n.Value)
			result += num

		}
	}

	return result, nil
}

type Pos struct {
	X, Y int
}

type Number struct {
	Pos
	Value string
}
