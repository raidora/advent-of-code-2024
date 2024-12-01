package one

import (
	"advent/utils/errors"
	"advent/utils/files"
	"testing"
)

func TestDayOneTaskOne(t *testing.T) {
	fileLines, err := files.ReadFileLines("input")
	errors.Check(err)
	taskOne(fileLines)
}

func TestDayOneTaskTwo(t *testing.T) {
	fileLines, err := files.ReadFileLines("input")
	errors.Check(err)
	taskTwo(fileLines)
}
