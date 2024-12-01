package three

import (
	"advent/utils/errors"
	"advent/utils/pretty"
	"regexp"
	"strconv"
	"strings"
)

func taskOne(file []byte) {
	sum := 0
	r, err := regexp.Compile("mul\\(\\d+,\\d+\\)")
	errors.Check(err)

	matches := r.FindAll(file, -1)

	for _, match := range matches {
		product, err := executeMultiplyToken(string(match))
		errors.Check(err)
		sum += product
	}

	pretty.PrintResult(sum, 3, 1)
}

func tokenize(input []byte) (validTokens []string, err error) {
	enableRead := true
	doExp, err := regexp.Compile("^do\\(\\)")
	if err != nil {
		return nil, err
	}
	dontExp, err := regexp.Compile("^don't\\(\\)")
	if err != nil {
		return nil, err
	}
	mulExp, err := regexp.Compile("^mul\\(\\d+,\\d+\\)")
	if err != nil {
		return nil, err
	}

	for i := range input {
		match := doExp.Match(input[i:])
		if match {
			enableRead = true
			continue
		}

		match = dontExp.Match(input[i:])
		if match {
			enableRead = false
			continue
		}

		find := mulExp.Find(input[i:])
		if len(find) > 0 && enableRead {
			validTokens = append(validTokens, string(find))
			continue
		}
	}

	return validTokens, nil
}

func executeMultiplyToken(token string) (int, error) {
	cleanToken := token
	cleanToken = strings.TrimLeft(cleanToken, "mul(")
	cleanToken = strings.TrimRight(cleanToken, ")")

	numbers := strings.Split(cleanToken, ",")
	numOne, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, err
	}
	numTwo, err := strconv.Atoi(numbers[1])
	if err != nil {
		return 0, err
	}

	return numOne * numTwo, nil
}

func taskTwo(file []byte) {
	sum := 0

	tokens, err := tokenize(file)
	errors.Check(err)

	for _, token := range tokens {
		product, err := executeMultiplyToken(token)
		errors.Check(err)
		sum += product
	}

	pretty.PrintResult(sum, 3, 2)
}
