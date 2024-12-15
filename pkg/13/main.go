package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	content, err := lib.ReadFile("pkg/13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

type Button struct {
	x    int
	y    int
	cost int
}

type Prize struct {
	x int
	y int
}

type Machine struct {
	A, B  Button
	Prize Prize
}

func InvertMatrix(matrix [2][2]int) [2][2]float64 {
	a := matrix[0][0]
	b := matrix[0][1]
	c := matrix[1][0]
	d := matrix[1][1]

	determinant := float64(1) / float64(a*d-b*c)

	return [2][2]float64{
		{determinant * float64(d), determinant * float64(-b)},
		{determinant * float64(-c), determinant * float64(a)},
	}
}

const epsilon = 1e-4 // Margin of error

func IsWholeNumber(n float64) bool {
	if _, frac := math.Modf(math.Abs(n)); frac < epsilon || frac > 1.0-epsilon {
		return true
	} else {
		return false
	}
}

func (m Machine) MinTokensForPrize() int {
	inverse := InvertMatrix([2][2]int{
		{m.A.x, m.B.x},
		{m.A.y, m.B.y},
	})

	result := [2]float64{
		inverse[0][0]*float64(m.Prize.x) + inverse[0][1]*float64(m.Prize.y),
		inverse[1][0]*float64(m.Prize.x) + inverse[1][1]*float64(m.Prize.y),
	}

	if IsWholeNumber(result[0]) && IsWholeNumber(result[1]) {
		return int(math.Round(result[0]))*3 + int(math.Round(result[1]))
	}

	return 0
}

var BUTTON_REGEX = regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
var PRIZE_REGEX = regexp.MustCompile(`X=(\d+), Y=(\d+)`)

func ParseMachines(content string) ([]Machine, error) {
	content = strings.TrimRight(content, "\n")

	parts := strings.Split(content, "\n\n")
	machines := make([]Machine, len(parts))

	for i, part := range parts {
		lines := strings.Split(part, "\n")
		if len(lines) != 3 {
			return nil, fmt.Errorf("want 3 lines of machine configuration, got %d", len(lines))
		}

		machine := Machine{}

		{
			button := BUTTON_REGEX.FindStringSubmatch(lines[0])
			if len(button) != 3 {
				return nil, fmt.Errorf("machine %d; button %s; want 2 matches, got %s", i, lines[0], button)
			}
			x, err := strconv.Atoi(button[1])
			if err != nil {
				return nil, fmt.Errorf("machine %d; button %s; invalid Button A X position %s: %w", i, lines[0], button[1], err)
			}
			y, err := strconv.Atoi(button[2])
			if err != nil {
				return nil, fmt.Errorf("machine %d; button %s; invalid Button A Y position %s: %w", i, lines[0], button[2], err)
			}
			machine.A = Button{x, y, 3}
		}

		{
			button := BUTTON_REGEX.FindStringSubmatch(lines[1])
			if len(button) != 3 {
				return nil, fmt.Errorf("machine %d; want 2 matches, got %s", i, button)
			}
			x, err := strconv.Atoi(button[1])
			if err != nil {
				return nil, fmt.Errorf("machine %d; invalid Button B X position %s: %w", i, button[1], err)
			}
			y, err := strconv.Atoi(button[2])
			if err != nil {
				return nil, fmt.Errorf("machine %d; invalid Button B Y position %s: %w", i, button[2], err)
			}
			machine.B = Button{x, y, 1}
		}

		{
			prize := PRIZE_REGEX.FindStringSubmatch(lines[2])
			if len(prize) != 3 {
				return nil, fmt.Errorf("machine %d prize; want 2 matches, got %s", i, prize)
			}
			x, err := strconv.Atoi(prize[1])
			if err != nil {
				return nil, fmt.Errorf("machine %d prize; invalid X position %s: %w", i, prize[1], err)
			}
			y, err := strconv.Atoi(prize[2])
			if err != nil {
				return nil, fmt.Errorf("machine %d prize; invalid Y position %s: %w", i, prize[2], err)
			}
			machine.Prize = Prize{x, y}
		}

		machines[i] = machine
	}
	return machines, nil
}

func Part1(content string) (int, error) {
	total := 0
	machines, err := ParseMachines(content)
	if err != nil {
		return 0, fmt.Errorf("machine parsing: %w", err)
	}

	for _, machine := range machines {
		total += machine.MinTokensForPrize()
	}

	return total, nil
}

func Part2(content string) (int, error) {
	total := 0
	machines, err := ParseMachines(content)
	if err != nil {
		return 0, fmt.Errorf("machine parsing: %w", err)
	}

	for _, machine := range machines {
		machine.Prize.x += 10000000000000
		machine.Prize.y += 10000000000000
		total += machine.MinTokensForPrize()
	}

	return total, nil
}
