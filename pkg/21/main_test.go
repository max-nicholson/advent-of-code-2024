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
			input: `029A
980A
179A
456A
379A`,
			want: 126384,
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
