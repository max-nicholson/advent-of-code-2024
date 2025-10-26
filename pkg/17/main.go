package main

import (
	"fmt"
	"iter"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2024/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %s\n", part1)

	part2, err := Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

type Register struct {
	Value int
}

type Opcode int

const (
	adv Opcode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func ParseOpcode(raw string) (Opcode, error) {
	opcode, err := strconv.ParseInt(raw, 10, 4)
	if err != nil {
		return 0, err
	}
	return Opcode(opcode), nil
}

type Operand int

func ParseOperand(raw string) (Operand, error) {
	operand, err := strconv.ParseInt(raw, 10, 4)
	if err != nil {
		return 0, err
	}

	return Operand(operand), nil
}

func (o Operand) Combo(registers Registers) int {
	switch o {
	case 0, 1, 2, 3:
		return int(o)
	case 4:
		return registers.A.Value
	case 5:
		return registers.B.Value
	case 6:
		return registers.C.Value
	}

	panic(fmt.Sprintf("invalid operand %d", o))
}

type Instruction struct {
	opcode  Opcode
	operand Operand
}

type Registers struct {
	A Register
	B Register
	C Register
}

type Program struct {
	registers    Registers
	instructions []Instruction
}

func ParseRegister(input string) (Register, error) {
	value, err := strconv.Atoi(input[12:])
	if err != nil {
		return Register{}, err
	}
	return Register{value}, nil
}

func ParseInstructions(input string) ([]Instruction, error) {
	rawInstructions := strings.Split(input[9:], ",")
	instructions := make([]Instruction, len(rawInstructions)/2)

	instructionPointer := 0
	for rawOpcode, rawOperand := range lib.Pairs(slices.Values(rawInstructions)) {
		opcode, err := ParseOpcode(rawOpcode)
		if err != nil {
			return instructions, fmt.Errorf("failed to parse opcode at instruction pointer %d: %w", instructionPointer, err)
		}
		operand, err := ParseOperand(rawOperand)
		if err != nil {
			return instructions, fmt.Errorf("failed to parse operand at instruction pointer %d: %w", instructionPointer+1, err)
		}

		instructions[instructionPointer/2] = Instruction{
			opcode,
			operand,
		}
		instructionPointer += 2
	}

	return instructions, nil
}

func ParseProgram(input []string) (Program, error) {
	registers := Registers{}
	a, err := ParseRegister(input[0])
	if err != nil {
		return Program{}, fmt.Errorf("failed to parse A register: %w", err)
	}
	registers.A = a

	b, err := ParseRegister(input[1])
	if err != nil {
		return Program{}, fmt.Errorf("failed to parse B register: %w", err)
	}
	registers.B = b

	c, err := ParseRegister(input[2])
	if err != nil {
		return Program{}, fmt.Errorf("failed to parse C register: %w", err)
	}
	registers.C = c

	instructions, err := ParseInstructions(input[4])
	if err != nil {
		return Program{}, fmt.Errorf("failed to parse program: %w", err)
	}

	return Program{
		registers,
		instructions,
	}, nil
}

func (p Program) Output() iter.Seq[int] {
	return func(yield func(int) bool) {
		registers := p.registers

		instructionPointer := 0
		for instructionPointer < len(p.instructions)*2 {
			instruction := p.instructions[instructionPointer/2]

			switch instruction.opcode {
			case adv:
				registers.A.Value = int(registers.A.Value / lib.PowInt(2, instruction.operand.Combo(registers)))
			case bxl:
				registers.B.Value ^= int(instruction.operand)
			case bst:
				registers.B.Value = instruction.operand.Combo(registers) % 8
			case bxc:
				registers.B.Value ^= registers.C.Value
			case out:
				s := instruction.operand.Combo(registers) % 8
				if !yield(s) {
					return
				}
			case bdv:
				registers.B.Value = int(registers.A.Value / lib.PowInt(2, instruction.operand.Combo(registers)))
			case cdv:
				registers.C.Value = int(registers.A.Value / lib.PowInt(2, instruction.operand.Combo(registers)))
			}

			if instruction.opcode == jnz && registers.A.Value != 0 {
				instructionPointer = int(instruction.operand)
			} else {
				instructionPointer += 2
			}

		}
	}
}

func (p Program) InstructionsByOpcode(opcode Opcode) iter.Seq[Instruction] {
	return lib.Filter(slices.Values(p.instructions), func(i Instruction) bool {
		return i.opcode == adv
	})
}

func Part1(lines []string) (string, error) {
	program, err := ParseProgram(lines)
	if err != nil {
		return "", err
	}

	output := []string{}
	for o := range program.Output() {
		output = append(output, strconv.Itoa(o))
	}

	return strings.Join(output, ","), nil
}

func Part2(lines []string) (int, error) {
	program, err := ParseProgram(lines)
	if err != nil {
		return 0, err
	}

	o := strings.Split(lines[4][9:], ",")
	output := make([]int, len(o))
	for i, v := range o {
		output[i], err = strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
	}

	// Expect program to have a certain shape, otherwise it's probably not solveable?
	// Or at least, the solution is much more complicated
	if lastInstruction := program.instructions[len(program.instructions)-1]; lastInstruction.opcode != jnz || lastInstruction.operand != 0 {
		return 0, fmt.Errorf("can only handle programs which loop entirely")
	}

	advInstructions := slices.Collect(program.InstructionsByOpcode(adv))
	if len(advInstructions) != 1 {
		return 0, fmt.Errorf("can only handle programs with a single `adv` instruction")
	}
	if advInstructions[0].operand > 3 {
		return 0, fmt.Errorf("can only handle programs where `adv` operand is a literal")
	}

	if len(slices.Collect(program.InstructionsByOpcode(out))) != 1 {
		return 0, fmt.Errorf("can only handle programs with a single `out` instruction")
	}

	// Only care about the loop body
	program.instructions = program.instructions[:len(program.instructions)-1]

	var solve func(instruction int, target int) int
	solve = func(instruction int, target int) int {
		if instruction < 0 {
			return target
		}

		// Only 8 possible values to check with 8-bit integers
		for t := range 8 {
			program.registers.A.Value = (target << 3) + t
			for v := range program.Output() {
				if v == output[instruction] {
					sub := solve(instruction-1, program.registers.A.Value)
					if sub != 0 {
						return sub
					}
				}
			}
		}

		return 0
	}

	return solve(len(output)-1, 0), nil
}
