package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`2333133121414131402`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 1928 {
		t.Fatalf("expected 1928, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`2333133121414131402`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 2858 {
		t.Fatalf("expected 2858, got %d", result)
	}
}
