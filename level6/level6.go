package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	answers := readAnswers()

	fmt.Println(sumYesQuestionsByGroups(answers))
}

func sumYesQuestionsByGroups(answers []map[byte]int) int {
	sum := 0

	for _, group := range answers {
		sum += len(group)
	}

	return sum
}

func readAnswers() []map[byte]int {
	reader := bufio.NewReader(os.Stdin)

	answers := make([]map[byte]int, 0)
	current := make(map[byte]int)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if len(current) != 0 {
				answers = append(answers, current)
			}
			break
		}

		line = strings.TrimRight(line, "\n ")

		if line == "" {
			answers = append(answers, current)
			current = make(map[byte]int)
		}

		readPersonAnswers(current, line)
	}

	return answers
}

func readPersonAnswers(current map[byte]int, line string) {
	for _, char := range line {
		key := byte(char)

		if val, ok := current[key]; ok {
			current[key] = val + 1
		} else {
			current[key] = 1
		}
	}
}
