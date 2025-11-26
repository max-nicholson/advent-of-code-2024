package main

import (
	"container/heap"
	"fmt"
	"log"
	"maps"
	"math"
	"slices"
	"strconv"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/21/input.txt")
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

type Position struct {
	row    int
	column int
}

type Move struct {
	from rune
	to   rune
}

type NumericKeypad struct {
	position rune
}

func NewNumericKeypad() *NumericKeypad {
	return &NumericKeypad{
		position: 'A',
	}
}

func (k *NumericKeypad) Type(button rune) []string {
	lookup := GetNumericalKeyboardLookup()
	sequence := lookup[Move{from: k.position, to: button}]
	k.position = button

	return sequence
}

func Find(needle rune, haystack [][]rune) Position {
	for row := range len(haystack) {
		for column := range len(haystack[0]) {
			if haystack[row][column] == needle {
				return Position{row, column}
			}
		}
	}

	panic("unreachable")
}

func Dijkstra(grid [][]rune, start rune) map[Position]int {
	pq := make(lib.PriorityQueue[Position], 0)
	heap.Init(&pq)
	costs := make(map[Position]int, len(grid)*len(grid[0]))

	for row := range grid {
		for column := range len(grid[0]) {
			costs[Position{row, column}] = math.MaxInt32
		}
	}

	source := Find(start, grid)
	heap.Push(&pq, &lib.PriorityQueueItem[Position]{Value: source})

	costs[source] = 0
	for pq.Len() > 0 {
		from := heap.Pop(&pq).(*lib.PriorityQueueItem[Position]).Value

		for _, direction := range []Position{
			{0, -1},
			{-1, 0},
			{0, 1},
			{1, 0},
		} {
			to := Position{
				row:    direction.row + from.row,
				column: direction.column + from.column,
			}

			if to.row < 0 || to.row >= len(grid) || to.column < 0 || to.column >= len(grid[0]) {
				continue
			}

			if grid[to.row][to.column] == ' ' {
				continue
			}

			cost := costs[from] + 1

			if cost < costs[to] {
				costs[to] = cost
				heap.Push(&pq, &lib.PriorityQueueItem[Position]{Value: to, Priority: cost})
			}
		}
	}

	return costs
}

func UniqueButtons(grid [][]rune) map[rune]struct{} {
	buttons := make(map[rune]struct{}, len(grid)*len(grid[0]))

	for _, line := range grid {
		for _, v := range line {
			if v != ' ' {
				buttons[v] = struct{}{}
			}
		}
	}

	return buttons
}

func Paths(costs map[Position]int, positionLookup map[rune]Position, start rune, end rune) []string {
	if start == end {
		return []string{"A"}
	}

	startPosition := positionLookup[start]
	endPosition := positionLookup[end]

	stack := []struct {
		current Position
		path    []rune
	}{
		{
			current: endPosition,
			path:    []rune{'A'},
		},
	}

	candidates := make([][]rune, 0)

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		from := item.current

		for _, direction := range []Position{
			{0, -1},
			{-1, 0},
			{0, 1},
			{1, 0},
		} {
			to := Position{
				row:    direction.row + from.row,
				column: direction.column + from.column,
			}

			cost, ok := costs[to]
			if !ok {
				continue
			}

			if cost != costs[from]-1 {
				continue
			}

			var direction rune
			if to.row == from.row {
				if to.column == from.column-1 {
					direction = '>'
				} else {
					direction = '<'
				}
			} else {
				if to.row == from.row-1 {
					direction = 'v'
				} else {
					direction = '^'
				}
			}

			path := append(make([]rune, 0, len(item.path)), item.path...)
			path = append(path, direction)

			if to.column == startPosition.column && to.row == startPosition.row {
				candidates = append(candidates, path)
			} else {
				stack = append(stack, struct {
					current Position
					path    []rune
				}{
					to,
					path,
				})
			}
		}
	}

	paths := []string{}
	for _, c := range candidates {
		slices.Reverse(c)
		paths = append(paths, string(c))
	}

	return paths
}

var numericalKeyboardLookup map[Move][]string

func GetNumericalKeyboardLookup() map[Move][]string {
	if len(numericalKeyboardLookup) == 0 {
		grid := [][]rune{
			{'7', '8', '9'},
			{'4', '5', '6'},
			{'1', '2', '3'},
			{' ', '0', 'A'},
		}

		numericalKeyboardLookup = make(map[Move][]string, len(grid)*len(grid[0]))

		positionLookup := make(map[rune]Position, len(grid)*len(grid[0]))
		for row, line := range grid {
			for column, value := range line {
				positionLookup[value] = Position{row, column}
			}
		}

		buttons := UniqueButtons(grid)

		for start := range buttons {
			costs := Dijkstra(grid, start)

			for end := range buttons {
				numericalKeyboardLookup[Move{from: start, to: end}] = Paths(costs, positionLookup, start, end)
			}
		}
	}

	return numericalKeyboardLookup
}

var directionalKeypadLookup map[Move][]string

