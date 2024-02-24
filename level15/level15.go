package main

import (
	"fmt"
	"strconv"
	"strings"
)

const searchLen = 2020

func main() {
	startingNumbers := readStartingNumbers()

	if len(startingNumbers) == 0 {
		fmt.Println("Empty starting numbers list")
		return
	}

	fmt.Println(find2020thNumber(startingNumbers))
}

func find2020thNumber(startingNumbers []int) int {
	numbers := make([]int, searchLen)
	copy(numbers, startingNumbers)

	for i := len(startingNumbers); i < searchLen; i++ {
		last := numbers[i-1]
		diff := 0

		for j := i - 2; j >= 0; j-- {
			if numbers[j] == last {
				diff = i - j - 1
				break
			}
		}

		numbers[i] = diff
	}

	return numbers[searchLen-1]
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
