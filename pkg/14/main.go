package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines, 101, 103)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(lines, 101, 103)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

type Point struct {
	x int
	y int
}

type Robot struct {
	position Point
	velocity Point
}

var ROBOT_REGEX = regexp.MustCompile(`p\=(-?\d+),(-?\d+) v\=(-?\d+),(-?\d+)`)

func ParseRobots(lines []string) ([]Robot, error) {
	robots := make([]Robot, len(lines))

	for i, line := range lines {
		match := ROBOT_REGEX.FindStringSubmatch(line)
		if len(match) == 0 {
			return nil, fmt.Errorf("invalid robot %v", line)
		}

		values := make([]int, 4)
		for i := range 4 {
			v, err := strconv.Atoi(match[i+1])
			if err != nil {
				return nil, fmt.Errorf("invalid robot %v property %d: %w", line, i, err)
			}
			values[i] = v
		}

		robots[i] = Robot{
			position: Point{
				values[0], values[1],
			},
			velocity: Point{
				values[2], values[3],
			},
		}
	}

	return robots, nil
}

func Part1(lines []string, width int, height int) (int, error) {
	if width%2 == 0 || height%2 == 0 {
		return 0, fmt.Errorf("invalid width/height; must be odd")
	}

	robots, err := ParseRobots(lines)
	if err != nil {
		return 0, fmt.Errorf("robot parsing: %w", err)
	}

	const duration = 100

	quadrants := [2][2]int{
		{0, 0},
		{0, 0},
	}

	for _, robot := range robots {
		visited := map[Point]int{
			robot.position: 0,
		}
		current := robot.position

		for seconds := 1; seconds <= duration; seconds++ {
			next := Point{
				current.x + robot.velocity.x,
				current.y + robot.velocity.y,
			}

			// wrap position
			// assumption that no robot has a velocity greater than the width/height of the grid
			// as we'd need more complicated wrapping
			if next.x < 0 {
				next.x += width
			} else if next.x >= width {
				next.x -= width
			}

			if next.y < 0 {
				next.y += height
			} else if next.y >= height {
				next.y -= height
			}

			if _, ok := visited[next]; ok {
				// determine cycle length, and therefore final position
				cycle := seconds

				left := duration - seconds

				r := left % cycle
				if r == 0 {
					current = next
					break
				}

				var found bool
				for p, s := range visited {
					if s == r {
						found = true
						current = p
						break
					}
				}
				if !found {
					panic("unable to find final point from cycle")
				} else {
					break
				}

			}

			visited[next] = seconds

			current = next
		}

		// what quadrant are we in
		if current.x == width/2 || current.y == height/2 {
			// middle horizontally/vertically are excluded
			continue
		}

		var qx int
		if current.x > width/2 {
			qx = 1
		} else {
			qx = 0
		}
		var qy int
		if current.y > height/2 {
			qy = 1
		} else {
			qy = 0
		}

		quadrants[qy][qx] += 1
	}

	return quadrants[0][0] * quadrants[0][1] * quadrants[1][0] * quadrants[1][1], nil
}

func Part2(lines []string, width int, height int) (int, error) {
	robots, err := ParseRobots(lines)
	if err != nil {
		return 0, fmt.Errorf("robot parsing: %w", err)
	}

	for seconds := 1; ; seconds++ {
		current := map[Point]struct{}{}
		var overlap bool = false

		for i, robot := range robots {
			next := Point{
				robot.position.x,
				robot.position.y,
			}
			next.x += robot.velocity.x
			next.y += robot.velocity.y

			// wrap position
			// assumption that no robot has a velocity greater than the width/height of the grid
			// as we'd need more complicated wrapping
			if next.x < 0 {
				next.x += width
			} else if next.x >= width {
				next.x -= width
			}

			if next.y < 0 {
				next.y += height
			} else if next.y >= height {
				next.y -= height
			}

			if _, ok := current[next]; ok {
				overlap = true
			} else {
				current[next] = struct{}{}
			}

			robot.position = next
			robots[i] = robot
		}

		if !overlap {
			var b strings.Builder
			b.Grow(height * width)
			for y := range height {
				for x := range width {
					p := Point{x, y}
					if _, ok := current[p]; ok {
						b.WriteRune('.')
					} else {
						b.WriteRune(' ')
					}
				}
				b.WriteRune('\n')
			}
			fmt.Println(b.String())
			return seconds, nil
		}
	}
}
