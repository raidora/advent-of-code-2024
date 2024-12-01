package four

import (
	"advent/utils/errors"
	"advent/utils/files"
	"testing"
)

func TestDayOneTaskOneDebug(t *testing.T) {
	file, err := files.ReadFileLines("testinput")
	errors.Check(err)
	taskOne(file)
}

func TestDayOneTaskOne(t *testing.T) {
	file, err := files.ReadFileLines("input")
	errors.Check(err)
	taskOne(file)
}

func TestDayOneTaskTwo(t *testing.T) {
	file, err := files.ReadFileLines("input")
	errors.Check(err)
	taskTwo(file)
}
