package main

import (
	"fmt"
	"os"
	"strconv"
)

const defaultPreamble = 25

func main() {
	preambleSize := readPreambleSize()
	portData := readPortData()

	invalidNumber := findFirstInvalidNumber(portData, preambleSize)
	fmt.Println(invalidNumber)
	fmt.Println(findXMASWeakness(portData, invalidNumber))
}

func findFirstInvalidNumber(portData []int, preambleSize int) int {
	index := preambleSize

	for index < len(portData) {
		current := portData[index]

		if !isSummableByPairInSlice(portData[index-preambleSize:index], current) {
			return current
		}

		index++
	}

	return -1
}

func isSummableByPairInSlice(dataSlice []int, current int) bool {
	for i, first := range dataSlice[:len(dataSlice)-1] {
		for _, second := range dataSlice[i+1:] {
			if first+second == current {
				return true
			}
		}
	}

	return false
}

func findXMASWeakness(portData []int, invalidNumber int) int {
	for i, first := range portData {
		sum := first
		minValue, maxValue := first, first

		for _, current := range portData[i+1:] {
			sum += current

			if sum == invalidNumber {
				return minValue + maxValue
			} else if sum > invalidNumber {
				break
			}

			if current < minValue {
				minValue = current
			} else if current > maxValue {
				maxValue = current
			}
		}
	}

	return -1
}

func readPortData() []int {
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

func readPreambleSize() int {
	if len(os.Args) < 2 {
		return defaultPreamble
	}

	if value, err := strconv.Atoi(os.Args[1]); err != nil {
		return defaultPreamble
	} else {
		return value
	}
}
