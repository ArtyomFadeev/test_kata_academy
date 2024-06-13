package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Римские цифры и их значения
var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []struct {
	Value   int
	Numeral string
}{
	{10, "X"}, {9, "IX"}, {8, "VIII"}, {7, "VII"}, {6, "VI"},
	{5, "V"}, {4, "IV"}, {3, "III"}, {2, "II"}, {1, "I"},
}

// Преобразование римских цифр в арабские
func romanToArabic(roman string) (int, error) {
	value, exists := romanNumerals[roman]
	if !exists {
		return 0, fmt.Errorf("неверное римское число: %s", roman)
	}
	return value, nil
}

// Преобразование арабских цифр в римские
func arabicToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 {
		return "", fmt.Errorf("в римской системе нет отрицательных чисел или нуля")
	}
	var result strings.Builder
	for _, numeral := range arabicToRoman {
		for arabic >= numeral.Value {
			arabic -= numeral.Value
			result.WriteString(numeral.Numeral)
		}
	}
	return result.String(), nil
}

// Разделение строки на компоненты (числа и операторы)
func parseInput(input string) (string, string, string, error) {
	re := regexp.MustCompile(`(\d+|I|II|III|IV|V|VI|VII|VIII|IX|X)([\+\-\*/])(\d+|I|II|III|IV|V|VI|VII|VIII|IX|X)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 4 {
		return "", "", "", fmt.Errorf("некорректный формат ввода")
	}
	return matches[1], matches[2], matches[3], nil
}

// Основная функция калькулятора
func calculate(input string) (string, error) {
	aStr, operator, bStr, err := parseInput(input)
	if err != nil {
		return "", err
	}

	var a, b int
	var isRoman bool

	if value, err := strconv.Atoi(aStr); err == nil {
		a = value
		isRoman = false
	} else {
		if value, err := romanToArabic(aStr); err == nil {
			a = value
			isRoman = true
		} else {
			return "", fmt.Errorf("некорректный формат числа: %s", aStr)
		}
	}

	if value, err := strconv.Atoi(bStr); err == nil {
		b = value
		if isRoman {
			return "", fmt.Errorf("используются одновременно разные системы счисления")
		}
	} else {
		if value, err := romanToArabic(bStr); err == nil {
			b = value
			if !isRoman {
				return "", fmt.Errorf("используются одновременно разные системы счисления")
			}
		} else {
			return "", fmt.Errorf("некорректный формат числа: %s", bStr)
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		return "", fmt.Errorf("числа должны быть в диапазоне от 1 до 10")
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", fmt.Errorf("деление на ноль")
		}
		result = a / b
	default:
		return "", fmt.Errorf("некорректный оператор: %s", operator)
	}

	if isRoman {
		if result < 1 {
			return "", fmt.Errorf("в римской системе нет отрицательных чисел или нуля")
		}
		return arabicToRomanNumeral(result)
	}
	return strconv.Itoa(result), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 2 + 3 или VI / III):")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	input = strings.TrimSpace(input)
	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}
