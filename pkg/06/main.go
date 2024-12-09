package main

import (
	"fmt"
	"log"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Point struct {
	x, y int
}

type Step struct {
	x, y      int
	direction Direction
}

func (d Direction) Delta() Point {
	switch d {
	case North:
		return Point{0, -1}
	case East:
		return Point{1, 0}
	case South:
		return Point{0, 1}
	case West:
		return Point{-1, 0}
	}

	panic("unreachable")
}

func (d Direction) Rotate() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}

	panic("unreachable")
}

func main() {
	grid, err := lib.ReadLines("pkg/06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(grid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(grid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

func FindGuard(grid []string) (Point, error) {
	for y, line := range grid {
		for x, position := range line {
			if position == '^' {
				return Point{x, y}, nil
			}
		}
	}

	return Point{}, fmt.Errorf("guard not found")
}

func UniquePositions(grid []string, start Point) map[Point]struct{} {
	rows := len(grid)
	columns := len(grid[0])

	current := start
	var direction Direction = North

	visited := make(map[Point]struct{})
	visited[current] = struct{}{}

	for {
		delta := direction.Delta()

		x := current.x + delta.x
		y := current.y + delta.y

		if !(0 <= x && x < columns) {
			break
		}

		if !(0 <= y && y < rows) {
			break
		}

		if grid[y][x] == '#' {
			direction = direction.Rotate()
		} else {
			current.x = x
			current.y = y
			visited[current] = struct{}{}
		}
	}

	return visited
}

func Part1(grid []string) (int, error) {
	start, err := FindGuard(grid)

	if err != nil {
		return 0, err
	}

	visited := UniquePositions(grid, start)

	return len(visited), nil
}

func Part2(grid []string) (int, error) {
	total := 0

	rows := len(grid)
	columns := len(grid[0])

	start, err := FindGuard(grid)
	if err != nil {
		return 0, err
	}

	visited := UniquePositions(grid, start)

	// Obstacle only has a chance of creating an infinite loop if it's on the original guard route
	// (without any obstacles)
	for obstacle := range visited {
		if obstacle == start {
			// cannot put obstacle at start
			continue
		}

		path := make(map[Step]struct{})
		path[Step{start.x, start.y, North}] = struct{}{}

		current := start
		var direction Direction = North

		for {
			delta := direction.Delta()

			x := current.x + delta.x
			y := current.y + delta.y

			if !(0 <= x && x < columns) {
				// Guard left area on x-axis
				break
			}

			if !(0 <= y && y < rows) {
				// Guard left area on y-axis
				break
			}

			if grid[y][x] == '#' || (x == obstacle.x && y == obstacle.y) {
				direction = direction.Rotate()
			} else {
				step := Step{x, y, direction}
				if _, ok := path[step]; ok {
					total += 1
					// Guard has hit a loop
					break
				}
				current.x = x
				current.y = y
				path[step] = struct{}{}
			}
		}
	}

	return total, nil
}
