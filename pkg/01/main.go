package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/01/input.txt")
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

func ParseLine(line string) (int, int, error) {
	parts := strings.Split(line, "   ")
	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("left location: %w", err)
	}

	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("right location: %w", err)
	}
	return left, right, nil
}

func Part1(lines []string) (int, error) {
	total := 0
	locations := len(lines)

	leftList := make([]int, locations)
	rightList := make([]int, locations)

	for i, line := range lines {
		left, right, err := ParseLine(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", i, err)
		}

		leftList[i] = left
		rightList[i] = right
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i := range locations {
		total += lib.Abs(leftList[i] - rightList[i])
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	similarityScore := 0
	totalLocations := len(lines)

	locations := make([]int, totalLocations)
	locationCounts := make(map[int]int, totalLocations/3)

	for i, line := range lines {
		left, right, err := ParseLine(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", i, err)
		}
		locations = append(locations, left)
		locationCounts[right] += 1
	}

	for _, location := range locations {
		similarityScore += locationCounts[location] * location
	}

	return similarityScore, nil
}
