package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var requiredPassportFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	passports := readPassports()

	fmt.Println(countValidPassports(passports))
}

func countValidPassports(passports []map[string]string) int {
	validPassports := 0

	for _, passport := range passports {
		if validatePassport(passport) {
			validPassports++
		}
	}

	return validPassports
}

func validatePassport(passport map[string]string) bool {
	fields := 0

	for _, field := range requiredPassportFields {
		if _, ok := passport[field]; ok {
			fields++
		}
	}

	return fields == len(requiredPassportFields)
}

func readPassports() []map[string]string {
	reader := bufio.NewReader(os.Stdin)

	passports := make([]map[string]string, 0)
	current := make(map[string]string)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if len(current) != 0 {
				passports = append(passports, current)
			}
			break
		}

		line = strings.TrimRight(line, "\n ")

		if line == "" {
			passports = append(passports, current)
			current = make(map[string]string)
		}

		readPassportFields(current, line)
	}

	return passports
}

func readPassportFields(passport map[string]string, line string) {
	for _, entry := range strings.Split(line, " ") {
		keyValue := strings.Split(entry, ":")

		if len(keyValue) != 2 {
			continue
		}

		key, value := keyValue[0], keyValue[1]
		(passport)[key] = value
	}
}
