package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// Конвертація префікс -> постфікс
func PrefixToPostfix(input string) (string, error) {
	var parts = strings.Split(input, " ")
	var arr []string

	for i := len(parts) - 1; i >= 0; i-- {
		if isStringOperator(parts[i]) {
			if len(arr) < 2 {
				return "", fmt.Errorf("невірний префіксний вираз")
			}

			var first = arr[len(arr)-1]
			var second = arr[len(arr)-2]
			arr = arr[:len(arr)-2]

			var res = first + " " + second + " " + parts[i]
			arr = append(arr, res)
		} else if isStringNumber(parts[i]) {
			arr = append(arr, parts[i])
		} else {
			return "", fmt.Errorf("невірний префіксний вираз")
		}
	}
	return strings.Join(arr, " "), nil
}

// Перевірка, чи є символ оператором
func isStringOperator(value string) bool {
	switch value {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

// Перевірка, чи є строка числом

func isStringNumber(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}
