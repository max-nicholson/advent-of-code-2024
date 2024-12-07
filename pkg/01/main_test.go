package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`3   4
4   3
2   5
1   3
3   9
3   3`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 11 {
		t.Fatalf("expected 11, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`3   4
4   3
2   5
1   3
3   9
3   3`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 31 {
		t.Fatalf("expected 31, got %d", result)
	}
}
