package files

import (
	"bufio"
	"os"
)

func ReadFileLines(path string) (lines []string, err error) {
	var fileLines []string

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	return fileLines, nil
}
