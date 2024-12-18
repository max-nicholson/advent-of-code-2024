package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPart1(t *testing.T) {
	result, err := Part1(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`)
	if err != nil {
		t.Fatal(err)
	}
	if result != 2028 {
		t.Fatalf("expected 2028, got %d", result)
	}
}

func setupWarehouse(warehouse string) [][]rune {
	lines := strings.Split(warehouse, "\n")

	acc := make([][]rune, len(lines))

	for i, line := range lines {
		acc[i] = []rune(line)
	}

	return acc
}

func TestTryMoveBoxVertically(t *testing.T) {
	cases := []struct {
		name      string
		move      Point
		warehouse string
		want      string
	}{
		{
			name: "single column",
			move: Point{r: -1, c: 0},
			warehouse: `####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##...[].......[]..##
##[]##....[]......##
##[]......[]..[]..##
##..[][]..@[].[][]##
##........[]......##
####################`,
			want: `####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##...[]...[]..[]..##
##[]##....[]......##
##[]......@...[]..##
##..[][]...[].[][]##
##........[]......##
####################`,
		},
		{
			name: "fan out",
			move: Point{r: -1},
			warehouse: `##############
##......##..##
##..........##
##...[][]...##
##....[]....##
##.....@....##
##############`,
			want: `##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############`,
		},
		{
			name: "complex",
			move: Point{r: 1},
			warehouse: `##############
#..[].[].[][]#
#.....[]...[]#
#...@[][][]###
#..[][][][]..#
#.[][][][].###
#..[]....[][]#
#[][]..[]..###
#............#
#......[]....#
##############`,
			want: `##############
#..[].[].[][]#
#.....[]...[]#
#....[][][]###
#...@[][][]..#
#..[].[][].###
#.[][]...[][]#
#[][]..[]..###
#..[]........#
#......[]....#
##############`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			warehouse := setupWarehouse(c.warehouse)

			start, ok := FindRobot(warehouse)
			if !ok {
				t.Errorf("TryMoveBoxVertically() no robot found in warehouse\n%s", PrintWarehouse(warehouse))
			}

			TryMoveBoxVertically(warehouse, start, c.move)

			if !cmp.Equal(warehouse, setupWarehouse(c.want)) {
				t.Errorf("TryMoveBoxVertically() want \n%s \ngot \n%s", c.want, PrintWarehouse(warehouse))
			}
		})
	}
}

func TestTryMoveBoxHorizontally(t *testing.T) {
	cases := []struct {
		name      string
		move      Point
		warehouse string
		want      string
	}{
		{
			name: "one box",
			move: Point{r: 0, c: -1},
			warehouse: `####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##....[]@.....[]..##
##[]##....[]......##
##[]....[]....[]..##
##..[][]..[]..[][]##
##........[]......##
####################`,
			want: `####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##...[]@......[]..##
##[]##....[]......##
##[]....[]....[]..##
##..[][]..[]..[][]##
##........[]......##
####################`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			warehouse := setupWarehouse(c.warehouse)

			start, ok := FindRobot(warehouse)
			if !ok {
				t.Errorf("TryMoveBoxHorizontally() no robot found in warehouse\n%s", PrintWarehouse(warehouse))
			}

			TryMoveBoxHorizontally(warehouse, start, c.move)

			if !cmp.Equal(warehouse, setupWarehouse(c.want)) {
				t.Errorf("TryMoveBoxHorizontally() want %s got %s", c.want, PrintWarehouse(warehouse))
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "Small example",
			input: `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`,
			want: 618,
		},
		{
			name: "larger example",
			input: `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
`,
			want: 9021,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Part2(c.input)
			if err != nil {
				t.Error(err)
			}
			if got != c.want {
				t.Errorf("Part2() want %d got %d", c.want, got)
			}
		})
	}
}
