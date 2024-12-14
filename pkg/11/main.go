package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/11/input.txt")
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

func ParseStones(line string) []int {
	parts := strings.Split(line, " ")
	stones := make([]int, len(parts))
	for i, v := range parts {
		stone, _ := strconv.Atoi(v)
		stones[i] = stone
	}
	return stones
}

func Next(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if len(strconv.Itoa(stone))%2 == 0 {
		digits := strconv.Itoa(stone)
		left, _ := strconv.Atoi(digits[:len(digits)/2])
		right, _ := strconv.Atoi(digits[len(digits)/2:])

		return []int{left, right}
	} else {
		return []int{stone * 2024}
	}
}

func Part1(lines []string) (int, error) {
	stones := ParseStones(lines[0])

	for range 25 {
		next := make([]int, 0, len(stones))

		for _, stone := range stones {
			next = append(next, Next(stone)...)
		}

		stones = next
	}

	return len(stones), nil
}

type CacheKey struct {
	stone int
	times int
}

func Blink(cache map[CacheKey]int, stone int, times int) int {
	key := CacheKey{stone, times}
	v, ok := cache[key]
	if ok {
		return v
	}

	if times == 1 {
		l := len(Next(stone))
		cache[key] = l
		return l
	}

	next := Next(stone)

	var sum int
	for _, s := range next {
		t := Blink(cache, s, times-1)
		cache[CacheKey{s, times - 1}] = t
		sum += t
	}
	cache[key] = sum
	return sum
}

func Part2(lines []string) (int, error) {
	stones := ParseStones(lines[0])

	var total int
	cache := map[CacheKey]int{}

	for _, stone := range stones {
		total += Blink(cache, stone, 75)
	}

	return total, nil
}
