package main

import (
	"fmt"
	"strconv"
	"strings"
)

type directionType int

const (
	east directionType = iota
	north
	west
	south
)

type instructionType struct {
	instrType byte
	value     int
}

type shipStateType struct {
	x, y      int
	direction directionType
}

type shipWithWaypointType struct {
	x, y   int
	wx, wy int
}

func main() {
	instructions := readInstructions()

	fmt.Println(tracePath(instructions).manhattanPosition())
	fmt.Println(traceWaypointPath(instructions).manhattanPosition())
}

func (receiver *shipStateType) manhattanPosition() int {
	return abs(receiver.x) + abs(receiver.y)
}

func (receiver *shipStateType) mapMovement(dir directionType, value int) {
	switch dir {
	case east:
		receiver.x += value
	case west:
		receiver.x -= value
	case north:
		receiver.y += value
	case south:
		receiver.y -= value
	default:
		fmt.Println("Warning: bad direction")
	}
}

func (receiver *shipStateType) rotatePositive(units int) {
	receiver.direction = (receiver.direction + directionType(units)) % 4
}

func (receiver *shipStateType) rotateNegative(units int) {
	receiver.direction = (receiver.direction + 16 - directionType(units)) % 4
}

func tracePath(instructions []instructionType) *shipStateType {
	shipState := shipStateType{0, 0, east}

	for _, instruction := range instructions {
		switch instruction.instrType {
		case 'F':
			shipState.mapMovement(shipState.direction, instruction.value)
		case 'E':
			shipState.mapMovement(east, instruction.value)
		case 'N':
			shipState.mapMovement(north, instruction.value)
		case 'W':
			shipState.mapMovement(west, instruction.value)
		case 'S':
			shipState.mapMovement(south, instruction.value)
		case 'L':
			shipState.rotatePositive(instruction.value / 90)
		case 'R':
			shipState.rotateNegative(instruction.value / 90)
		default:
			fmt.Println("Warning: bad instruction")
		}
	}

	return &shipState
}

func (receiver *shipWithWaypointType) manhattanPosition() int {
	return abs(receiver.x) + abs(receiver.y)
}

func (receiver *shipWithWaypointType) moveToWaypoint(value int) {
	receiver.x += receiver.wx * value
	receiver.y += receiver.wy * value
}

func (receiver *shipWithWaypointType) moveWaypoint(dir directionType, value int) {
	switch dir {
	case east:
		receiver.wx += value
	case west:
		receiver.wx -= value
	case north:
		receiver.wy += value
	case south:
		receiver.wy -= value
	default:
		fmt.Println("Warning: bad waypoint direction")
	}
}

func (receiver *shipWithWaypointType) rotatePositive(units int) {
	for i := 0; i < units; i++ {
		wx, wy := receiver.wx, receiver.wy
		receiver.wx = -wy
		receiver.wy = wx
	}
}

func (receiver *shipWithWaypointType) rotateNegative(units int) {
	for i := 0; i < units; i++ {
		wx, wy := receiver.wx, receiver.wy
		receiver.wx = wy
		receiver.wy = -wx
	}
}

func traceWaypointPath(instructions []instructionType) *shipWithWaypointType {
	shipState := shipWithWaypointType{0, 0, 10, 1}

	for _, instruction := range instructions {
		switch instruction.instrType {
		case 'F':
			shipState.moveToWaypoint(instruction.value)
		case 'E':
			shipState.moveWaypoint(east, instruction.value)
		case 'N':
			shipState.moveWaypoint(north, instruction.value)
		case 'W':
			shipState.moveWaypoint(west, instruction.value)
		case 'S':
			shipState.moveWaypoint(south, instruction.value)
		case 'L':
			shipState.rotatePositive(instruction.value / 90)
		case 'R':
			shipState.rotateNegative(instruction.value / 90)
		default:
			fmt.Println("Warning: bad instruction")
		}
	}

	return &shipState
}

func readInstructions() []instructionType {
	instructions := make([]instructionType, 0)

	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil || len(line) < 2 {
			return instructions
		}

		line = strings.Trim(line, "\n")

		instrType := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Warning, couldn't convert string to numeric", err)
			return instructions
		}

		instructions = append(instructions, instructionType{instrType, value})
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
