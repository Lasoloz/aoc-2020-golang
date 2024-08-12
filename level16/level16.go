package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type intervalDef struct {
	start int
	end   int
}

type ticketRangeDef struct {
	name      string
	intervals []intervalDef
}

type ticketRangeStat struct {
	rangeDefs      []ticketRangeDef
	personalTicket []int
	nearbyTickets  [][]int
}

var ticketTypeRegex = regexp.MustCompile("(\\w+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")

func main() {
	ticketStat := readTicketRanges()
	if ticketStat == nil {
		fmt.Println("Erroneous input!")
		os.Exit(1)
	}

	fmt.Println(calculateTicketScanningErrorRate(*ticketStat))
}

func calculateTicketScanningErrorRate(stat ticketRangeStat) int {
	errorRate := 0

	for _, ticket := range stat.nearbyTickets {
		for _, num := range ticket {
			found := false
			for _, rd := range stat.rangeDefs {
				for _, id := range rd.intervals {
					if num >= id.start && num <= id.end {
						found = true
						break
					}
				}
				if found {
					break
				}
			}

			if !found {
				errorRate += num
			}
		}
	}

	return errorRate
}

func readTicketRanges() *ticketRangeStat {
	reader := bufio.NewReader(os.Stdin)
	ticketStat := &ticketRangeStat{make([]ticketRangeDef, 0), make([]int, 0), make([][]int, 0)}

	mode := 0

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimRight(line, "\n")
		if err != nil {
			break
		}

		if len(line) == 0 {
			continue
		}

		if line == "your ticket:" {
			mode = 1
			continue
		}

		if line == "nearby tickets:" {
			mode = 2
			continue
		}

		switch mode {
		case 0:
			rangeDef := readRangeDef(line)
			if rangeDef == nil {
				return nil
			}
			ticketStat.rangeDefs = append(ticketStat.rangeDefs, *rangeDef)
		case 1:
			ticket := readTicket(line)
			if ticket == nil {
				return nil
			}
			ticketStat.personalTicket = ticket
		case 2:
			ticket := readTicket(line)
			if ticket == nil {
				return nil
			}
			ticketStat.nearbyTickets = append(ticketStat.nearbyTickets, ticket)
		default:
			break
		}
	}

	return ticketStat
}

func readRangeDef(line string) *ticketRangeDef {
	matches := ticketTypeRegex.FindStringSubmatch(line)
	if matches == nil || len(matches) != 6 {
		return nil
	}

	s1, _ := strconv.Atoi(matches[2])
	e1, _ := strconv.Atoi(matches[3])
	s2, _ := strconv.Atoi(matches[4])
	e2, _ := strconv.Atoi(matches[5])

	return &ticketRangeDef{name: matches[1], intervals: []intervalDef{
		{s1, e1},
		{s2, e2},
	}}
}

func readTicket(rawTicket string) []int {
	rawNums := strings.Split(rawTicket, ",")
	result := make([]int, 0)

	for _, rawNum := range rawNums {
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			return nil
		}
		result = append(result, num)
	}

	return result
}
