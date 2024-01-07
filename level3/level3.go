package main

import "fmt"

const stepX = 3
const stepY = 1
const tree = '#'

type slopeSection struct {
	lines  []string
	width  int
	height int
}

func main() {
	slope := readSlope()

	fmt.Println(countTrees(slope))
}

func countTrees(slope slopeSection) int {
	width, height := slope.width, slope.height
	x, y := stepX%width, stepY

	trees := 0

	for y < height {
		if slope.lines[y][x] == tree {
			trees++
		}

		x = (x + stepX) % width
		y = y + stepY
	}

	return trees
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
