package string_sum

// package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorHasLetter      = errors.New("got letter")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "+", " ")
	input = strings.ReplaceAll(input, "-", " -")

	if input == "" || len(strings.TrimSpace(input)) == 0 {
		return "", fmt.Errorf("empty input, %w", errorEmptyInput)
	}

	for _, v := range input {
		if !strings.Contains("0123456789+- ", string(v)) {
			return "", fmt.Errorf("maybe it contains letter or symbol: %w", errorHasLetter)
		}
	}

	operands := strings.Fields(input)
	if len(operands) != 2 {
		err = fmt.Errorf("error: %w", errorNotTwoOperands)
		return
	}

	firstOperand, firstErr := strconv.Atoi(operands[0])
	if firstErr != nil {
		err = fmt.Errorf("error: %w", firstErr)
		return
	}

	secondOperand, SecondErr := strconv.Atoi(operands[1])
	if SecondErr != nil {
		err = fmt.Errorf("error: %w", SecondErr)
		return
	}

	output = strconv.Itoa(firstOperand + secondOperand)
	err = nil
	return
}

// func main() {
// 	a := "-f55+-33"
// 	fmt.Println(StringSum(a))
// 	fmt.Println(a)
// }
