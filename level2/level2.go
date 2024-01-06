package main

import (
	"fmt"
)

type pwdPolicy struct {
	min  int
	max  int
	char int32
	pwd  string
}

func main() {
	policies := readPolicies()

	fmt.Println(countValidPasswords(policies, part1Validator))
	fmt.Println(countValidPasswords(policies, part2Validator))
}

func countValidPasswords(policies []pwdPolicy, pwdValidator func(policy pwdPolicy) bool) (correct int) {
	for _, p := range policies {
		if pwdValidator(p) {
			correct++
		}
	}

	return
}

func part1Validator(p pwdPolicy) bool {
	charCount := 0

	for _, ch := range p.pwd {
		if ch == p.char {
			charCount++
		}
	}

	return p.min <= charCount && charCount <= p.max
}

func part2Validator(p pwdPolicy) bool {
	positionCount := 0

	for index, ch := range p.pwd {
		pos := index + 1
		if pos != p.min && pos != p.max {
			continue
		}

		if ch == p.char {
			positionCount++
		}
	}

	return positionCount == 1
}

func readPolicies() []pwdPolicy {
	result := make([]pwdPolicy, 0)

	for {
		policy := pwdPolicy{}
		_, ok := fmt.Scanf("%d-%d %c: %s\n", &policy.min, &policy.max, &policy.char, &policy.pwd)

		if ok != nil {
			break
		}

		result = append(result, policy)
	}

	return result
}
