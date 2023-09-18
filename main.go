package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicNumerals = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func isRomanNumeral(s string) bool {
	_, ok := romanNumerals[s]
	return ok
}

func parseRomanNumeral(s string) int {
	return romanNumerals[s]
}

func toRomanNumeral(n int) string {
	if n <= 0 {
		return "Error"
	}
	roman := ""
	for _, numeral := range []int{10, 9, 5, 4, 1} {
		for n >= numeral {
			roman += arabicNumerals[numeral]
			n -= numeral
		}
	}
	return roman
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Simple Calculator")
	fmt.Println("Available operators: +, -, *, /")
	fmt.Println("Enter 'exit' to quit")

	for {
		fmt.Print("Enter an expression: ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Error: Invalid input format")
			continue
		}

		num1, num2 := 0, 0

		if isRomanNumeral(parts[0]) && isRomanNumeral(parts[2]) {
			num1 = parseRomanNumeral(parts[0])
			num2 = parseRomanNumeral(parts[2])
		} else {
			fmt.Println("Error: Invalid number")
			continue
		}

		operator := parts[1]
		result := 0

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 != 0 {
				result = num1 / num2
			} else {
				fmt.Println("Error: Division by zero")
				continue
			}
		default:
			fmt.Println("Error: Invalid operator")
			continue
		}

		if isRomanNumeral(parts[0]) && isRomanNumeral(parts[2]) {
			if result < 1 {
				fmt.Println("Error: Result cannot be represented as a Roman numeral")
			} else {
				fmt.Printf("Result: %s\n", toRomanNumeral(result))
			}
		} else {
			fmt.Printf("Result: %d\n", result)
		}
	}
}
