package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// Префікс -> постфікс
func PrefixToPostfix(input string) (string, error) {
	parts := strings.Fields(input)
	var arr []string

	for i := len(parts) - 1; i >= 0; i-- {
		switch {
		case isStringOperator(parts[i]):
			if len(arr) < 2 {
				return "", fmt.Errorf("Невірний префіксний вираз")
			}
			first := arr[len(arr)-1]
			second := arr[len(arr)-2]
			arr = arr[:len(arr)-2]

			result := first + " " + second + " " + parts[i]
			arr = append(arr, result)
		case isStringNumber(parts[i]):
			arr = append(arr, parts[i])
		default:
			return "", fmt.Errorf("Невірний префіксний вираз")
		}
	}
	return strings.Join(arr, " "), nil
}

// Перевірка, чи є строка числом
func isStringNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Перевірка, чи є символ оператором
func isStringOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

