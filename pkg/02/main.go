package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

type Report struct {
	Levels []int
}

type Direction int

const (
	Unknown Direction = iota
	Ascending
	Descending
)

func (r Report) IsSafe() bool {
	levels := len(r.Levels)

	var direction Direction
	var prev int = r.Levels[0]
	for i := 1; i < levels; i++ {
		level := r.Levels[i]

		diff := lib.Abs(level - prev)

		if diff > 3 || diff == 0 {
			return false
		}

		if direction == Unknown {
			if level > prev {
				direction = Ascending
			} else {
				direction = Descending
			}
		} else if direction == Ascending {
			if level < prev {
				return false
			}
		} else {
			if level > prev {
				return false
			}
		}
		prev = level
	}

	return true
}

func (r Report) IsSafeWithProblemDampener() bool {
	// Could use BFS here, but it's more effort
	if r.IsSafe() {
		return true
	}

	levels := len(r.Levels)

	for i := range r.Levels {
		dampedLevels := make([]int, levels-1)
		var offset int
		for j := range levels {
			if j == i {
				offset = 1
			} else {
				dampedLevels[j-offset] = r.Levels[j]
			}
		}

		if (Report{Levels: dampedLevels}).IsSafe() {
			return true
		}
	}

	return false
}

func main() {
	lines, err := lib.ReadLines("pkg/02/input.txt")
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

func ParseReport(line string) (*Report, error) {
	parts := strings.Split(line, " ")
	levels := make([]int, len(parts))

	for i, part := range parts {
		level, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("level %s: %w", part, err)
		}
		levels[i] = level
	}

	return &Report{
		Levels: levels,
	}, nil
}

func Part1(lines []string) (int, error) {
	safeReports := 0

	for i, line := range lines {
		report, err := ParseReport(line)
		if err != nil {
			return 0, fmt.Errorf("report %d: %w", i+1, err)
		}

		if report.IsSafe() {
			safeReports += 1
		}
	}

	return safeReports, nil
}

func Part2(lines []string) (int, error) {
	safeReportsWithProblemDampener := 0

	for i, line := range lines {
		report, err := ParseReport(line)
		if err != nil {
			return 0, fmt.Errorf("report %d: %w", i+1, err)
		}

		if report.IsSafeWithProblemDampener() {
			safeReportsWithProblemDampener += 1
		}
	}

	return safeReportsWithProblemDampener, nil
}
