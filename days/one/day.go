package one

import (
	"advent/utils/errors"
	"advent/utils/pretty"
	"math"
	"slices"
	"strconv"
	"strings"
)

func parseColumns(lines []string) (leftCol []int, rightCol []int, err error) {
	for _, line := range lines {
		cols := strings.Split(line, "  ")

		leftNum, err := strconv.Atoi(strings.TrimRight(cols[0], " "))
		if err != nil {
			return leftCol, rightCol, err
		}
		rightNum, err := strconv.Atoi(strings.TrimLeft(cols[1], " "))
		if err != nil {
			return leftCol, rightCol, err
		}

		leftCol = append(leftCol, leftNum)
		rightCol = append(rightCol, rightNum)
	}

	return leftCol, rightCol, err
}

func taskOne(lines []string) {
	var sumDistance int

	leftCol, rightCol, err := parseColumns(lines)
	errors.Check(err)

	slices.Sort(leftCol)
	slices.Sort(rightCol)

	if len(leftCol) != len(rightCol) {
		panic("Column lengths don't match, can't calculate distances")
	}

	for idx := range len(leftCol) {
		sumDistance += int(math.Abs(float64(leftCol[idx] - rightCol[idx])))
	}

	pretty.PrintResult(sumDistance, 1, 1)
}

func taskTwo(lines []string) {
	var similarityScore int

	leftCol, rightCol, err := parseColumns(lines)
	errors.Check(err)

	for _, leftID := range leftCol {
		var matchTimes int
		for _, rightID := range rightCol {
			if leftID == rightID {
				matchTimes++
			}
		}
		similarityScore += leftID * matchTimes
	}

	pretty.PrintResult(similarityScore, 1, 2)
}
