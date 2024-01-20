package main

import "fmt"

// Debug version
type operation string

const (
	acc operation = "acc"
	jmp           = "jmp"
	nop           = "nop"
)

// Trying out "iota"
//type operation int
//
//const (
//	acc operation = iota
//	jmp
//	nop
//)

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

func (receiver *machineType) step() (int, bool) {
	if receiver.inPath[receiver.instructionPtr] {
		return receiver.globalAcc, false
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

	return receiver.globalAcc, true
}

func main() {
	instructions := readInstructions()

	fmt.Println(findAccAtFirstLoop(instructions))
}

func findAccAtFirstLoop(instructions []instructionType) int {
	machine := makeMachine(instructions)

	for {
		value, ok := machine.step()

		if !ok {
			return value
		}
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
