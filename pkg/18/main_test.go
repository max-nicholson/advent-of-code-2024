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
		bytes int
		size  int
	}{
		{
			input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`,
			want:  22,
			bytes: 12,
			size:  6,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Part1(strings.Split(tc.input, "\n"), tc.bytes, tc.size)
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
		want  Coordinate
		size  int
	}{
		{
			input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`,
			want: Coordinate{6, 1},
			size: 6,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Part2(strings.Split(tc.input, "\n"), tc.size)
			if err != nil {
				t.Error(err)
			}
			if !got.Equal(tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
