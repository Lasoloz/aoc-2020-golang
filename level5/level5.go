package main

import (
	"fmt"
	"strconv"
	"strings"
)

type seat struct {
	row    int
	column int
	id     int
}

func main() {
	passes := readBoardingPasses()

	fmt.Println(findHighestId(passes))
}

func findHighestId(passes []seat) int {
	var maxPassId int

	for _, pass := range passes {
		if pass.id > maxPassId {
			maxPassId = pass.id
		}
	}

	return maxPassId
}

func readBoardingPasses() []seat {
	seats := make([]seat, 0)

	for {
		var pass string
		if _, err := fmt.Scanln(&pass); err != nil {
			break
		}

		seats = append(seats, parseSeat(pass))
	}

	return seats
}

func parseSeat(pass string) seat {
	rawRow := pass[:7]
	rawCol := pass[7:]
	convertedRow := strings.ReplaceAll(strings.ReplaceAll(rawRow, "F", "0"), "B", "1")
	convertedCol := strings.ReplaceAll(strings.ReplaceAll(rawCol, "L", "0"), "R", "1")
	row, _ := strconv.ParseInt(convertedRow, 2, 0)
	col, _ := strconv.ParseInt(convertedCol, 2, 0)

	return seat{int(row), int(col), int(row*8 + col)}
}
