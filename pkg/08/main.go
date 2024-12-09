package main

import (
	"fmt"
	"log"

	"github.com/max-nicholson/advent-of-code-2024/lib"
	"github.com/mowshon/iterium"
)

func main() {
	lines, err := lib.ReadLines("pkg/08/input.txt")
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

type Point struct {
	x int
	y int
}

func PrintAntennae(lines []string, antinodes map[Point]struct{}) {
	width := len(lines[0])

	for i, line := range lines {
		acc := make([]rune, width)
		for j, node := range line {
			if node != '.' {
				acc[j] = node
			} else {
				point := Point{j, i}
				if _, ok := antinodes[point]; ok {
					acc[j] = '#'
				} else {
					acc[j] = '.'
				}
			}
		}
		fmt.Println(string(acc))
	}
}

func ParseAntennae(lines []string) map[rune][]Point {
	antennae := make(map[rune][]Point)

	for r, line := range lines {
		for c, node := range line {
			if node == '.' {
				continue
			}

			point := Point{c, r}

			_, ok := antennae[node]
			if !ok {
				antennae[node] = []Point{point}
			} else {
				antennae[node] = append(antennae[node], point)
			}
		}
	}
	return antennae
}

func UniqueAntinodes(lines []string) map[Point]struct{} {
	antennae := ParseAntennae(lines)

	uniqueAntinodes := make(map[Point]struct{})

	rows := len(lines)
	columns := len(lines[0])

	for _, antenna := range antennae {
		if len(antenna) == 1 {
			continue
		}

		combinations := iterium.Combinations(antenna, 2)
		for pair := range combinations.Chan() {
			a := pair[0]
			b := pair[1]

			dy := lib.Abs(a.y - b.y)
			dx := lib.Abs(a.x - b.x)

			var first Point
			var second Point
			if a.x < b.x {
				first = Point{
					a.x - dx,
					a.y - dy,
				}
				second = Point{
					b.x + dx,
					b.y + dy,
				}
			} else {
				first = Point{
					a.x + dx,
					a.y - dy,
				}
				second = Point{
					b.x - dx,
					b.y + dy,
				}
			}

			if 0 <= first.x && first.x < columns && 0 <= first.y && first.y < rows {
				uniqueAntinodes[first] = struct{}{}
			}

			if 0 <= second.x && second.x < columns && 0 <= second.y && second.y < rows {
				uniqueAntinodes[second] = struct{}{}
			}
		}
	}

	return uniqueAntinodes
}

func UniqueAntinodesWithResonance(lines []string) map[Point]struct{} {
	antennae := ParseAntennae(lines)

	uniqueAntinodes := make(map[Point]struct{})

	rows := len(lines)
	columns := len(lines[0])

	for _, antenna := range antennae {
		if len(antenna) == 1 {
			continue
		}

		combinations := iterium.Combinations(antenna, 2)
		for pair := range combinations.Chan() {
			a := pair[0]
			b := pair[1]

			uniqueAntinodes[a] = struct{}{}
			uniqueAntinodes[b] = struct{}{}

			dy := lib.Abs(a.y - b.y)
			dx := b.x - a.x

			var next Point
			// Up
			next = a
			for {
				next = Point{next.x - dx, next.y - dy}
				if next.x < 0 || next.x >= columns || next.y < 0 || next.y >= rows {
					break
				}

				uniqueAntinodes[next] = struct{}{}
			}

			// Down
			next = b
			for {
				next = Point{next.x + dx, next.y + dy}
				if next.x < 0 || next.x >= columns || next.y < 0 || next.y >= rows {
					break
				}

				uniqueAntinodes[next] = struct{}{}
			}
		}
	}

	return uniqueAntinodes
}

func Part1(lines []string) (int, error) {
	uniqueAntinodes := UniqueAntinodes(lines)

	return len(uniqueAntinodes), nil
}

func Part2(lines []string) (int, error) {
	uniqueAntinodes := UniqueAntinodesWithResonance(lines)

	return len(uniqueAntinodes), nil
}
