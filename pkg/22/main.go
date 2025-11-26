package main

import (
	"fmt"
	"log"
	"maps"
	"slices"
	"strconv"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/22/input.txt")
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

type Secret int

func NewSecret(s string) (Secret, error) {
	secret, err := strconv.Atoi(s)
	return Secret(secret), err
}

func (secret Secret) Mix(n int) Secret {
	return Secret(n ^ int(secret))
}

func (secret Secret) Prune() Secret {
	return Secret(int(secret) % 16777216)
}

func (secret Secret) Next() Secret {
	secret = secret.Mix(int(secret) * 64).Prune()
	secret = secret.Mix(int(secret) / 32).Prune()
	secret = secret.Mix(int(secret) * 2048).Prune()

	return secret
}

func (secret Secret) Price() int {
	return int(secret) % 10
}

func Part1(lines []string) (int, error) {
	total := 0

	for i, line := range lines {

		secret, err := NewSecret(line)
		if err != nil {
			return 0, fmt.Errorf("unable to parse line %d (%s) as number: %w", i+1, line, err)
		}

		for range 2000 {
			secret = secret.Next()
		}

		total += int(secret)
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	bananasByMonkeyBySequence := make([]map[[4]int]int, len(lines))

	for lineNumber, line := range lines {
		bananasBySequence := make(map[[4]int]int, 0)

		secret, err := NewSecret(line)
		if err != nil {
			return 0, fmt.Errorf("unable to parse line %d (%s) as number: %w", lineNumber+1, line, err)
		}

		prices := make([]int, 2001)
		prices[0] = secret.Price()

		for i := range 2000 {
			secret = secret.Next()
			prices[i+1] = secret.Price()
		}

		changes := make([]int, 2000)
		for i := range 2000 {
			changes[i] = prices[i+1] - prices[i]
		}

		for i := 0; i < 1996; i++ {
			window := ([4]int)(changes[i : i+5])

			if _, ok := bananasBySequence[window]; ok {
				continue
			}

			bananasBySequence[window] = prices[i+4]
		}

		bananasByMonkeyBySequence[lineNumber] = bananasBySequence
	}

	acc := make(map[[4]int]int, 0)
	for _, bananasBySequence := range bananasByMonkeyBySequence {
		for sequence, bananas := range bananasBySequence {
			acc[sequence] += bananas
		}
	}

	max := slices.Max(slices.Collect(maps.Values(acc)))

	return max, nil
}
