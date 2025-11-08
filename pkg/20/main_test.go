package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

	for i, tc := range []struct {
		minSaving int
		want      int
	}{
		{
			minSaving: 64,
			want:      1,
		},
		{
			minSaving: 40,
			want:      2,
		},
		{
			minSaving: 38,
			want:      3,
		},
		{
			minSaving: 36,
			want:      4,
		},
		{
			minSaving: 20,
			want:      5,
		},
		{
			minSaving: 12,
			want:      8,
		},
		{
			minSaving: 10,
			want:      10,
		},
		{
			minSaving: 8,
			want:      14,
		},
		{
			minSaving: 6,
			want:      16,
		},
		{
			minSaving: 4,
			want:      30,
		},
		{
			minSaving: 2,
			want:      44,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Part1(strings.Split(input, "\n"), tc.minSaving)
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
	input := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

	for i, tc := range []struct {
		minSaving int
		want      int
	}{
		{
			minSaving: 76,
			want:      3,
		},
		{
			minSaving: 74,
			want:      7,
		},
		{
			minSaving: 72,
			want:      29,
		},
		{
			minSaving: 70,
			want:      41,
		},
		{
			minSaving: 68,
			want:      55,
		},
		{
			minSaving: 66,
			want:      67,
		},
		{
			minSaving: 64,
			want:      86,
		},
		{
			minSaving: 62,
			want:      106,
		},
		{
			minSaving: 60,
			want:      129,
		},
		{
			minSaving: 58,
			want:      154,
		},
		{
			minSaving: 56,
			want:      193,
		},
		{
			minSaving: 54,
			want:      222,
		},
		{
			minSaving: 52,
			want:      253,
		},
		{
			minSaving: 50,
			want:      285,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Part2(strings.Split(input, "\n"), tc.minSaving)
			if err != nil {
				t.Error(err)
			}
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
