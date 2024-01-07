package main

import "fmt"

const stepX = 3
const stepY = 1
const tree = '#'

type step struct {
	x int
	y int
}

var steps = []step{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

type slopeSection struct {
	lines  []string
	width  int
	height int
}

func main() {
	slope := readSlope()

	fmt.Println(countTrees(slope, stepX, stepY))
	fmt.Println(countTreeConfigurations(slope))
}

func countTrees(slope slopeSection, dx, dy int) int {
	width, height := slope.width, slope.height
	x, y := dx%width, dy

	trees := 0

	for y < height {
		if slope.lines[y][x] == tree {
			trees++
		}

		x = (x + dx) % width
		y = y + dy
	}

	return trees
}

func countTreeConfigurations(slope slopeSection) int {
	multipliedHits := 1

	for _, step := range steps {
		multipliedHits *= countTrees(slope, step.x, step.y)
	}

	return multipliedHits
}

func readSlope() slopeSection {
	result := make([]string, 0)

	for {
		var line string
		_, ok := fmt.Scanln(&line)

		if ok != nil {
			break
		}

		result = append(result, line)
	}

	if len(result) == 0 {
		return slopeSection{}
	}

	return slopeSection{result, len(result[0]), len(result)}
}
