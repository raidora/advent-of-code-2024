package three

import (
	"advent/utils/errors"
	"advent/utils/files"
	"testing"
)

func TestDayOneTaskOne(t *testing.T) {
	file, err := files.ReadFile("input")
	errors.Check(err)
	taskOne(file)
}

func TestDayOneTaskTwo(t *testing.T) {
	file, err := files.ReadFile("input")
	errors.Check(err)
	taskTwo(file)
}
