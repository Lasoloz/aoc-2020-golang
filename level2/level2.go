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

	count := countValidPasswords(policies)
	fmt.Println(count)
}

func countValidPasswords(policies []pwdPolicy) (correct int) {
	for _, p := range policies {
		charCount := 0

		for _, ch := range p.pwd {
			if ch == p.char {
				charCount++
			}
		}

		if p.min <= charCount && charCount <= p.max {
			correct++
		}
	}

	return
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
