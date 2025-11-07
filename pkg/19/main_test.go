package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	for i, tc := range []struct {
		input string
		want  int
	}{
		{
			input: `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`,
			want: 6,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Part1(strings.Split(tc.input, "\n"))
			if err != nil {
				t.Error(err)
			}
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	for i, tc := range []struct {
		input string
		want  int
	}{
		{
			input: `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`,
			want: 16,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Part2(strings.Split(tc.input, "\n"))
			if err != nil {
				t.Error(err)
			}
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
