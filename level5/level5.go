package main

import (
	"fmt"
	"sort"
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
	fmt.Println(findMissingId(passes)) // Printing the array and some string replacement gave me quicker answer tho...
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

func findMissingId(passes []seat) any {
	ids := make([]int, 0)

	for _, pass := range passes {
		ids = append(ids, pass.id)
	}

	sort.Ints(ids)

	missingId := 28

	for _, id := range ids[1:] {
		if id-missingId == 1 {
			missingId = id
		} else {
			missingId++
			break
		}
	}

	return missingId
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
