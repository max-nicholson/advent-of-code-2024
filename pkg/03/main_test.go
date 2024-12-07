package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 161 {
		t.Fatalf("expected 161, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 48 {
		t.Fatalf("expected 48, got %d", result)
	}
}
