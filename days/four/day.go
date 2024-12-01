package four

import (
	"fmt"
	"math"
	"slices"
)

func convertInput(lines []string) (chars [][]byte) {
	for _, line := range lines {
		chars = append(chars, []byte(line))
	}
	return chars
}

func matchBothWays(arr []byte, word string) int {
	var matches int

	scnFwd := string(arr)
	slices.Reverse(arr)
	scnBwd := string(arr)

	if scnFwd == word {
		matches++
	}
	if scnBwd == word {
		matches++
	}

	return matches
}

func findHorizontalWordOccurrences(b [][]byte, word string) int {
	var occurrences int

	for rID := range len(b) {
		for cID := range len(b[rID]) - len(word) + 1 {
			var rowWd []byte
			for letterID := range len(word) {
				rowWd = append(rowWd, b[rID][cID+letterID])
			}

			occurrences += matchBothWays(rowWd, word)
		}
	}

	return occurrences
}

func findVerticalWordOccurrences(b [][]byte, word string) int {
	var occurrences int

	for rID := range len(b) - len(word) + 1 {
		for cID := range len(b[rID]) {
			var colWd []byte
			for letterID := range len(word) {
				colWd = append(colWd, b[rID+letterID][cID])
			}

			occurrences += matchBothWays(colWd, word)
		}
	}
	return occurrences
}

func findDiagonalWordOccurrences(b [][]byte, word string) int {
	var occurrences int

	for rID := range len(b) - len(word) + 1 {
		for cID := range len(b[rID]) - len(word) + 1 {
			// Scan left to right
			var wdLR []byte
			for lrID := range len(word) {
				wdLR = append(wdLR, b[rID+lrID][cID+lrID])
			}
			occurrences += matchBothWays(wdLR, word)

			// Scan right to left
			var wdRL []byte
			for rlID := range len(word) {
				rlIDReverseCounter := int(math.Abs(float64(rlID - len(word) + 1)))
				wdRL = append(wdRL, b[rID+rlID][cID+rlIDReverseCounter])
			}
			occurrences += matchBothWays(wdLR, word)
		}
	}
	return occurrences
}

func taskOne(inputLines []string) {
	sum := 0
	chars := convertInput(inputLines)
	sum += findHorizontalWordOccurrences(chars, "XMAS")
	fmt.Println(sum)
	sum += findVerticalWordOccurrences(chars, "XMAS")
	fmt.Println(sum)
	sum += findDiagonalWordOccurrences(chars, "XMAS")
	fmt.Println(sum)
}

func taskTwo(inputLines []string) {
}
