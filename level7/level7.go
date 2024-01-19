package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ruleSplitter = regexp.MustCompile("^([\\w ]+) bags contain (.+)\\.$")
var ruleMatcher = regexp.MustCompile("^(\\d+) ([\\w ]+) bags?$")

type bagRuleMapping struct {
	holdingBag string
	contents   map[string]int
}

func main() {
	rules := readBagRules()

	fmt.Println(countBagInEventuallyOthers(rules, "shiny gold"))
}

func countBagInEventuallyOthers(rules map[string]bagRuleMapping, bag string) int {
	count := 0

	for key := range rules {
		if canEventuallyContain(rules, key, bag) {
			count++
		}
	}

	return count
}

func canEventuallyContain(rules map[string]bagRuleMapping, current string, searched string) bool {
	toCheck := make([]bagRuleMapping, 0)
	toCheck = append(toCheck, rules[current])

	for len(toCheck) > 0 {
		checked := toCheck[len(toCheck)-1]
		toCheck = toCheck[:len(toCheck)-1]

		for key := range checked.contents {
			if key == searched {
				return true
			}

			toCheck = append(toCheck, rules[key])
		}
	}

	return false
}

func readBagRules() (rules map[string]bagRuleMapping) {
	reader := bufio.NewReader(os.Stdin)

	rules = make(map[string]bagRuleMapping)

	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n ")

		if err != nil || line == "" {
			return
		}

		ruleMapping, ok := processLine(line)

		if !ok {
			continue
		}

		rules[ruleMapping.holdingBag] = ruleMapping
	}
}

func processLine(line string) (ruleMapping bagRuleMapping, ok bool) {
	subMatches := ruleSplitter.FindStringSubmatch(line)

	if len(subMatches) < 3 {
		return
	}

	ruleMapping = bagRuleMapping{subMatches[1], make(map[string]int)}
	ok = true

	processMappings(ruleMapping.contents, subMatches[2])

	return
}

func processMappings(contents map[string]int, rawContents string) {
	for _, rawContent := range strings.Split(rawContents, ", ") {
		subMatches := ruleMatcher.FindStringSubmatch(rawContent)

		if len(subMatches) < 3 {
			continue
		}

		count, err := strconv.Atoi(subMatches[1])
		if err != nil {
			continue
		}

		bag := subMatches[2]

		contents[bag] = count
	}
}
