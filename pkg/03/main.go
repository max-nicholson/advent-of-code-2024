package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/03/input.txt")
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

func ParseInstruction(match []string) (int, int, error) {
	a, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

func Part1(lines []string) (int, error) {
	total := 0

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	for i, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, b, err := ParseInstruction(match)
			if err != nil {
				return 0, fmt.Errorf("line %d, %s: %w", i, match[0], err)
			}
			total += a * b
		}
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	total := 0

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)

	var enabled = true
	for i, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if enabled {
				a, b, err := ParseInstruction(match)
				if err != nil {
					return 0, fmt.Errorf("line %d, %s: %w", i, match[0], err)
				}
				total += a * b
			}
		}
	}

	return total, nil
}
