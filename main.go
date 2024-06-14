package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Римские цифры и их значения
var romanNumerals = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
}

var arabicToRoman = []struct {
	Value   int
	Numeral string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// Преобразование римских цифр в арабские
func romanToArabic(roman string) (int, error) {
	validRoman := regexp.MustCompile(`^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	if !validRoman.MatchString(roman) {
		return 0, errors.New("неверное римское число: " + roman)
	}

	result := 0
	i := 0
	for i < len(roman) {
		if i+1 < len(roman) {
			twoChar := roman[i : i+2]
			if value, exists := romanNumerals[twoChar]; exists {
				result += value
				i += 2
				continue
			}
		}
		oneChar := roman[i : i+1]
		if value, exists := romanNumerals[oneChar]; exists {
			result += value
			i++
		} else {
			return 0, errors.New("неверное римское число: " + roman)
		}
	}
	return result, nil
}

// Преобразование арабских цифр в римские
func arabicToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 {
		return "", errors.New("в римской системе нет отрицательных чисел или нуля")
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

// Основная функция калькулятора
func calculate(input string) (string, error) {
	// Регулярное выражение для разбора строки
	re := regexp.MustCompile(`^(\d+|[IVXLCDM]+)\s*([+\-*/])\s*(\d+|[IVXLCDM]+)$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		return "", errors.New("некорректный формат ввода")
	}

	aStr, operator, bStr := matches[1], matches[2], matches[3]

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
			return "", errors.New("некорректный формат числа: " + aStr)
		}
	}

	if value, err := strconv.Atoi(bStr); err == nil {
		b = value
		if isRoman {
			return "", errors.New("используются одновременно разные системы счисления")
		}
	} else {
		if value, err := romanToArabic(bStr); err == nil {
			b = value
			if !isRoman {
				return "", errors.New("используются одновременно разные системы счисления")
			}
		} else {
			return "", errors.New("некорректный формат числа: " + bStr)
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		return "", errors.New("числа должны быть в диапазоне от 1 до 10")
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
			return "", errors.New("в римской системе нет отрицательных чисел или нуля")
		}
		return arabicToRomanNumeral(result)
	}
	return strconv.Itoa(result), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 2+3 или VI/III):")

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
