package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

type Point struct {
	row    int
	column int
}

func (a Point) Equal(b Point) bool {
	return a.row == b.row && a.column == b.column
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    Node // The value of the item; arbitrary.
	priority int  // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
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

func main() {
	lines, err := lib.ReadLines("pkg/16/input.txt")
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

type Grid [][]rune

func parseGrid(lines []string) Grid {
	grid := make(Grid, len(lines))

	for row_index, row := range lines {
		grid[row_index] = make([]rune, len(row))
		for column_index, cell := range row {
			grid[row_index][column_index] = cell
		}
	}

	return grid
}
func (g Grid) Find(target rune) (Point, error) {
	for i, row := range g {
		for j, cell := range row {
			if cell == target {
				return Point{row: i, column: j}, nil
			}
		}
	}

	return Point{}, fmt.Errorf("cell not found")
}

type Direction Point

func (a Direction) Equal(b Direction) bool {
	return a.row == b.row && a.column == b.column
}

func (d Direction) Reverse() Direction {
	return Direction{row: -d.row, column: -d.column}
}

var directions = []Direction{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type Node struct {
	point     Point
	direction Direction
}

type Edge struct {
	node   Node
	weight int
}

func (point Point) Move(direction Direction) Point {
	return Point{
		row:    point.row + direction.row,
		column: point.column + direction.column,
	}
}

func (point Point) MinCost(costs map[Node]int) int {
	min := math.MaxInt32
	for _, direction := range directions {
		min = lib.Min(min, costs[Node{point: point, direction: direction}])
	}
	return min
}

func (g Grid) Edges(from Node) []Edge {
	edges := make([]Edge, 0)

	for _, direction := range []Direction{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	} {
		to := from.point.Move(direction)

		if g[to.row][to.column] == '#' {
			continue
		}

		node := Node{
			point:     to,
			direction: direction,
		}

		if direction.Equal(from.direction) {
			edges = append(edges, Edge{
				node:   node,
				weight: 1,
			})
		} else if direction.Equal(from.direction.Reverse()) {
			edges = append(edges, Edge{
				node:   node,
				weight: 2001,
			})
		} else {
			edges = append(edges, Edge{
				node:   node,
				weight: 1001,
			})
		}
	}

	return edges
}

func dijkstra(grid Grid, start Point) map[Node]int {
	costs := make(map[Node]int, len(grid)*len(grid[0]))
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for i, row := range grid {
		for j := range row {
			for _, direction := range directions {
				costs[Node{point: Point{row: i, column: j}, direction: direction}] = math.MaxInt32
			}
		}
	}

	costs[Node{point: start, direction: Direction{0, 1}}] = 0
	heap.Push(&pq, &Item{value: Node{point: start, direction: Direction{0, 1}}})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item).value

		for _, edge := range grid.Edges(u) {
			v := edge.node
			weight := edge.weight

			if costs[u]+weight < costs[v] {
				costs[v] = costs[u] + weight
				heap.Push(&pq, &Item{value: v, priority: costs[v]})
			}

		}
	}

	return costs
}

func Part1(lines []string) (int, error) {
	grid := parseGrid(lines)

	start, err := grid.Find('S')
	if err != nil {
		return 0, fmt.Errorf("start not found")
	}

	end, err := grid.Find('E')
	if err != nil {
		return 0, fmt.Errorf("end not found")
	}

	costs := dijkstra(grid, start)

	return end.MinCost(costs), nil
}

func Part2(lines []string) (int, error) {
	grid := parseGrid(lines)

	start, err := grid.Find('S')
	if err != nil {
		return 0, fmt.Errorf("start not found")
	}

	end, err := grid.Find('E')
	if err != nil {
		return 0, fmt.Errorf("end not found")
	}

	costs := dijkstra(grid, start)

	min := end.MinCost(costs)

	visited := map[Node]struct{}{}
	queue := []Node{}

	for _, direction := range directions {
		node := Node{point: end, direction: direction}
		if costs[node] == min {
			queue = append(queue, node)
			visited[node] = struct{}{}
		}
	}

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, prevDirection := range directions {
			previous := Node{
				point:     current.point.Move(current.direction.Reverse()),
				direction: prevDirection,
			}

			var cost int
			if current.direction.Equal(prevDirection) {
				cost = 1
			} else if current.direction.Equal(prevDirection.Reverse()) {
				cost = 2001
			} else {
				cost = 1001
			}

			if costs[previous]+cost == costs[current] {
				if _, seen := visited[previous]; !seen {
					visited[previous] = struct{}{}
					queue = append(queue, previous)
				}
			}
		}
	}

	points := map[Point]struct{}{}
	for node := range visited {
		points[node.point] = struct{}{}
	}

	return len(points), nil
}
