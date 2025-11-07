package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/19/input.txt")
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

type Design string

type Towel string

func ParseTowels(line string) map[Towel]struct{} {
	parts := strings.Split(line, ", ")
	towels := make(map[Towel]struct{}, len(parts))

	for _, part := range parts {
		towels[Towel(part)] = struct{}{}
	}

	return towels
}

func ParseDesigns(lines []string) map[Design]struct{} {
	designs := make(map[Design]struct{}, len(lines))

	for _, line := range lines {
		designs[Design(line)] = struct{}{}
	}

	return designs
}

func Part1(lines []string) (int, error) {
	total := 0

	towels := ParseTowels(lines[0])
	designs := ParseDesigns(lines[2:])
	cache := make(map[Design]bool)

	var canDisplay func(d Design) bool
	canDisplay = func(d Design) bool {
		if result, ok := cache[d]; ok {
			return result
		}

		// One towel can produce the entire design
		if _, ok := towels[Towel(d)]; ok {
			cache[d] = true
			return true
		}

		// Can a towel produce _part_ of the design
		for i := len(d) - 1; i > 0; i-- {
			if _, ok := towels[Towel(d[:i])]; ok {
				remainder := Design(d[i:])
				if canDisplay(remainder) {
					cache[d] = true
					return true
				}
			}
		}

		cache[d] = false
		return false
	}

	for design := range designs {
		if canDisplay(design) {
			total += 1
		}
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	total := 0

	towels := ParseTowels(lines[0])
	designs := ParseDesigns(lines[2:])

	cache := make(map[Design]int)

	var permutations func(d Design) int
	permutations = func(d Design) int {
		if result, ok := cache[d]; ok {
			return result
		}

		var total = 0
		if _, ok := towels[Towel(d)]; ok {
			total += 1
		}

		for i := len(d) - 1; i > 0; i-- {
			if _, ok := towels[Towel(d[:i])]; ok {
				remainder := Design(d[i:])
				total += permutations(remainder)
			}
		}

		cache[d] = total
		return total
	}

	for design := range designs {
		total += permutations(design)
	}

	return total, nil
}
