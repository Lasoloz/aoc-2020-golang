package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instructionDef struct {
	address int
	value   string
}

type memoryDef struct {
	memoryMap map[int]int64
	maskForm  int64
	maskValue int64
}

func (m memoryDef) String() string {
	return fmt.Sprintf("{%v %b %b}", m.memoryMap, m.maskForm, m.maskValue)
}

var maskingInstRegex = regexp.MustCompile("mask = (.+)")
var memInstRegex = regexp.MustCompile("mem\\[(\\d+)] = (\\d+)")

func main() {
	instructions := readInstructions()

	fmt.Println(performInstructions(instructions))
}

func performInstructions(instructions []instructionDef) int64 {
	memory := memoryDef{make(map[int]int64), 0, 0}

	for _, instruction := range instructions {
		if instruction.address < 0 {
			setMask(&memory, instruction)
		} else {
			setMem(&memory, instruction)
		}
	}

	return sumMemoryMap(memory)
}

func setMask(m *memoryDef, instruction instructionDef) {
	maskForm, maskValue := int64(0), int64(0)
	for _, ch := range instruction.value {
		maskForm <<= 1
		maskValue <<= 1

		if ch == '1' || ch == '0' {
			maskValue += int64(ch - '0')
		} else {
			maskForm += 1
			// Mask format is inverse, keeping only original bits
		}
	}
	m.maskForm, m.maskValue = maskForm, maskValue
}

func setMem(m *memoryDef, instruction instructionDef) {
	// It should be safe to ignore this error
	value, _ := strconv.ParseInt(instruction.value, 10, 64)
	value &= m.maskForm
	m.memoryMap[instruction.address] = value | m.maskValue
}

func sumMemoryMap(memory memoryDef) int64 {
	sum := int64(0)
	for _, value := range memory.memoryMap {
		sum += value
	}
	return sum
}

func readInstructions() []instructionDef {
	reader := bufio.NewReader(os.Stdin)
	instructions := make([]instructionDef, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return instructions
		}

		instruction := readInstruction(strings.Trim(line, "\n"))
		if instruction == nil {
			_, _ = fmt.Fprintln(os.Stderr, "Warning: incorrect instruction line:", line)
			continue
		}
		instructions = append(instructions, *instruction)
	}
}

func readInstruction(line string) *instructionDef {
	matches := maskingInstRegex.FindStringSubmatch(line)
	if matches != nil && len(matches) == 2 {
		return &instructionDef{-1, matches[1]}
	}

	matches = memInstRegex.FindStringSubmatch(line)
	if matches != nil && len(matches) == 3 {
		address, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil
		}
		return &instructionDef{address, matches[2]}
	}

	return nil
}
