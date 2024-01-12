package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type validator struct {
	field    string
	validate func(value string) bool
}

var requiredPassportFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

var colorRegex, _ = regexp.Compile("^#[0-9a-f]{6}$")
var eyeColorRegex, _ = regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
var pidRegex, _ = regexp.Compile("^[0-9]{9}$")
var passportFieldValidators = []validator{
	{"byr", createDateValidator(1920, 2002)},
	{"iyr", createDateValidator(2010, 2020)},
	{"eyr", createDateValidator(2020, 2030)},
	{"hgt", heightValidator},
	{"hcl", createRegexValidator(colorRegex)},
	{"ecl", createRegexValidator(eyeColorRegex)},
	{"pid", createRegexValidator(pidRegex)},
}

func main() {
	passports := readPassports()

	fmt.Println(countValidPassports(passports, validatePassportFieldsExist))
	fmt.Println(countValidPassports(passports, validatePassportFields))
}

func countValidPassports(passports []map[string]string, validator func(map[string]string) bool) int {
	validPassports := 0

	for _, passport := range passports {
		if validator(passport) {
			validPassports++
		}
	}

	return validPassports
}

func validatePassportFieldsExist(passport map[string]string) bool {
	fields := 0

	for _, field := range requiredPassportFields {
		if _, ok := passport[field]; ok {
			fields++
		}
	}

	return fields == len(requiredPassportFields)
}

func validatePassportFields(passport map[string]string) bool {
	fields := 0

	for _, v := range passportFieldValidators {
		if value, ok := passport[v.field]; ok && v.validate(value) {
			fields++
		}
	}

	return fields == len(passportFieldValidators)
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
		passport[key] = value
	}
}

func createDateValidator(min, max int) func(string) bool {
	return func(value string) bool {
		if year, err := strconv.Atoi(value); err == nil {
			return year >= min && year <= max
		}

		return false
	}
}

func createRegexValidator(regex *regexp.Regexp) func(string) bool {
	return func(value string) bool { return regex.MatchString(value) }
}

func heightValidator(value string) bool {
	var height int
	if _, err := fmt.Sscanf(value, "%dcm", &height); err == nil {
		return height >= 150 && height <= 193
	}
	if _, err := fmt.Sscanf(value, "%din", &height); err == nil {
		return height >= 59 && height <= 76
	}

	return false
}
