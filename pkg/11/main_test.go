package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`125 17`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 55312 {
		t.Fatalf("expected 55312, got %d", result)
	}
}
