package two

import (
	"advent/utils/errors"
	"advent/utils/files"
	"testing"
)

func TestDayTwoTaskOne(t *testing.T) {
	fileLines, err := files.ReadFileLines("input")
	errors.Check(err)
	taskOne(fileLines)
}

func TestDayTwoTaskTwo(t *testing.T) {
	fileLines, err := files.ReadFileLines("input")
	errors.Check(err)
	taskTwo(fileLines)
}
