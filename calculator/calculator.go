package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

// Calculate parses and calculates the result of the given expression
func Calculate(expression string) (string, error) {
	// Remove spaces
	expression = strings.ReplaceAll(expression, " ", "")

	// Parse the expression
	var str1, str2 string
	var num int
	var operator string

	// Detect operation type
	if strings.Contains(expression, "+") {
		operator = "+"
	} else if strings.Contains(expression, "-") {
		operator = "-"
	} else if strings.Contains(expression, "*") {
		operator = "*"
	} else if strings.Contains(expression, "/") {
		operator = "/"
	} else {
		return "", fmt.Errorf("unsupported operation")
	}

	// Split the expression
	parts := strings.Split(expression, operator)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid expression format")
	}

	// Trim quotes and validate
	str1 = strings.Trim(parts[0], "\"")
	if len(str1) > 10 {
		return "", fmt.Errorf("first string is too long")
	}

	if operator == "+" || operator == "-" {
		str2 = strings.Trim(parts[1], "\"")
		if len(str2) > 10 {
			return "", fmt.Errorf("second string is too long")
		}
	} else {
		var err error
		num, err = strconv.Atoi(parts[1])
		if err != nil || num < 1 || num > 10 {
			return "", fmt.Errorf("invalid number")
		}
	}

	// Perform the operation
	switch operator {
	case "+":
		return str1 + str2, nil
	case "-":
		return strings.ReplaceAll(str1, str2, ""), nil
	case "*":
		return strings.Repeat(str1, num), nil
	case "/":
		partLen := len(str1) / num
		return str1[:partLen], nil
	}

	return "", fmt.Errorf("unknown error")
}
