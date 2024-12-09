package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	content, err := lib.ReadFile("pkg/05/input.txt")
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

func ParseRules(raw string) (map[int]map[int]struct{}, error) {
	lines := strings.Split(raw, "\n")
	rules := make(map[int]map[int]struct{}, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid rule %s on line %d", line, i+1)
		}

		before, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("left page number on line %d: %w", i+1, err)
		}

		after, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("right page number on line %d: %w", i+1, err)
		}

		if _, ok := rules[before]; !ok {
			rules[before] = make(map[int]struct{})
		}

		rules[before][after] = struct{}{}
	}

	return rules, nil
}

func ParseUpdates(raw string) ([][]int, error) {
	lines := strings.Split(raw, "\n")
	updates := make([][]int, len(lines))

	for i, line := range lines {
		values := strings.Split(line, ",")
		updates[i] = make([]int, len(values))

		for j, v := range values {
			page, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("invalid page %s on line %d: %w", v, i+1, err)
			}
			updates[i][j] = page
		}
	}

	return updates, nil
}

func Order(rules map[int]map[int]struct{}, _update []int) []int {
	update := make([]int, len(_update))
	copy(update, _update)

	slices.SortFunc(update, func(a, b int) int {
		if afterA, ok := rules[a]; ok {
			_, ok := afterA[b]
			if ok {
				return -1
			}
		}
		if afterB, ok := rules[b]; ok {
			if _, ok := afterB[a]; ok {
				return 1
			}
		}

		return 0
	})

	return update
}

func InOrder(rules map[int]map[int]struct{}, update []int) bool {
	seen := make(map[int]struct{})
	for _, page := range update {
		after, ok := rules[page]
		if ok && lib.Intersect(after, seen) {
			return false
		}
		seen[page] = struct{}{}
	}
	return true
}

func Part1(content string) (int, error) {
	total := 0

	parts := strings.Split(content, "\n\n")

	rules, err := ParseRules(parts[0])
	if err != nil {
		return 0, fmt.Errorf("rule parsing: %w", err)
	}

	updates, err := ParseUpdates(parts[1])
	if err != nil {
		return 0, fmt.Errorf("update parsing: %w", err)
	}

	for _, update := range updates {
		if InOrder(rules, update) {
			total += update[len(update)/2]
		}
	}

	return total, nil
}

func Part2(content string) (int, error) {
	total := 0

	parts := strings.Split(content, "\n\n")

	rules, err := ParseRules(parts[0])
	if err != nil {
		return 0, fmt.Errorf("rule parsing: %w", err)
	}

	updates, err := ParseUpdates(parts[1])
	if err != nil {
		return 0, fmt.Errorf("update parsing: %w", err)
	}

	for _, update := range updates {
		if InOrder(rules, update) {
			continue
		}

		ordered := Order(rules, update)

		total += ordered[len(update)/2]

	}

	return total, nil
}
