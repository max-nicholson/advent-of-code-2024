package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/20/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(lines, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

type Point struct {
	row, column int
}

func (a Point) Equal(b Point) bool {
	return a.row == b.row && a.column == b.column
}

type Item[T any] struct {
	value    T
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type Racetrack struct {
	grid  []string
	start Point
	end   Point
}

func (racetrack Racetrack) InBounds(p Point) bool {
	return !(p.row < 0 || p.row >= len(racetrack.grid) || p.column < 0 || p.column >= len(racetrack.grid[0]))
}

func ParseRacetrack(lines []string) Racetrack {
	var start Point
	for r, line := range lines {
		for c, cell := range line {
			if cell == 'S' {
				start = Point{r, c}
				break
			}
		}
	}

	var end Point
	for r, line := range lines {
		for c, cell := range line {
			if cell == 'E' {
				end = Point{r, c}
				break
			}
		}
	}

	racetrack := Racetrack{
		grid:  lines,
		start: start,
		end:   end,
	}

	return racetrack
}

var directions = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func Part1(lines []string, minSaving int) (int, error) {
	racetrack := ParseRacetrack(lines)

	pq := make(PriorityQueue[Point], 0)
	heap.Init(&pq)
	costs := make(map[Point]int, len(racetrack.grid)*len(racetrack.grid[0]))

	for row, line := range racetrack.grid {
		for column := range line {
			costs[Point{row: row, column: column}] = math.MaxInt32
		}
	}

	costs[racetrack.start] = 0
	heap.Push(&pq, &Item[Point]{value: racetrack.start})

	for pq.Len() > 0 {
		from := heap.Pop(&pq).(*Item[Point]).value

		for _, direction := range directions {
			to := Point{
				row:    direction.row + from.row,
				column: direction.column + from.column,
			}

			if !racetrack.InBounds(to) {
				continue
			}

			if racetrack.grid[to.row][to.column] == '#' {
				continue
			}

			cost := costs[from] + 1

			if cost < costs[to] {
				costs[to] = cost
				if !to.Equal(racetrack.end) {
					heap.Push(&pq, &Item[Point]{value: to, priority: cost})
				}
			}
		}
	}

	visited := map[Point]int{racetrack.end: 0}
	queue := []Point{racetrack.end}

	for len(queue) > 0 {
		from := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, direction := range directions {
			point := Point{row: direction.row + from.row, column: direction.column + from.column}

			if costs[point] == costs[from]-1 {
				queue = append(queue, point)
				visited[point] = costs[point]
			}
		}
	}

	delete(visited, racetrack.start)

	cheats := 0
	for from := range visited {
		for _, direction := range directions {
			to := Point{row: direction.row + from.row, column: direction.column + from.column}

			if !racetrack.InBounds(to) {
				continue
			}

			// Must be a wall to start a cheat
			if racetrack.grid[to.row][to.column] != '#' {
				continue
			}

			var otherDirections []Point
			for _, b := range directions {
				if b.row == -direction.row && b.column == -direction.column {
					continue
				}
				otherDirections = append(otherDirections, b)
			}

			for _, direction := range otherDirections {
				second := Point{row: direction.row + to.row, column: direction.column + to.column}

				if !racetrack.InBounds(second) {
					continue
				}

				// "At most 2" means this doesn't _have_ to be a wall

				// If we've not made a worthwhile saving, don't bother exploring this path further
				if racetrack.grid[second.row][second.column] != '#' {
					saving := costs[from] - 2 - costs[second]
					if saving >= minSaving {
						cheats += 1
						continue
					}
				}

				if second.Equal(racetrack.end) {
					continue
				}

				var otherDirections []Point
				for _, b := range directions {
					if b.row == -direction.row && b.column == -direction.column {
						continue
					}
					otherDirections = append(otherDirections, b)
				}

				for _, direction := range otherDirections {
					third := Point{row: direction.row + second.row, column: direction.column + second.column}

					if !racetrack.InBounds(second) {
						continue
					}

					// Used up our cheats
					if racetrack.grid[to.row][to.column] == '#' {
						continue
					}

					moves := 3
					saving := costs[from] - moves - costs[third]
					if saving >= minSaving {
						cheats += 1
					}
				}
			}
		}
	}

	return cheats, nil
}

func Part2(lines []string, minSaving int) (int, error) {
	racetrack := ParseRacetrack(lines)

	costs := make(map[Point]int, len(racetrack.grid)*len(racetrack.grid[0]))

	costs[racetrack.start] = 0
	queue := []Point{racetrack.start}

	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]

		for _, to := range []Point{
			{from.row + 1, from.column},
			{from.row - 1, from.column},
			{from.row, from.column + 1},
			{from.row, from.column - 1},
		} {

			if !racetrack.InBounds(to) {
				continue
			}

			if racetrack.grid[to.row][to.column] == '#' {
				continue
			}

			if _, ok := costs[to]; ok {
				continue
			}

			costs[to] = costs[from] + 1

			queue = append(queue, to)
		}
	}

	cheats := 0
	for row := range len(racetrack.grid) {
		for column := range len(racetrack.grid[0]) {
			if racetrack.grid[row][column] == '#' {
				continue
			}

			from := Point{row, column}

			for moves := 2; moves < 21; moves++ {
				for dr := range moves + 1 {
					dc := moves - dr
					for to := range map[Point]struct{}{
						{from.row + dr, from.column + dc}: {},
						{from.row + dr, from.column - dc}: {},
						{from.row - dr, from.column + dc}: {},
						{from.row - dr, from.column - dc}: {},
					} {

						if !racetrack.InBounds(to) || racetrack.grid[to.row][to.column] == '#' {
							continue
						}

						if (costs[from] - costs[to]) >= (minSaving + moves) {
							cheats += 1
						}
					}
				}
			}
		}
	}

	return cheats, nil
}
