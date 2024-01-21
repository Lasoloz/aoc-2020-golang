package main

import (
	"fmt"
	"slices"
)

func main() {
	joltages := readJoltages()
	slices.Sort(joltages)

	fmt.Println(findMultipliedDistribution(joltages))
}

func findMultipliedDistribution(joltages []int) int {
	currentJoltage := 0

	jolt1Count := 0
	jolt3Count := 0

	for _, joltage := range joltages {
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

	jolt3Count++
	return jolt1Count * jolt3Count
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
