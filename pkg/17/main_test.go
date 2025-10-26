package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	for i, tc := range []struct {
		input string
		want  string
	}{
		{
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
			want: "4,6,3,5,6,3,5,2,1,0",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Part1(strings.Split(tc.input, "\n"))
			if err != nil {
				t.Error(err)
			}
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
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
			input: `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`,
			want: 117440,
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
