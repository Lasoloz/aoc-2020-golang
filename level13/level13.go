package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const stop = 100_000_000_000

type wrappedInt struct {
	value int
}

func (receiver *wrappedInt) String() string {
	if receiver == nil {
		return "{ }"
	}
	return fmt.Sprintf("{%d}", receiver.value)
}

type timetableEntry struct {
	index int
	value int
}

func main() {
	departure, timetable, ok := readDepartureAndTimetable()

	if !ok {
		fmt.Println("Incorrect input")
		return
	}

	fmt.Println(differenceMultipliedById(departure, timetable))
	fmt.Println(findIncrementalDeparture(timetable))
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

func findIncrementalDeparture(timetable []*wrappedInt) int {
	simplifiedTimetable := simplifyTimetable(timetable)
	for timestamp := 0; timestamp < stop; timestamp++ {
		ok := true

		if timestamp%100_000_000 == 0 {
			fmt.Print(".")
		}

		for _, entry := range simplifiedTimetable {
			if (timestamp+entry.index)%entry.value != 0 {
				ok = false
				break
			}
		}

		if ok {
			return timestamp
		}
	}

	return 0
}

func simplifyTimetable(timetable []*wrappedInt) []timetableEntry {
	result := make([]timetableEntry, 0)

	for index, time := range timetable {
		if time != nil {
			result = append(result, timetableEntry{index, time.value})
		}
	}

	return result
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
