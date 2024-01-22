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

var adjacentDeltas = []point{
	{-1, -1}, {-1, 00}, {-1, 01},
	{00, -1} /* 0, 0*/, {00, 01},
	{01, -1}, {01, 00}, {01, 01},
}

func main() {
	seatConfig := readSeatConfig()
	iterateUntilStable(seatConfig)

	fmt.Println(countOccupiedSeats(seatConfig))
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

func iterateUntilStable(seatConfig [][]seatConfigType) {
	for i := 0; i < maxIterations; i++ {
		if !iterate(seatConfig) {
			break
		}
	}
}

func iterate(seatConfig [][]seatConfigType) bool {
	posToToggle := findToggleablePositions(seatConfig)

	for _, pos := range posToToggle {
		seatConfig[pos.y][pos.x] = toggleSeat(seatConfig[pos.y][pos.x])
	}

	return len(posToToggle) > 0
}

func findToggleablePositions(seatConfig [][]seatConfigType) []point {
	posToToggle := make([]point, 0)

	for y := range seatConfig {
		for x, value := range seatConfig[y] {
			if value == floor {
				continue
			}

			_, occup := countAdjacentSeats(seatConfig, y, x)
			if value == emptySeat && occup == 0 {
				posToToggle = append(posToToggle, point{x, y})
			} else if value == occupSeat && occup >= 4 {
				posToToggle = append(posToToggle, point{x, y})
			}
		}
	}
	return posToToggle
}

func countAdjacentSeats(seatConfig [][]seatConfigType, y int, x int) (empty, occup int) {
	for _, delta := range adjacentDeltas {
		aX, aY := x+delta.x, y+delta.y
		if !isValid(aX, aY, seatConfig) {
			continue
		}

		val := seatConfig[aY][aX]
		if val == emptySeat {
			empty++
		} else if val == occupSeat {
			occup++
		}
	}

	return
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
