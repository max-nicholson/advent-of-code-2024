package main

import (
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
	"github.com/mowshon/iterium"
)

func main() {
	lines, err := lib.ReadLines("pkg/23/input.txt")
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
	fmt.Printf("part2: %s\n", part2)
}

type Computer struct {
	name      string
	connected map[string]struct{}
}

type NetworkMap struct {
	computers map[string]Computer
}

func (networkMap NetworkMap) Sets() [][]Computer {
	sets := make([][]Computer, 0)
	seen := map[string]struct{}{}

	for _, computer := range networkMap.computers {
		combinations := iterium.Combinations(slices.Collect(maps.Keys(computer.connected)), 2)
		for pair := range combinations.Chan() {
			names := []string{computer.name, pair[0], pair[1]}
			slices.Sort(names)
			id := strings.Join(names, ",")

			if _, ok := seen[id]; ok {
				continue
			}

			if _, ok := networkMap.computers[pair[0]].connected[pair[1]]; ok {
				sets = append(sets, []Computer{
					computer,
					networkMap.computers[pair[0]],
					networkMap.computers[pair[1]],
				})
				seen[id] = struct{}{}
			}
		}
	}

	return sets
}

type LANParty struct {
	computers []Computer
}

func (networkMap NetworkMap) LANParty() LANParty {
	maxSize := 0
	maxComputers := []Computer{}
	seen := map[string]struct{}{}

	for name, computer := range networkMap.computers {
		pool := slices.Collect(maps.Keys(computer.connected))
		pool = append(pool, name)

		for size := lib.Max(maxSize+1, 2); size < len(pool); size++ {
			var found = false
			permutations := iterium.Combinations(pool, size)

			for permutation := range permutations.Chan() {
				slices.Sort(permutation)
				id := strings.Join(permutation, ",")

				if _, ok := seen[id]; ok {
					continue
				}

				seen[id] = struct{}{}

				var set = true
				for i, name := range permutation {
					connected := networkMap.computers[name].connected

					for _, other := range permutation[i+1:] {
						if _, ok := connected[other]; !ok {
							set = false
							break
						}
					}

					if !set {
						break
					}
				}

				if !set {
					continue
				}

				found = true
				maxSize = size
				computers := make([]Computer, size)

				for i, name := range permutation {
					computers[i] = networkMap.computers[name]
				}

				maxComputers = computers
			}

			if !found {
				break
			}
		}
	}

	return LANParty{
		computers: maxComputers,
	}
}

func (party LANParty) Password() string {
	names := make([]string, len(party.computers))

	for i, computer := range party.computers {
		names[i] = computer.name
	}

	slices.Sort(names)

	return strings.Join(names, ",")
}

func NewNetworkMap(lines []string) NetworkMap {
	computers := make(map[string]Computer, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "-")
		a := parts[0]
		b := parts[1]

		{
			computer, ok := computers[a]
			if !ok {
				computer = Computer{
					name: a,
					connected: map[string]struct{}{
						b: {},
					},
				}
				computers[a] = computer
			} else {
				computer.connected[b] = struct{}{}
			}
		}

		{
			computer, ok := computers[b]
			if !ok {
				computer = Computer{
					name: b,
					connected: map[string]struct{}{
						a: {},
					},
				}
				computers[b] = computer
			} else {
				computer.connected[a] = struct{}{}
			}
		}
	}

	return NetworkMap{
		computers,
	}
}

func Part1(lines []string) (int, error) {
	total := 0

	networkMap := NewNetworkMap(lines)

	for _, set := range networkMap.Sets() {
		if slices.ContainsFunc(set, func(computer Computer) bool {
			return strings.HasPrefix(computer.name, "t")
		}) {
			total += 1
		}
	}

	return total, nil
}

func Part2(lines []string) (string, error) {
	networkMap := NewNetworkMap(lines)

	return networkMap.LANParty().Password(), nil
}
