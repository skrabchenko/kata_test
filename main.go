package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumStrings(a, b string) string {
	return a + b
}

func subtractStrings(a, b string) string {
	return strings.Replace(a, b, "", 1)
}

func multiplyStringByNumber(a string, b int) string {
	result := ""
	for i := 0; i < b; i++ {
		result += a
	}
	return result
}

func divideStringByNumber(a string, b int) string {
	if b == 0 {
		return ""
	}
	partLen := len(a) / b
	return a[:partLen]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Убираем пробелы вокруг строки
	input = strings.TrimSpace(input)

	// Проверяем, что строка начинается и заканчивается на кавычки
	if !(strings.HasPrefix(input, "\"") && strings.Contains(input, "\" ")) {
		fmt.Println("Неверный формат ввода")
		return
	}

	// Находим первое строковое значение
	firstQuoteEnd := strings.Index(input[1:], "\"") + 1
	x := input[:firstQuoteEnd+1]
	operatorAndRest := strings.TrimSpace(input[firstQuoteEnd+1:])

	// Разделяем оставшуюся часть строки по пробелам
	fields := strings.Fields(operatorAndRest)
	if len(fields) < 2 {
		fmt.Println("Неверный формат ввода")
		return
	}

	operator := fields[0]
	y := strings.Join(fields[1:], " ")

	// Удаляем кавычки из x и y
	x = x[1 : len(x)-1]
	if strings.HasPrefix(y, "\"") && strings.HasSuffix(y, "\"") {
		y = y[1 : len(y)-1]
	}

	// Проверка и выполнение операций
	var result string
	var num int
	var err error
	if operator == "*" || operator == "/" {
		num, err = strconv.Atoi(y)
		if err != nil || num < 1 || num > 10 {
			fmt.Println("Пожалуйста, введите число от 1 до 10")
			return
		}
	}

	switch operator {
	case "+":
		result = sumStrings(x, y)
	case "-":
		result = subtractStrings(x, y)
	case "*":
		result = multiplyStringByNumber(x, num)
	case "/":
		result = divideStringByNumber(x, num)
	default:
		fmt.Println("Некорректный оператор")
		return
	}

	// Ограничение длины результата и вывод
	const maxLength = 40
	if len(result) > maxLength {
		cutResult := result[:maxLength] + "..."
		fmt.Println(cutResult)
	} else {
		fmt.Println("Результат:", result)
	}
}
