package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type wrappedInt struct {
	value int
}

func (receiver *wrappedInt) String() string {
	if receiver == nil {
		return "{ }"
	}
	return fmt.Sprintf("{%d}", receiver.value)
}

func main() {
	departure, timetable, ok := readDepartureAndTimetable()

	if !ok {
		fmt.Println("Incorrect input")
		return
	}

	fmt.Println(differenceMultipliedById(departure, timetable))
}

func differenceMultipliedById(departure int, timetable []*wrappedInt) int {
	lowestDiff := math.MaxInt
	lowestId := -1
	for _, time := range timetable {
		if time == nil {
			continue
		}

		diff := time.value - (departure % time.value)
		if diff < lowestDiff {
			lowestDiff = diff
			lowestId = time.value
		}
	}

	return lowestId * lowestDiff
}

func readDepartureAndTimetable() (departure int, timetable []*wrappedInt, ok bool) {
	_, dErr := fmt.Scanln(&departure)
	if dErr != nil {
		return
	}

	var rawTimetable string
	_, tErr := fmt.Scanln(&rawTimetable)
	if tErr != nil {
		return
	}

	timetable = make([]*wrappedInt, 0)

	for _, rawVal := range strings.Split(rawTimetable, ",") {
		val, atoiErr := strconv.Atoi(rawVal)
		if atoiErr != nil {
			timetable = append(timetable, nil)
		} else {
			timetable = append(timetable, &wrappedInt{val})
		}
	}

	ok = true
	return
}
