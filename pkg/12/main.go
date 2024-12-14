package main

import (
	"fmt"
	"log"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

type Plane int

const (
	Horizontal Plane = iota + 1
	Vertical
)

type Edge int

const (
	Top Edge = iota + 1
	Bottom
	Left
	Right
)

func (e Edge) Plane() Plane {
	switch e {
	case Top:
		return Horizontal
	case Bottom:
		return Horizontal
	case Left:
		return Vertical
	case Right:
		return Vertical
	}
	panic(fmt.Sprintf("invalid edge %v", e))
}

func (e Edge) Deltas() []Point {
	switch e.Plane() {
	case Horizontal:
		return []Point{{0, 1}, {0, -1}}
	case Vertical:
		return []Point{{1, 0}, {-1, 0}}
	}
	panic("invalid plane")
}

func main() {
	lines, err := lib.ReadLines("pkg/12/input.txt")
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
	row    int
	column int
}

func (p Point) Adjacent(edge Edge) Point {
	switch edge {
	case Top:
		return Point{p.row - 1, p.column}
	case Bottom:
		return Point{p.row + 1, p.column}
	case Left:
		return Point{p.row, p.column - 1}
	case Right:
		return Point{p.row, p.column + 1}
	}
	panic("invalid edge")
}

var CARDINAL_DIRECTIONS = []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

type Region struct {
	plots map[Point]struct{}
	plant byte
}

func (r Region) Add(point Point) {
	r.plots[point] = struct{}{}
}

func (r Region) Area() int {
	return len(r.plots)
}

type Fence struct {
	start Point
	end   Point
}

func (f Fence) Plane() Plane {
	if f.start.row == f.end.row {
		return Horizontal
	} else {
		return Vertical
	}
}

func (r Region) Fences() map[Fence]struct{} {
	perimeter := map[Fence]struct{}{}

	for p := range r.plots {
		for _, delta := range CARDINAL_DIRECTIONS {
			neighbour := Point{row: delta.row + p.row, column: delta.column + p.column}

			if _, ok := r.plots[neighbour]; !ok {
				var fence Fence
				if delta.row == -1 {
					// Horizontal above
					fence.start = p
					fence.end = Point{p.row, p.column + 1}
				} else if delta.row == 1 {
					// Horizontal below
					fence.start = neighbour
					fence.end = Point{neighbour.row, neighbour.column + 1}
				} else if delta.column == -1 {
					// Vertical left
					fence.start = p
					fence.end = Point{p.row + 1, p.column}
				} else {
					// Vertical right
					fence.start = neighbour
					fence.end = Point{neighbour.row + 1, neighbour.column}
				}

				perimeter[fence] = struct{}{}
				continue
			}
		}
	}

	return perimeter
}

func (r Region) Perimeter() int {
	return len(r.Fences())
}

func (r Region) Sides() int {
	var sides int
	type visit struct {
		Point
		edge Edge
	}
	visited := map[visit]struct{}{}

	for p := range r.plots {
		var current Point
		for _, edge := range []Edge{Top, Bottom, Left, Right} {
			if _, ok := visited[visit{Point: p, edge: edge}]; ok {
				continue
			}

			if _, ok := r.plots[p.Adjacent(edge)]; ok {
				continue
			}

			visited[visit{Point: p, edge: edge}] = struct{}{}
			for _, delta := range edge.Deltas() {
				current = p
				for {
					next := Point{current.row + delta.row, current.column + delta.column}
					if _, ok := r.plots[next]; !ok {
						break
					}

					if _, ok := r.plots[next.Adjacent(edge)]; ok {
						break
					}

					visited[visit{Point: next, edge: edge}] = struct{}{}
					current = next
				}
			}
			sides += 1
		}
	}

	return sides
}

func NewRegion(plant byte) Region {
	return Region{plant: plant, plots: map[Point]struct{}{}}
}

type Garden struct {
	regions []Region
}

func NewGarden(grid []string) Garden {
	garden := Garden{regions: []Region{}}
	rows := len(grid)
	columns := len(grid[0])

	visited := map[Point]struct{}{}

	for r := range rows {
		for c := range columns {
			plant := grid[r][c]
			point := Point{row: r, column: c}
			if _, ok := visited[point]; ok {
				continue
			}

			region := NewRegion(plant)

			stack := []Point{point}
			for len(stack) > 0 {
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if grid[p.row][p.column] != plant {
					continue
				}

				region.Add(p)
				visited[p] = struct{}{}

				for _, delta := range CARDINAL_DIRECTIONS {
					next := Point{row: delta.row + p.row, column: delta.column + p.column}

					if next.row < 0 || next.row >= rows || next.column < 0 || next.column >= columns {
						continue
					}

					if _, ok := visited[next]; ok {
						continue
					}

					if grid[next.row][next.column] == plant {
						stack = append(stack, next)
					}
				}
			}

			garden.regions = append(garden.regions, region)
		}
	}

	return garden
}

func Part1(grid []string) (int, error) {
	totalPrice := 0

	garden := NewGarden(grid)

	for _, region := range garden.regions {
		totalPrice += region.Area() * region.Perimeter()
	}

	return totalPrice, nil
}

func Part2(grid []string) (int, error) {
	totalPrice := 0

	garden := NewGarden(grid)

	for _, region := range garden.regions {
		area := region.Area()
		sides := region.Sides()
		price := area * sides
		// fmt.Printf("%s = %d * %d = %d\n", string(region.plant), area, sides, price)
		totalPrice += price
	}

	return totalPrice, nil
}
