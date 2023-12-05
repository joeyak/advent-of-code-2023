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

	var gears []*Gear
	var parts []*Part
	for y, line := range strings.Split(data, "\n") {
		part := &Part{}
		for x, r := range line {
			if unicode.IsDigit(r) {
				if part.Value == "" {
					part = &Part{Pos: Pos{X: x, Y: y}}
				}
				part.Value += string(r)
				continue
			}

			if part.Value != "" {
				parts = append(parts, part)
			}

			if r != '.' {
				gears = append(gears, &Gear{Pos: Pos{X: x, Y: y}, Value: r})
			}
			part = &Part{}
		}
		if part.Value != "" {
			parts = append(parts, part)
		}
	}

	for _, part := range parts {
		for _, gear := range gears {
			for mod := range part.Value {
				diffX := gear.X - (part.X + mod)
				diffY := gear.Y - part.Y
				if diffX <= 1 && diffX >= -1 && diffY <= 1 && diffY >= -1 {
					part.Gears = append(part.Gears, gear)
					gear.Parts = append(gear.Parts, part)
					break
				}
			}
		}
		fmt.Printf("Part: %+v\n", part)
	}

	for _, gear := range gears {
		fmt.Printf("Gear: %v %v", gear.Pos, gear.Value)
		if gear.Value == '*' {
			if len(gear.Parts) == 2 {
				fmt.Print(" -")
				gearValue := 1
				for _, part := range gear.Parts {
					fmt.Printf(" %v", part.Value)
					num, _ := strconv.Atoi(part.Value)
					gearValue *= num
				}
				result += gearValue
				fmt.Printf(" - %d", gearValue)
			}
		}
		fmt.Println()
	}

	return result, nil
}

type Pos struct {
	X, Y int
}

type Part struct {
	Pos
	Value string
	Gears []*Gear
}

type Gear struct {
	Pos
	Value rune
	Parts []*Part
}
