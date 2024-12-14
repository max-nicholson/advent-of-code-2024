package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 36 {
		t.Fatalf("expected 36, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 81 {
		t.Fatalf("expected 81, got %d", result)
	}
}
