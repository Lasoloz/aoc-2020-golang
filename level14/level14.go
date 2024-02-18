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
	memoryMap  map[int]int
	maskForm   int
	maskValues []int
}

func (m memoryDef) String() string {
	return fmt.Sprintf("{%v %b %b}", m.memoryMap, m.maskForm, m.maskValues)
}

var maskingInstRegex = regexp.MustCompile("mask = (.+)")
var memInstRegex = regexp.MustCompile("mem\\[(\\d+)] = (\\d+)")

func main() {
	instructions := readInstructions()

	fmt.Println(performInstructions(instructions, setSimpleMask, setSimpleMem))
	fmt.Println(performInstructions(instructions, setAddressMask, setMemOfFloatingAddresses))
}

func performInstructions(
	instructions []instructionDef,
	setMask func(m *memoryDef, instruction instructionDef),
	setMem func(m *memoryDef, instruction instructionDef),
) int {
	memory := memoryDef{make(map[int]int), 0, make([]int, 0)}

	for _, instruction := range instructions {
		if instruction.address < 0 {
			setMask(&memory, instruction)
		} else {
			setMem(&memory, instruction)
		}
	}

	return sumMemoryMap(memory)
}

func setSimpleMask(m *memoryDef, instruction instructionDef) {
	maskForm, maskValue := 0, 0
	for _, ch := range instruction.value {
		maskForm <<= 1
		maskValue <<= 1

		if ch == '1' || ch == '0' {
			maskValue += int(ch - '0')
		} else {
			maskForm += 1
			// Mask format is inverse, keeping only original bits
		}
	}
	m.maskForm = maskForm
	m.maskValues = []int{maskValue}
}

func setSimpleMem(m *memoryDef, instruction instructionDef) {
	if len(m.maskValues) < 1 {
		_, _ = fmt.Fprintln(os.Stderr, "Warning: empty mask slice!")
		return
	}
	// It should be safe to ignore this error
	value, _ := strconv.Atoi(instruction.value)
	maskValue := m.maskValues[0]
	value &= m.maskForm
	m.memoryMap[instruction.address] = value | maskValue
}

func setAddressMask(m *memoryDef, instruction instructionDef) {
	maskForm := 0
	maskValues := make([]int, 1)

	for _, ch := range instruction.value {
		maskForm <<= 1

		if ch == 'X' {
			maskValues = duplicateMasksWithVariedLSB(maskValues)
		} else if ch == '1' {
			maskValues = shiftAndSetLSB(maskValues, 1)
		} else {
			maskForm |= 1
			maskValues = shiftAndSetLSB(maskValues, 0)
		}
	}

	m.maskForm = maskForm
	m.maskValues = maskValues
}

func setMemOfFloatingAddresses(m *memoryDef, instruction instructionDef) {
	// It should be safe to ignore this error
	value, _ := strconv.Atoi(instruction.value)
	baseAddress := instruction.address
	for _, maskValue := range m.maskValues {
		address := (baseAddress & m.maskForm) | maskValue
		m.memoryMap[address] = value
	}
}

func duplicateMasksWithVariedLSB(masks []int) []int {
	result := make([]int, len(masks)*2)
	for i := range masks {
		result[2*i] = masks[i] << 1
		result[2*i+1] = (masks[i] << 1) | 1
	}
	return result
}

func shiftAndSetLSB(masks []int, lsbValue int) []int {
	for i := range masks {
		masks[i] = (masks[i] << 1) | lsbValue
	}
	return masks
}

func sumMemoryMap(memory memoryDef) int {
	sum := 0
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
