package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"test_kata_academy/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your expression:")
	expression, err := reader.ReadString('\n')
	if err != nil {
		panic("Failed to read input")
	}
	expression = strings.TrimSpace(expression)

	result, err := calculator.Calculate(expression)
	if err != nil {
		panic(err)
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}
	fmt.Println(result)
}
