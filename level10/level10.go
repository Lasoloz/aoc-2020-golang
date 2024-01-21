package main

import (
	"fmt"
	"sort"
)

func main() {
	joltages := readJoltages()
	adjustedJoltages := sortAndAppendMinMaxJoltages(joltages)

	fmt.Println(findMultipliedDistribution(adjustedJoltages))
	fmt.Println(countArrangements(adjustedJoltages))
}

func findMultipliedDistribution(joltages []int) int {
	currentJoltage := joltages[0]

	jolt1Count := 0
	jolt3Count := 0

	for _, joltage := range joltages[1:] {
		diff := joltage - currentJoltage

		if diff > 3 {
			fmt.Println("Warning: difference is too high!", diff)
		}

		if diff == 1 {
			jolt1Count++
		} else if diff == 3 {
			jolt3Count++
		}

		currentJoltage = joltage
	}

	return jolt1Count * jolt3Count
}

func countArrangements(joltages []int) int {
	length := len(joltages)
	pathCounts := make([]int, length)
	pathCounts[length-1] = 1 // End has one possible path

	for i := length - 2; i >= 0; i-- {
		pathCount := 0
		current := joltages[i]

		for j, joltage := range joltages[i+1 : min(length, i+4)] {
			if joltage-current > 3 {
				break
			}

			pathCountIndex := i + j + 1
			pathCount += pathCounts[pathCountIndex]
		}

		pathCounts[i] = pathCount
	}

	return pathCounts[0]
}

func readJoltages() []int {
	buf := make([]int, 0)

	for {
		var value int
		_, err := fmt.Scanln(&value)

		if err != nil {
			return buf
		}

		buf = append(buf, value)
	}
}

func sortAndAppendMinMaxJoltages(joltages []int) []int {
	adjustedJoltages := make([]int, len(joltages)+1, len(joltages)+2)
	copy(adjustedJoltages[1:], joltages)
	sort.Ints(adjustedJoltages)
	adjustedJoltages = append(adjustedJoltages, adjustedJoltages[len(joltages)]+3)
	return adjustedJoltages
}
