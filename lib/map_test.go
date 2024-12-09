package lib_test

import (
	"testing"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func TestIntersect(t *testing.T) {
	a := map[int]struct{}{
		1: struct{}{},
		2: struct{}{},
	}

	b := map[int]struct{}{
		2: struct{}{},
		3: struct{}{},
	}

	if !lib.Intersect(a, b) {
		t.Fatalf("want a and b to intersect")
	}
}

func TestNoIntersect(t *testing.T) {
	a := map[int]struct{}{
		1: struct{}{},
		2: struct{}{},
	}

	b := map[int]struct{}{
		3: struct{}{},
		4: struct{}{},
	}

	if lib.Intersect(a, b) {
		t.Fatalf("want a and b not to intersect")
	}
}
