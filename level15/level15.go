package main

import (
	"fmt"
	"strconv"
	"strings"
)

const searchLen1 = 2020
const searchLen2 = 30_000_000
const buildLen = searchLen2

func main() {
	startingNumbers := readStartingNumbers()

	if len(startingNumbers) == 0 {
		fmt.Println("Empty starting numbers list")
		return
	}

	result := buildList(startingNumbers)
	fmt.Println(result[searchLen1-1])
	fmt.Println(result[searchLen2-1])
}

func buildList(startingNumbers []int) []int {
	numbers := make([]int, buildLen)
	previousIndices := make(map[int]int)
	copy(numbers, startingNumbers)

	for i := 0; i < len(startingNumbers)-1; i++ {
		num := startingNumbers[i]
		previousIndices[num] = i
	}

	for i := len(startingNumbers); i < buildLen; i++ {
		lastIndex := i - 1
		last := numbers[lastIndex]
		diff := 0

		if prevIndex, ok := previousIndices[last]; ok {
			diff = lastIndex - prevIndex
		}

		numbers[i] = diff
		previousIndices[last] = i - 1
	}

	return numbers
}

func readStartingNumbers() []int {
	var inputLine string
	_, err := fmt.Scanln(&inputLine)
	if err != nil {
		fmt.Println("Couldn't read input", err)
		return []int{}
	}

	result := make([]int, 0)
	rawNums := strings.Split(inputLine, ",")
	for _, rawNum := range rawNums {
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			fmt.Println("Couldn't read number from input", err)
		}
		result = append(result, num)
	}

	return result
}