func GetDirectionalKeypadLookup() map[Move][]string {
	if len(directionalKeypadLookup) == 0 {
		grid := [][]rune{
			{' ', '^', 'A'},
			{'<', 'v', '>'},
		}
		directionalKeypadLookup = make(map[Move][]string, len(grid)*len(grid[0]))

		buttons := UniqueButtons(grid)
		positionLookup := make(map[rune]Position, len(grid)*len(grid[0]))
		for row, line := range grid {
			for column, value := range line {
				positionLookup[value] = Position{row, column}
			}
		}

		for start := range buttons {
			costs := Dijkstra(grid, start)

			for end := range buttons {
				directionalKeypadLookup[Move{from: start, to: end}] = Paths(costs, positionLookup, start, end)
			}
		}

		// for move, path := range directionalKeypadLookup {
		// 	fmt.Printf("%s -> %s = %s\n", string(move.from), string(move.to), path)
		// }
	}

	return directionalKeypadLookup
}

type Keypad interface {
	Type(button rune) []string
}

type DirectionalKeypad struct {
	position rune
}

func NewDirectionalKeypad() *DirectionalKeypad {
	return &DirectionalKeypad{
		position: 'A',
	}
}

func (k *DirectionalKeypad) Type(button rune) []string {
	lookup := GetDirectionalKeypadLookup()
	sequence := lookup[Move{from: k.position, to: button}]
	k.position = button

	return sequence
}

type Code string

func (c Code) Numeric() int {
	numeric, _ := strconv.Atoi(string(c)[:3])

	return numeric
}

func Permutations(sequence string, keypad Keypad) []string {
	var acc []string
	for _, button := range sequence {
		subsequences := keypad.Type(button)
		if len(acc) == 0 {
			acc = subsequences
		} else {
			sequences := make([]string, 0, len(acc)*len(subsequences))
			for _, head := range acc {
				for _, tail := range subsequences {
					sequences = append(sequences, head+tail)
				}
			}
			acc = sequences
		}
	}

	return acc
}

func (code Code) Sequence(keypads []Keypad) string {
	previousLevel := map[string]struct{}{
		string(code): {},
	}
	var currentLevel map[string]struct{}

	for _, keypad := range keypads {
		currentLevel = make(map[string]struct{}, len(previousLevel))

		for sequence := range previousLevel {
			acc := Permutations(sequence, keypad)
			for _, sequence := range acc {
				currentLevel[sequence] = struct{}{}
			}
		}

		min := len(slices.MinFunc(slices.Collect(maps.Keys(currentLevel)), func(a string, b string) int {
			return len(a) - len(b)
		}))

		for sequence := range currentLevel {
			if len(sequence) > min {
				delete(currentLevel, sequence)
			}
		}
		previousLevel = currentLevel
	}

	combinations := slices.Collect(maps.Keys(currentLevel))

	return slices.MinFunc(combinations, func(a string, b string) int {
		return len(a) - len(b)
	})
}

func (c Code) Complexity(keypads []Keypad) int {
	return c.Numeric() * len(c.Sequence(keypads))
}

func Part1(lines []string) (int, error) {
	total := 0

	keypads := []Keypad{
		NewNumericKeypad(),
		NewDirectionalKeypad(),
		NewDirectionalKeypad(),
	}

	for _, line := range lines {
		code := Code(line)
		total += code.Complexity(keypads)
	}

	return total, nil
}

func Pairs(s string) [][2]rune {
	pairs := make([][2]rune, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		pairs[i] = [2]rune{rune(s[i]), rune(s[i+1])}
	}

	return pairs
}

type SequenceCacheItem struct {
	start rune
	end   rune
	depth int
}

var sequenceLengthCache = map[SequenceCacheItem]int{}

func SequenceLength(start rune, end rune, depth int) int {
	l, ok := sequenceLengthCache[SequenceCacheItem{start, end, depth}]
	if ok {
		return l
	}

	lookup := GetDirectionalKeypadLookup()

	if depth == 1 {
		v := len(lookup[Move{start, end}][0])
		sequenceLengthCache[SequenceCacheItem{start, end, depth}] = v
		return v
	}

	min := math.MaxInt
	for _, sequence := range lookup[Move{start, end}] {
		length := 0

		for _, pair := range Pairs("A" + sequence) {
			a := pair[0]
			b := pair[1]

			length += SequenceLength(a, b, depth-1)
		}

		min = lib.Min(min, length)
	}

	sequenceLengthCache[SequenceCacheItem{start, end, depth}] = min
	return min
}

func Part2(lines []string) (int, error) {
	total := 0

	for _, line := range lines {
		numericKeypad := NewNumericKeypad()
		code := Code(line)

		permutations := Permutations(string(code), numericKeypad)

		min := math.MaxInt
		for _, permutation := range permutations {
			length := 0

			for _, pair := range Pairs("A" + permutation) {
				a := pair[0]
				b := pair[1]

				length += SequenceLength(a, b, 25)
			}

			min = lib.Min(min, length)
		}

		total += code.Numeric() * min
	}

	return total, nil
}
