package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 41 {
		t.Fatalf("expected 41, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 6 {
		t.Fatalf("expected 6, got %d", result)
	}
}
