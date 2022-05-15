package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	if input == "" || len(strings.TrimSpace(input)) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	var sum int

	re := regexp.MustCompile("[0-9]+")

	for _, n := range re.FindAllString(input, -1) {
		num, err := strconv.Atoi(n)
		if err != nil {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
		sum += num

	}
	output = strconv.Itoa(sum)
	return output, nil
}

func main() {
	a := "13+55"
	fmt.Println(StringSum(a))
	fmt.Println(a)
}
