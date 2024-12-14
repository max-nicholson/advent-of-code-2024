package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 1930 {
		t.Fatalf("expected 1930, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	type test struct {
		input string
		want  int
	}
	tests := []test{
		{input: `AAAA
BBCD
BBCC
EEEC`,
			want: 80},
		{input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			want: 236},
		{input: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			want: 368},
		{input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`, want: 1206},
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Part2(strings.Split(tc.input, "\n"))
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Fatalf("test: %d; expected %d, got %d", i+1, tc.want, got)
			}
		})
	}
}
