package four

import (
	"fmt"
	"slices"
)

type HorizontalDirection int
type VerticalDirection int
type DiagonalDirection int

const (
	DirectionLeft HorizontalDirection = iota
	DirectionRight
)

const (
	DirectionUp VerticalDirection = iota
	DirectionDown
)

func convertInput(lines []string) (chars [][]byte) {
	for _, line := range lines {
		chars = append(chars, []byte(line))
	}
	return chars
}

func copy2DSlice(b [][]byte) [][]byte {
	tmp := make([][]byte, len(b))
	copy(tmp, b)
	return tmp
}

func copySlice(b []byte) []byte {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}

func flipVertical(b [][]byte) [][]byte {
	tmp := copy2DSlice(b)
	slices.Reverse(tmp)
	return tmp
}

func readColumn(b [][]byte, colIdx int) []byte {
	var column []byte
	height := len(b)

	for idx := range height - 1 {
		column = append(column, b[idx][colIdx])
	}
	return column
}

func findHorizontalWordOccurrences(b [][]byte, word string) int {
	var occurrences int

	for _, bRow := range b {
		for idx := range bRow {
			var scannable []byte
			wLen := len(word)
			if idx > len(bRow)-wLen {
				break
			} else {
				scannable = bRow[idx : idx+wLen]
			}
			scnFwd := string(scannable)
			tmp := copySlice(scannable)
			slices.Reverse(tmp)
			scnBwd := string(tmp)

			if scnFwd == word || scnBwd == word {
				occurrences++
			}
		}
	}
	return occurrences
}

func findVerticalWordOccurrences(b [][]byte, word string) int {
	var occurrences int
	width := len(b[0])

	for colIdx := range width - 1 {
		column := readColumn(b, colIdx)

		scnFwd := string(column)
		tmp := copySlice(column)
		slices.Reverse(tmp)
		scnBwd := string(tmp)
		if scnFwd == word || scnBwd == word {
			occurrences++
		}
	}
	return occurrences
}

func findDiagonalWordOccurrences(b [][]byte, word string) int {
	var occurrences int

	for rowIdx, row := range b {
		for colIdx := range row {
			var wdLR []byte
			var wdRL []byte
			if colIdx > len(row)-len(word) {
				break
			}

			//left to right diagonal
			wdLR = append(wdLR, b[rowIdx][colIdx])
			wdLR = append(wdLR, b[rowIdx][colIdx+1])
			wdLR = append(wdLR, b[rowIdx][colIdx+2])
			wdLR = append(wdLR, b[rowIdx][colIdx+3])

			//left to right diagonal
			wdRL = append(wdRL, b[rowIdx][colIdx+3])
			wdRL = append(wdRL, b[rowIdx][colIdx+2])
			wdRL = append(wdRL, b[rowIdx][colIdx+1])
			wdRL = append(wdRL, b[rowIdx][colIdx])

			scnFwdLR := string(wdLR)
			tmpLR := copySlice(wdLR)
			slices.Reverse(tmpLR)
			scnBwdLR := string(tmpLR)

			if scnFwdLR == word || scnBwdLR == word {
				occurrences++
			}

			scnFwdRL := string(wdRL)
			tmpRL := copySlice(wdRL)
			slices.Reverse(tmpRL)
			scnBwdRL := string(tmpRL)

			if scnFwdRL == word || scnBwdRL == word {
				occurrences++
			}
		}
	}
	return occurrences
}

func taskOne(inputLines []string) {
	sum := 0
	chars := convertInput(inputLines)
	sum += findHorizontalWordOccurrences(chars, "XMAS")
	sum += findVerticalWordOccurrences(chars, "XMAS")
	sum += findDiagonalWordOccurrences(chars, "XMAS")

	fmt.Println(chars)
}

func taskTwo(inputLines []string) {
}
