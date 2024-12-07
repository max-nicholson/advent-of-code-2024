package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 2 {
		t.Fatalf("expected 2, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 4 {
		t.Fatalf("expected 4, got %d", result)
	}
}
