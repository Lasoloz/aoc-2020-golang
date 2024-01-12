package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type groupDeclarations struct {
	peopleCount int
	answers     map[byte]int
}

func main() {
	answers := readAnswers()

	fmt.Println(sumYesQuestionsByAnyoneByGroups(answers))
	fmt.Println(sumYesQuestionsByEveryoneByGroups(answers))
}

func sumYesQuestionsByAnyoneByGroups(declarations []groupDeclarations) int {
	sum := 0

	for _, group := range declarations {
		sum += len(group.answers)
	}

	return sum
}

func sumYesQuestionsByEveryoneByGroups(declarations []groupDeclarations) int {
	sum := 0

	for _, group := range declarations {
		for _, answerCount := range group.answers {
			if answerCount == group.peopleCount {
				sum++
			}
		}
	}

	return sum
}

func readAnswers() []groupDeclarations {
	reader := bufio.NewReader(os.Stdin)

	answers := make([]groupDeclarations, 0)
	current := groupDeclarations{0, make(map[byte]int)}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if len(current.answers) != 0 {
				answers = append(answers, current)
			}
			break
		}

		line = strings.TrimRight(line, "\n ")

		if line == "" {
			answers = append(answers, current)
			current = groupDeclarations{0, make(map[byte]int)}
			continue
		}

		readPersonAnswers(current.answers, line)
		current.peopleCount++
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
