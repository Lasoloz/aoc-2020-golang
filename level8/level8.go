package main

import "fmt"

// Debug version
//type operation string
//
//const (
//	acc operation = "acc"
//	jmp           = "jmp"
//	nop           = "nop"
//)

// Trying out "iota"
type operation int

const (
	acc operation = iota
	jmp
	nop
)

type retCode int

const (
	cont retCode = iota
	loop
	exit
)

type instructionType struct {
	instruction operation
	value       int
}

type machineType struct {
	instructions   []instructionType
	inPath         []bool
	globalAcc      int
	instructionPtr int
}

func (receiver *machineType) step() (int, retCode) {
	if receiver.instructionPtr >= len(receiver.instructions) {
		return receiver.globalAcc, exit
	}

	if receiver.inPath[receiver.instructionPtr] {
		return receiver.globalAcc, loop
	}

	receiver.inPath[receiver.instructionPtr] = true
	curr := receiver.instructions[receiver.instructionPtr]

	switch curr.instruction {
	case nop:
		receiver.instructionPtr += 1
	case acc:
		receiver.globalAcc += curr.value
		receiver.instructionPtr += 1
	case jmp:
		receiver.instructionPtr += curr.value
	}

	return receiver.globalAcc, cont
}

func main() {
	instructions := readInstructions()

	fmt.Println(findAccAtFirstLoop(instructions))

	solution, ok := findAccByFixingInstruction(instructions)
	if !ok {
		fmt.Println("No solution for second part!")
	} else {
		fmt.Println(solution)
	}
}

func findAccAtFirstLoop(instructions []instructionType) int {
	machine := makeMachine(instructions)

	for {
		value, ret := machine.step()

		if ret == loop {
			return value
		}
	}
}

func findAccByFixingInstruction(instructions []instructionType) (int, bool) {
	for i := len(instructions) - 1; i >= 0; i-- {
		if result, ok := checkByTogglingInstruction(instructions, i); ok {
			return result, true
		}
	}

	return -1, false
}

func checkByTogglingInstruction(instructions []instructionType, ptr int) (int, bool) {
	if ok := toggleInstruction(instructions, ptr); !ok {
		return -1, false
	}
	defer toggleInstruction(instructions, ptr)

	machine := makeMachine(instructions)

	for {
		value, ret := machine.step()

		if ret == cont {
			continue
		}

		return value, ret == exit
	}
}

func toggleInstruction(instructions []instructionType, ptr int) bool {
	switch instructions[ptr].instruction {
	case jmp:
		instructions[ptr].instruction = nop
		return true
	case nop:
		instructions[ptr].instruction = jmp
		return true
	default:
		return false
	}
}

func readInstructions() []instructionType {
	machine := make([]instructionType, 0)

	for {
		var rawOp string
		var value int

		_, err := fmt.Scanln(&rawOp, &value)
		if err != nil {
			break
		}

		machine = append(machine, instructionType{readInstruction(rawOp), value})
	}

	return machine
}

func readInstruction(rawOp string) operation {
	switch rawOp {
	case "acc":
		return acc
	case "jmp":
		return jmp
	case "nop":
		return nop
	}
	return nop
}

func makeMachine(instructions []instructionType) machineType {
	return machineType{instructions, make([]bool, len(instructions)), 0, 0}
}
