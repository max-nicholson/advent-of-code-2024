package main

import (
	"fmt"
	"log"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/10/input.txt")
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

func ParseGrid(lines []string) [][]int {
	rows := len(lines)
	columns := len(lines[0])
	grid := make([][]int, rows)

	for i, line := range lines {
		grid[i] = make([]int, columns)
		for j, c := range line {
			height := int(c - '0')
			grid[i][j] = height
		}
	}
	return grid
}

type Point struct {
	row    int
	column int
}

type Path struct {
	Point
	height int
}

func Part1(lines []string) (int, error) {
	total := 0

	grid := ParseGrid(lines)
	rows := len(grid)
	columns := len(grid[0])

	for r, row := range grid {
		for c, node := range row {
			if node != 0 {
				continue
			}

			seen := make(map[Point]struct{})
			stack := []Path{{Point: Point{row: r, column: c}, height: 0}}

			for len(stack) > 0 {
				path := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				for _, delta := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					row := delta.row + path.Point.row
					column := delta.column + path.column

					if row < 0 || row >= rows || column < 0 || column >= columns {
						continue
					}

					next := Point{row, column}

					if grid[row][column] != path.height+1 {
						continue
					}

					if _, ok := seen[next]; ok {
						continue
					}

					seen[next] = struct{}{}

					if grid[row][column] == 9 {
						total += 1
					} else {
						stack = append(stack, Path{Point: next, height: grid[row][column]})
					}
				}
			}
		}
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	total := 0

	grid := ParseGrid(lines)
	rows := len(grid)
	columns := len(grid[0])

	for r, row := range grid {
		for c, node := range row {
			if node != 0 {
				continue
			}

			stack := []Path{{Point: Point{row: r, column: c}, height: 0}}

			for len(stack) > 0 {
				path := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				for _, delta := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					row := delta.row + path.Point.row
					column := delta.column + path.column

					if row < 0 || row >= rows || column < 0 || column >= columns {
						continue
					}

					next := Point{row, column}

					if grid[row][column] != path.height+1 {
						continue
					}

					if grid[row][column] == 9 {
						total += 1
					} else {
						stack = append(stack, Path{Point: next, height: grid[row][column]})
					}
				}
			}
		}
	}

	return total, nil
}
