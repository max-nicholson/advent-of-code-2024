package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(``, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 0 {
		t.Fatalf("expected 0, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(``, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 0 {
		t.Fatalf("expected 0, got %d", result)
	}
}
