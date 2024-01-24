package main

import "fmt"

const maxIterations = 10_000

type seatConfigType int

const (
	floor seatConfigType = iota
	emptySeat
	occupSeat
)

type point struct {
	x, y int
}

type countRulesType struct {
	seatCounter func([][]seatConfigType, int, int) int
	tolerance   int
}

var adjacentDeltas = []point{
	{-1, -1}, {-1, 00}, {-1, 01},
	{00, -1} /* 0, 0*/, {00, 01},
	{01, -1}, {01, 00}, {01, 01},
}

func main() {
	seatConfig := readSeatConfig()

	adjacencyRules := countRulesType{countAdjacentSeats, 4}
	directionalRules := countRulesType{countDirectionallyVisibleSeats, 5}

	fmt.Println(solve(seatConfig, adjacencyRules))
	fmt.Println(solve(seatConfig, directionalRules))
}

func solve(seatConfig [][]seatConfigType, adjacencyRules countRulesType) any {
	return countOccupiedSeats(iterateUntilStable(
		copyMatrix(seatConfig),
		adjacencyRules,
	))
}

func copyMatrix[Type any](matrix [][]Type) [][]Type {
	result := make([][]Type, 0, len(matrix))

	for _, row := range matrix {
		resultRow := make([]Type, len(row))
		copy(resultRow, row)
		result = append(result, resultRow)
	}

	return result
}

func countOccupiedSeats(config [][]seatConfigType) any {
	count := 0

	for _, row := range config {
		for _, value := range row {
			if value == occupSeat {
				count++
			}
		}
	}

	return count
}

func iterateUntilStable(seatConfig [][]seatConfigType, rules countRulesType) [][]seatConfigType {
	for i := 0; i < maxIterations; i++ {
		if !iterate(seatConfig, rules) {
			break
		}
	}

	return seatConfig
}

func iterate(seatConfig [][]seatConfigType, rules countRulesType) bool {
	posToToggle := findToggleablePositions(seatConfig, rules)

	for _, pos := range posToToggle {
		seatConfig[pos.y][pos.x] = toggleSeat(seatConfig[pos.y][pos.x])
	}

	return len(posToToggle) > 0
}

func findToggleablePositions(seatConfig [][]seatConfigType, rules countRulesType) []point {
	posToToggle := make([]point, 0)

	for y := range seatConfig {
		for x, value := range seatConfig[y] {
			if value == floor {
				continue
			}

			occup := rules.seatCounter(seatConfig, y, x)
			if value == emptySeat && occup == 0 {
				posToToggle = append(posToToggle, point{x, y})
			} else if value == occupSeat && occup >= rules.tolerance {
				posToToggle = append(posToToggle, point{x, y})
			}
		}
	}
	return posToToggle
}

func countAdjacentSeats(seatConfig [][]seatConfigType, y int, x int) int {
	occup := 0

	for _, delta := range adjacentDeltas {
		aX, aY := x+delta.x, y+delta.y
		if !isValid(aX, aY, seatConfig) {
			continue
		}

		val := seatConfig[aY][aX]
		if val == occupSeat {
			occup++
		}
	}

	return occup
}

func countDirectionallyVisibleSeats(seatConfig [][]seatConfigType, y int, x int) int {
	occup := 0

	for _, delta := range adjacentDeltas {
		cx, cy := x, y
		for {
			cx += delta.x
			cy += delta.y

			if !isValid(cx, cy, seatConfig) || seatConfig[cy][cx] == emptySeat {
				break
			}

			if seatConfig[cy][cx] == occupSeat {
				occup++
				break
			}
		}
	}
	return occup
}

func isValid(aX, aY int, config [][]seatConfigType) bool {
	return !(aX < 0 || aY < 0 || aX >= len(config[0]) || aY >= len(config))
}

func toggleSeat(configType seatConfigType) seatConfigType {
	if configType == emptySeat {
		return occupSeat
	} else {
		return emptySeat
	}
}

func readSeatConfig() [][]seatConfigType {
	result := make([][]seatConfigType, 0)

	for {
		var row string
		_, err := fmt.Scanln(&row)
		if err != nil {
			return result
		}

		result = append(result, processRowOfSeats(row))
	}
}

func processRowOfSeats(rawRow string) []seatConfigType {
	row := make([]seatConfigType, 0)

	for _, char := range rawRow {
		switch char {
		case 'L':
			row = append(row, emptySeat)
		default:
			row = append(row, floor)
		}
	}

	return row
}
