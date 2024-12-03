package two

import (
	"advent/utils/errors"
	"advent/utils/pretty"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type input struct {
	reports []report
}

type report struct {
	levels []int
}

func isIncreased(a int, b int) bool {
	return b > a
}

func isDecreased(a int, b int) bool {
	return b < a
}

func passesDistanceConstraint(a *int, b *int) bool {
	distance := int(math.Abs(float64(*a - *b)))
	return distance > 0 && distance <= 3
}

func (r *report) isSafe() bool {
	var previousLevel int
	var expectIncreasingLevels bool

	for idx, level := range r.levels {
		if idx == 0 {
			previousLevel = level
			continue
		}

		if idx == 1 && previousLevel < level {
			expectIncreasingLevels = true
		}

		if !passesDistanceConstraint(&previousLevel, &level) {
			return false
		}

		if expectIncreasingLevels {
			if isDecreased(previousLevel, level) {
				return false
			}
		} else {
			if isIncreased(previousLevel, level) {
				return false
			}
		}

		previousLevel = level
	}

	return true
}

func convertInput(inputLines []string) (*input, error) {
	inpObj := input{}

	for _, line := range inputLines {
		var levels []int

		chars := strings.Split(line, " ")
		for _, char := range chars {
			level, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}

			levels = append(levels, level)
		}

		inpObj.reports = append(inpObj.reports, report{
			levels: levels,
		})
	}
	return &inpObj, nil
}

func taskOne(inputList []string) {
	var safeReports int
	var totalReports int
	inpObj, err := convertInput(inputList)
	errors.Check(err)

	for _, report := range inpObj.reports {
		if report.isSafe() {
			safeReports++
		}
		totalReports++
	}

	fmt.Printf("%d reports of %d are safe\r\n", safeReports, totalReports)
	pretty.PrintResult(safeReports, 2, 1)
}

func removeItem(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func getDampenedReportVariants(rep report) []report {
	var variants []report

	for i := range len(rep.levels) {
		tmp := make([]int, len(rep.levels))
		copy(tmp, rep.levels)
		variants = append(variants, report{
			levels: removeItem(tmp, i),
		})
	}

	return variants
}

func taskTwo(inputList []string) {
	var safeReports int
	var totalReports int
	inpObj, err := convertInput(inputList)
	errors.Check(err)

	for _, report := range inpObj.reports {
		if report.isSafe() {
			safeReports++
		} else {
			for _, report := range getDampenedReportVariants(report) {
				if report.isSafe() {
					safeReports++
					break
				}
			}
		}

		totalReports++
	}

	fmt.Printf("%d reports of %d are safe\r\n", safeReports, totalReports)
	pretty.PrintResult(safeReports, 2, 2)
}
