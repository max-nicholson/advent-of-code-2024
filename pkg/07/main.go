package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

type Equation struct {
	numbers []int
	value   int
}

type Operator int

const (
	Add Operator = iota + 1
	Multiply
	Concatenate
)

type State struct {
	index    int
	value    int
	operator Operator
}

func (e Equation) SolveableWith(operators []Operator) bool {
	if len(e.numbers) == 1 {
		return e.numbers[0] == e.value
	}

	stack := []State{}
	for _, op := range operators {
		stack = append(stack, State{index: 1, value: e.numbers[0], operator: op})
	}

	for len(stack) > 0 {
		l := len(stack) - 1
		s := stack[l]
		stack = stack[:l]
		var value = s.value

		switch s.operator {
		case Add:
			value += e.numbers[s.index]
		case Multiply:
			value *= e.numbers[s.index]
		case Concatenate:
			v := strconv.Itoa(value)
			v += strconv.Itoa(e.numbers[s.index])
			value, _ = strconv.Atoi(v)
		}

		if s.index == len(e.numbers)-1 {
			if value == e.value {
				return true
			}

			// Try another combination
			continue
		}

		if value > e.value {
			// We've already gone past the target
			// Since our only operations are additive, it's now impossible to reach the target
			continue
		}

		for _, op := range operators {
			stack = append(stack, State{index: s.index + 1, value: value, operator: op})
		}
	}

	return false
}

func ParseEquations(lines []string) ([]Equation, error) {
	equations := make([]Equation, len(lines))
	for i, equation := range lines {
		parts := strings.Split(equation, ": ")
		v := parts[0]
		value, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("line %d equation %s value %s: %w", i+1, equation, v, err)
		}

		nums := strings.Split(parts[1], " ")
		numbers := make([]int, len(nums))

		for j, n := range nums {
			v, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("line %d equation %s number %s: %w", i+1, equation, n, err)
			}
			numbers[j] = v
		}

		equations[i].value = value
		equations[i].numbers = numbers
	}

	return equations, nil
}

func Part1(lines []string) (int, error) {
	total := 0

	equations, err := ParseEquations(lines)
	if err != nil {
		return 0, err
	}

	operators := []Operator{Add, Multiply}

	for _, equation := range equations {
		if equation.SolveableWith(operators) {
			total += equation.value
		}
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	total := 0

	equations, err := ParseEquations(lines)
	if err != nil {
		return 0, err
	}

	operators := []Operator{Add, Multiply, Concatenate}

	for _, equation := range equations {
		if equation.SolveableWith(operators) {
			total += equation.value
		}
	}

	return total, nil
}
