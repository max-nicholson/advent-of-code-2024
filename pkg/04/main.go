package main

import (
	"fmt"
	"log"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

func Part1(lines []string) (int, error) {
	total := 0

	rows := len(lines)
	columns := len(lines[0])
	var directions = [][2]int{{1, -1}, {1, 1}, {-1, 1}, {-1, -1}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for row, chars := range lines {
		for column, c := range chars {
			if c != 'X' {
				continue
			}

			for _, direction := range directions {
				dx := direction[0]
				dy := direction[1]

				for i, c := range []byte{'M', 'A', 'S'} {
					y := row + dy*(i+1)
					x := column + dx*(i+1)

					if x < 0 || x > rows-1 {
						break
					}

					if y < 0 || y > columns-1 {
						break
					}

					if lines[y][x] != c {
						break
					}

					if i == 2 {
						total += 1
					}
				}
			}

		}
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	total := 0

	rows := len(lines)
	columns := len(lines[0])

	for row := 1; row < rows-1; row++ {
		for column := 1; column < columns-1; column++ {
			if lines[row][column] != 'A' {
				continue
			}

			if lines[row-1][column-1] == 'M' && lines[row+1][column-1] == 'M' && lines[row-1][column+1] == 'S' && lines[row+1][column+1] == 'S' {
				total += 1
			} else if lines[row-1][column-1] == 'S' && lines[row+1][column-1] == 'S' && lines[row-1][column+1] == 'M' && lines[row+1][column+1] == 'M' {
				total += 1
			} else if lines[row-1][column-1] == 'M' && lines[row+1][column-1] == 'S' && lines[row-1][column+1] == 'M' && lines[row+1][column+1] == 'S' {
				total += 1
			} else if lines[row-1][column-1] == 'S' && lines[row+1][column-1] == 'M' && lines[row-1][column+1] == 'S' && lines[row+1][column+1] == 'M' {
				total += 1
			}
		}
	}

	return total, nil
}
