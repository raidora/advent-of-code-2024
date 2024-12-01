package main

import (
	"advent/utils/errors"
	"advent/utils/files"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var leftCol []int
	var rightCol []int
	var sumDistance int64

	lines, err := files.ReadFileLines("days/one/input")
	errors.Check(err)

	for _, line := range lines {
		cols := strings.Split(line, "  ")

		leftNum, err := strconv.Atoi(strings.TrimRight(cols[0], " "))
		errors.Check(err)
		rightNum, err := strconv.Atoi(strings.TrimLeft(cols[1], " "))
		errors.Check(err)

		leftCol = append(leftCol, leftNum)
		rightCol = append(rightCol, rightNum)
	}

	slices.Sort(leftCol)
	slices.Sort(rightCol)

	if len(leftCol) != len(rightCol) {
		panic("Column lengths don't match, can't calculate distances")
	}

	for idx := range len(leftCol) {
		sumDistance += int64(math.Abs(float64(leftCol[idx] - rightCol[idx])))
	}

	fmt.Println(sumDistance)
}
