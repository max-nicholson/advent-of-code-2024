package main

import (
	"fmt"
	"log"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

type Mode int

const (
	File Mode = iota + 1
	FreeSpace
)

func main() {
	lines, err := lib.ReadLines("pkg/09/input.txt")
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

func ParseLength(b byte) int {
	return int(rune(b) - '0')
}

func Part1(lines []string) (int, error) {
	diskMap := lines[0]
	length := len(diskMap)

	type State struct {
		index  int
		fileId int
		mode   Mode
		blocks int
	}

	left := State{index: 0, fileId: -1, mode: File}
	var right State
	if length%2 == 0 {
		// Last entry is free space, can be ignored
		right = State{index: length - 2, fileId: length / 2}
	} else {
		right = State{index: length - 1, fileId: length/2 + 1}
	}

	var checksum = 0
	var position = 0

	for left.index < right.index {
		if left.mode == File {
			if left.blocks > 0 {
				panic("partially consumed block in left File mode")
			}

			left.fileId += 1
			blocks := ParseLength(diskMap[left.index])

			for range blocks {
				checksum += left.fileId * position
				// fmt.Printf("%d * %d = %d\n", position, left.fileId, left.fileId*position)
				position += 1
			}

			left.mode = FreeSpace
			left.index += 1
		} else {
			var space int
			if left.blocks > 0 {
				// partially used free space
				space = left.blocks
			} else {
				space = ParseLength(diskMap[left.index])
			}

			var blocks int
			if right.blocks > 0 {
				// partially used file
				blocks = right.blocks
			} else {
				// new complete file block
				right.fileId -= 1
				blocks = ParseLength(diskMap[right.index])
			}

			for range lib.Min(blocks, space) {
				// fmt.Printf("%d * %d = %d\n", position, right.fileId, right.fileId*position)
				checksum += right.fileId * position
				position += 1
			}

			if blocks > space {
				// FreeSpace on left is fully consumed
				left.mode = File
				left.index += 1
				left.blocks = 0

				// Must have leftover file blocks on the right
				right.blocks = blocks - space
			} else if blocks < space {
				// File on right is fully consumed
				right.index -= 2
				right.blocks = 0

				// We must have leftover space on the left
				left.blocks = space - blocks
			} else {
				// Both left and right are fully consumed
				left.mode = File
				left.index += 1
				left.blocks = 0

				right.index -= 2
				right.blocks = 0
			}
		}
	}

	if right.blocks > 0 {
		// left caught up to right with some spare
		for range right.blocks {
			// fmt.Printf("%d * %d = %d\n", position, right.fileId, right.fileId*position)
			checksum += position * right.fileId
			position += 1
		}
		right.blocks = 0
	}

	return checksum, nil
}

func Part2(lines []string) (int, error) {
	diskMap := lines[0]
	length := len(diskMap)

	disk := make([]int, 0, length*4)
	var fileId = -1
	for i := 0; i < length; i++ {
		blocks := ParseLength(diskMap[i])
		if i%2 == 0 {
			fileId += 1
			for range blocks {
				disk = append(disk, fileId)
			}
		} else {
			for range blocks {
				disk = append(disk, -1)
			}
		}
	}

	var min = 0
	var minFileId = fileId
	positions := len(disk) - 1
	var i = positions
	for i >= 0 {
		fileId := disk[i]

		if fileId == -1 {
			// skip over free space
			i -= 1
			continue
		}

		// get number of blocks
		var blocks = 1
		for i-blocks >= 0 {
			if disk[i-blocks] == fileId {
				blocks += 1
			} else {
				break
			}
		}

		if fileId > minFileId {
			// we've already moved this file - can be skipped
			i -= blocks
			continue
		}

		// need to make space for this many blocks
		var space = 0
		var partialSpace = false
		var j int
		for j = min; j < i-blocks+1; j++ {
			if disk[j] != -1 {
				if space > 0 {
					// space has now ended, but is too small
					// reset the counter and keep looking
					space = 0
					// flag that we've skipped over some space that's too small
					// as subsequent files might fit here
					partialSpace = true
				} else if !partialSpace {
					// occupied space - don't bother checking this in future files
					min += 1
				}
			} else {
				space += 1
				if space == blocks {
					break
				}
			}
		}

		if space == blocks {
			// we've found somewhere to fit the file
			for b := range blocks {
				disk[j-b] = fileId
				disk[i-b] = -1
			}
		}

		minFileId = fileId
		i -= blocks
	}

	var checksum int
	for i, fileId := range disk {
		if fileId != -1 {
			checksum += i * fileId
		}
	}

	return checksum, nil
}
