package string_sum

// package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errSomeNewError     = errors.New("something go wrong")
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
	if input == "" || len(strings.TrimSpace(input)) == 0 {
		return "", fmt.Errorf("empty input, %w", errorEmptyInput)
	}

	var sum, count int

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	for _, n := range re.FindAllString(input, -1) {
		num, err := strconv.Atoi(n)
		if err != nil {
			return "", fmt.Errorf("some unexpecting error, %w", errSomeNewError)
		}
		count++
		sum += num
	}

	if count != 2 {
		return "", fmt.Errorf("an incredible number of operands, %w", errorNotTwoOperands)
	}

	output = strconv.Itoa(sum)
	return output, nil
}

// func main() {
// 	a := "-55+-33"
// 	fmt.Println(StringSum(a))
// 	fmt.Println(a)
// }
