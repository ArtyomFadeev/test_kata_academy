package calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Calculate parses and calculates the result of the given expression
func Calculate(expression string) (string, error) {
	// Remove spaces
	expression = strings.TrimSpace(expression)

	// Regular expression to match the expression pattern
	re := regexp.MustCompile(`^\"([^\"]{1,10})\"\s*([\+\-\*/])\s*(\"([^\"]{1,10})\"|([1-9]|10))$`)
	matches := re.FindStringSubmatch(expression)
	if matches == nil {
		return "", fmt.Errorf("invalid expression format")
	}

	str1 := matches[1]
	operator := matches[2]
	var str2 string
	var num int
	var err error

	if matches[4] != "" {
		str2 = matches[4]
	} else {
		num, err = strconv.Atoi(matches[5])
		if err != nil {
			return "", fmt.Errorf("invalid number")
		}
	}

	// Perform the operation
	switch operator {
	case "+":
		return str1 + str2, nil
	case "-":
		if strings.Contains(str1, str2) {
			return strings.ReplaceAll(str1, str2, ""), nil
		} else {
			return str1, nil
		}
	case "*":
		return strings.Repeat(str1, num), nil
	case "/":
		if num <= 0 || num > len(str1) {
			return "", fmt.Errorf("invalid division")
		}
		partLen := len(str1) / num
		return str1[:partLen], nil
	}

	return "", fmt.Errorf("unknown error")
}
