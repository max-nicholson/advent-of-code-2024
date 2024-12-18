package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	input, err := lib.ReadFile("pkg/15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

func ParseInput(input string) ([][]rune, []Point, error) {
	parts := strings.Split(input, "\n\n")

	warehouse := [][]rune{}
	lines := strings.Split(parts[0], "\n")
	for i, line := range lines {
		warehouse = append(warehouse, make([]rune, len(line)))
		for j, c := range line {
			warehouse[i][j] = c
		}
	}

	movements := make([]Point, 0, len(parts[1]))
	lines = strings.Split(parts[1], "\n")
	for r, line := range lines {
		for c, m := range line {
			move := Point{}
			switch m {
			case '<':
				move.c = -1
			case '>':
				move.c = 1
			case '^':
				move.r = -1
			case 'v':
				move.r = 1
			default:
				return nil, nil, fmt.Errorf("invalid movement %s at row %d column %d", string(m), r, c)
			}
			movements = append(movements, move)
		}
	}

	return warehouse, movements, nil
}

type Point struct {
	r int
	c int
}

func FindRobot(warehouse [][]rune) (Point, bool) {
	for r, row := range warehouse {
		for c, coordinate := range row {
			if coordinate == '@' {
				return Point{r, c}, true
			}
		}
	}

	return Point{}, false
}

func Part1(input string) (int, error) {
	warehouse, movements, err := ParseInput(input)
	if err != nil {
		return 0, err
	}

	current, found := FindRobot(warehouse)
	if !found {
		return 0, fmt.Errorf("could not find robot @ in warehouse %v", warehouse)
	}

	for i, move := range movements {
		next := Point{
			current.r + move.r,
			current.c + move.c,
		}

		switch warehouse[next.r][next.c] {
		case '@':
			return 0, fmt.Errorf("invalid move at index %d, robot already at %v", next, i)
		case '#':
			// blocked
		case '.':
			warehouse[next.r][next.c] = '@'
			warehouse[current.r][current.c] = '.'
			current = next
		case 'O':
			end := Point{
				next.r + move.r,
				next.c + move.c,
			}
			for {
				contents := warehouse[end.r][end.c]
				if contents == '.' {
					warehouse[next.r][next.c] = '@'
					warehouse[current.r][current.c] = '.'
					warehouse[end.r][end.c] = 'O'
					current = next
					break
				} else if contents == '#' {
					break
				} else {
					end.r += move.r
					end.c += move.c
				}
			}
		default:
			return 0, fmt.Errorf("unexpected location %s at %v", string(warehouse[next.r][next.c]), next)
		}
	}

	var sum int

	for r, row := range warehouse {
		for c, coordinate := range row {
			if coordinate != 'O' {
				continue
			}

			sum += 100*r + c
		}
	}

	return sum, nil
}

func ExpandWarehouse(warehouse [][]rune) [][]rune {
	expanded := make([][]rune, len(warehouse))

	for r, row := range warehouse {
		expanded[r] = make([]rune, len(row)*2)
		for c, tile := range row {
			switch tile {
			case '#':
				expanded[r][2*c] = '#'
				expanded[r][2*c+1] = '#'
			case 'O':
				expanded[r][2*c] = '['
				expanded[r][2*c+1] = ']'
			case '.':
				expanded[r][2*c] = '.'
				expanded[r][2*c+1] = '.'
			case '@':
				expanded[r][2*c] = '@'
				expanded[r][2*c+1] = '.'
			default:
				panic(fmt.Sprintf("invalid tile %s at row %d column %d", string(tile), r, c))
			}
		}
	}

	return expanded
}

func TryMoveBoxHorizontally(warehouse [][]rune, start Point, move Point) bool {
	if move.r != 0 {
		panic("attempting horizontal move with vertical delta")
	}

	box := Point{
		r: start.r,
		c: start.c + move.c,
	}
	step := Point{
		c: move.c * 2,
	}

	next := Point{
		start.r,
		box.c + step.c,
	}

	if move.c == 1 && warehouse[box.r][box.c] != '[' {
		panic(fmt.Sprintf("attempting move to right against %s; not left edge of a box", string(warehouse[box.r][box.c])))
	}

	if move.c == -1 && warehouse[box.r][box.c] != ']' {
		panic(fmt.Sprintf("attempting move to left against %s; not right edge of a box", string(warehouse[box.r][box.c])))
	}

	var boxes int = 1
	for {
		tile := warehouse[next.r][next.c]
		if tile == '.' {
			// slide every box across by 1
			// start with the space and work backwards
			for i := 0; i < boxes*2; i++ {
				warehouse[next.r][next.c-i*move.c] = warehouse[next.r][next.c-((i+1)*move.c)]
			}
			warehouse[start.r][box.c] = '@'
			warehouse[start.r][start.c] = '.'
			return true
		} else if tile == '#' {
			return false
		} else {
			boxes += 1
			next.c += step.c
		}
	}
}

func FindBox(warehouse [][]rune, side Point) (Point, Point) {
	switch warehouse[side.r][side.c] {
	case '[':
		left := side
		right := Point{
			side.r,
			side.c + 1,
		}
		if warehouse[right.r][right.c] != ']' {
			panic(fmt.Sprintf("expected right side of box; got %s", string(warehouse[side.r][side.c])))
		}
		return left, right
	case ']':
		left := Point{
			side.r,
			side.c - 1,
		}
		right := side
		if warehouse[left.r][left.c] != '[' {
			panic(fmt.Sprintf("expected left side of box; got %s", string(warehouse[side.r][side.c])))
		}
		return left, right
	default:
		panic(fmt.Sprintf("attempting vertical move against %s; not edge of a box", string(warehouse[side.r][side.c])))
	}
}

func TryMoveBoxVertically(warehouse [][]rune, start Point, move Point) bool {
	if move.c != 0 {
		panic("attempting vertical move with horizontal delta")
	}

	next := Point{
		start.r + move.r,
		start.c,
	}

	left, right := FindBox(warehouse, next)

	// NB: Using a map[Point]struct{} rather than []Point because we can have fan-in
	// [][]
	//  []
	// where we don't want to overcount (i.e. count same box twice), or fan-out
	//  []
	// [][]
	// where we don't want to undercount (only count one of the two boxes)
	// Using a set means it's easier to iterate across a frontier and account for every box exactly once

	rows := []map[Point]struct{}{}
	rows = append(rows, map[Point]struct{}{
		left:  {},
		right: {},
	})
	// Only care that the top edge of the boxes can move
	// e.g.
	// []
	// [][]
	//  []
	// Only need to check the `[]` on the top row, and the right of `[][]` for collisions
	frontier := map[int]int{
		left.c:  left.r,
		right.c: right.r,
	}

	for {
		row := map[Point]struct{}{}
		nextFrontier := map[int]int{}
		for c, r := range frontier {
			next := Point{
				r + move.r,
				c,
			}
			tile := warehouse[next.r][next.c]
			if tile == '#' {
				return false
			} else if tile == '.' {
				// OK, but need to check rest of frontier
			} else if tile == '[' || tile == ']' {
				previous := warehouse[r][c]
				if tile == previous {
					// []
					// []
					// Box is inline

					nextFrontier[next.c] = next.r
					row[next] = struct{}{}
				} else {
					// Box is staggered with another box
					// []
					//  []
					left, right := FindBox(warehouse, next)
					nextFrontier[left.c] = left.r
					nextFrontier[right.c] = right.r
					row[left] = struct{}{}
					row[right] = struct{}{}
				}
			} else {
				panic(fmt.Sprintf("unexpected tile %s", string(tile)))
			}
		}
		if len(nextFrontier) == 0 {
			for _, row := range slices.Backward(rows) {
				for point := range row {
					warehouse[point.r+move.r][point.c] = warehouse[point.r][point.c]
					warehouse[point.r][point.c] = '.'
				}
			}

			warehouse[start.r][start.c] = '.'
			if next.c == left.c {
				warehouse[left.r][left.c] = '@'
				warehouse[right.r][right.c] = '.'
			} else {
				warehouse[left.r][left.c] = '.'
				warehouse[right.r][right.c] = '@'
			}
			return true
		}

		frontier = nextFrontier
		rows = append(rows, row)
	}
}

func PrintWarehouse(warehouse [][]rune) string {
	var b strings.Builder
	b.Grow(len(warehouse) * (len(warehouse[0]) + 1))

	for _, row := range warehouse {
		for _, tile := range row {
			b.WriteRune(tile)
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func PrintMove(move Point) string {
	if move.c == 0 {
		if move.r == -1 {
			return "^"
		} else {
			return "v"
		}
	} else {
		if move.c == -1 {
			return "<"
		} else {
			return ">"
		}
	}
}

func CountWarehouse(warehouse [][]rune) (walls int, boxes int, empty int) {
	for r, row := range warehouse {
		for c, tile := range row {
			if tile == '#' {
				walls += 1
			} else if tile == '[' && c < len(row) && warehouse[r][c+1] == ']' {
				boxes += 1
			} else if tile == '.' {
				empty += 1
			}
		}
	}

	return walls, boxes, empty
}

func Part2(input string) (int, error) {
	warehouse, movements, err := ParseInput(input)
	if err != nil {
		return 0, err
	}

	warehouse = ExpandWarehouse(warehouse)

	sw, sb, se := CountWarehouse(warehouse)

	// fmt.Println(PrintWarehouse(warehouse))

	current, found := FindRobot(warehouse)
	if !found {
		return 0, fmt.Errorf("could not find robot @ in warehouse %v", warehouse)
	}

	for i, move := range movements {
		// fmt.Println(PrintMove(move))

		next := Point{
			current.r + move.r,
			current.c + move.c,
		}

		switch warehouse[next.r][next.c] {
		case '@':
			return 0, fmt.Errorf("invalid move at index %d, robot already at %v", next, i)
		case '#':
			// blocked
		case '.':
			warehouse[next.r][next.c] = '@'
			warehouse[current.r][current.c] = '.'
			current = next
		case '[':
			fallthrough
		case ']':
			if move.r == 0 {
				if ok := TryMoveBoxHorizontally(warehouse, current, move); ok {
					current = next
				}
			} else {
				if ok := TryMoveBoxVertically(warehouse, current, move); ok {
					current = next
				}
			}
		default:
			return 0, fmt.Errorf("unexpected location %s at %v", string(warehouse[next.r][next.c]), next)
		}

		// fmt.Println(PrintWarehouse(warehouse))

		if w, b, e := CountWarehouse(warehouse); w != sw || b != sb || e != se {
			panic(fmt.Sprintf("change in contents of warehouse from walls=%d, boxes=%d, empty=%d to walls=%d, boxes=%d, empty=%d", sw, sb, se, w, b, e))
		}
	}

	var sum int

	for r, row := range warehouse {
		for c, coordinate := range row {
			if coordinate != '[' {
				continue
			}

			sum += 100*r + c
		}
	}

	return sum, nil
}
