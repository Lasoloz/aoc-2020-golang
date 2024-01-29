package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const stop = math.MaxInt

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

	jump := simplifiedTimetable[0].value
	match := 0

	// This algorithm is found by trial and error, and it would probably break for non-prime
	// numbers. Maybe it is worth checking if it actually does.
	// Start by first value 'a' in timetable
	// Find first matching sequence by jumping 'a' increments, where the next value is contained
	// in the a,b series. For the example 7,13,... that is 77
	// Then start by incrementing by 7*13 jumps
	// Find the next matching a,b,c series, in this case that is 7,13,x,x,59 matching on a start
	// of 350. Then repeat until final value is found
	for timestamp := jump; timestamp < stop; timestamp += jump {
		ok := true

		for index, entry := range simplifiedTimetable {
			if (timestamp+entry.index)%entry.value != 0 {
				ok = false
				break
			} else if index > match {
				jump *= entry.value
				match = index
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
