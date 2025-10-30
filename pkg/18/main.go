package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines, 1024, 70)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(lines, 70)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %v\n", part2)
}

type Coordinate struct {
	X int
	Y int
}

func (a Coordinate) Equal(b Coordinate) bool {
	return a.X == b.X && a.Y == b.Y
}

type Item struct {
	value    Coordinate
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func ParseCoordinates(lines []string) ([]Coordinate, error) {
	coordinates := make([]Coordinate, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("expected 2 comma separated values on line %d: got %s", i+1, line)
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("x coordinate on line %d invalid: got %s", i+1, parts[0])
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("y coordinate on line %d invalid: got %s", i+1, parts[1])
		}

		coordinates[i] = Coordinate{x, y}
	}

	return coordinates, nil
}

func Part1(lines []string, bytes int, size int) (int, error) {
	coordinates, err := ParseCoordinates(lines)
	if err != nil {
		return 0, err
	}

	corrupted := make(map[Coordinate]struct{}, bytes)
	for i := range bytes {
		corrupted[coordinates[i]] = struct{}{}
	}

	end := Coordinate{size, size}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	costs := make(map[Coordinate]int, (size+1)*(size+1))

	for y := range size + 1 {
		for x := range size + 1 {
			costs[Coordinate{x, y}] = math.MaxInt32
		}
	}

	start := Coordinate{0, 0}
	costs[start] = 0
	heap.Push(&pq, &Item{value: start})

	for pq.Len() > 0 {
		from := heap.Pop(&pq).(*Item).value

		for _, direction := range []Coordinate{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
		} {
			to := Coordinate{
				direction.X + from.X,
				direction.Y + from.Y,
			}

			if _, ok := corrupted[to]; ok {
				continue
			}

			cost := costs[from] + 1

			if cost < costs[to] {
				costs[to] = cost
				if !to.Equal(end) {
					heap.Push(&pq, &Item{value: to, priority: cost})
				}
			}
		}
	}

	return costs[end], nil
}

func Part2(lines []string, size int) (Coordinate, error) {
	coordinates, err := ParseCoordinates(lines)
	if err != nil {
		return Coordinate{}, err
	}

	corrupted := make(map[Coordinate]struct{})

	start := Coordinate{0, 0}
	end := Coordinate{size, size}

	for _, coordinate := range coordinates {
		corrupted[coordinate] = struct{}{}
		visited := map[Coordinate]struct{}{}

		stack := []Coordinate{start}
		var traverseable bool

		for len(stack) > 0 {
			from := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			visited[from] = struct{}{}

			for _, direction := range []Coordinate{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
			} {
				to := Coordinate{
					direction.X + from.X,
					direction.Y + from.Y,
				}

				if _, ok := visited[to]; ok {
					continue
				}

				if to.X < 0 || to.X > size || to.Y < 0 || to.Y > size {
					continue
				}

				if _, ok := corrupted[to]; ok {
					continue
				}

				if to.Equal(end) {
					traverseable = true
					break
				} else {
					stack = append(stack, to)
				}
			}

		}

		if !traverseable {
			return coordinate, nil
		}
	}

	return Coordinate{}, fmt.Errorf("no blocking coordinate found")
}
