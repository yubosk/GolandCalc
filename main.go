package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isRoman(number string) bool {
	var r, _ = regexp.MatchString("^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", number)
	return r
}

func integerToRoman(number int) string {

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func romanToInteger(s string) int {

	know := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	lengthOfString := len(s)
	lastElement := s[len(s)-1 : lengthOfString]
	var result int
	result = know[lastElement]
	for i := len(s) - 1; i > 0; i-- {
		if know[s[i:i+1]] <= know[s[i-1:i]] {
			result += know[s[i-1:i]]
		} else {
			result -= know[s[i-1:i]]
		}
	}
	return result
}

func calc(a, b int, exp string) (result int) {
	if exp == "+" {
		result = a + b
	} else if exp == "-" {
		result = a - b
	} else if exp == "/" {
		result = a / b
	} else if exp == "*" {
		result = a * b
	} else {
		panic("Не известная операция")
	}
	return
}

func strToInt(symbol string) (result int) {
	result, err := strconv.Atoi(symbol)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func isSameTypeOfNumber(number1, number2 string) bool {
	if isRoman(number1) == isRoman(number2) {
		return true
	} else {
		panic("Я работаю с любым, но только одним типом цифр")
	}
}

func isLessTen(number1, number2 int) bool {
	if (number1 <= 10) && (number2 <= 10) {
		return true
	} else {
		panic("Одно из чисел больше 10!")
	}
}

func main() {
	//чтение строки ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение для вычисления: ")
	line, _ := reader.ReadString('\n')
	//проверка корректности ввода
	if len(strings.Fields(line)) > 3 {
		panic("Формат математической операции не удовлетворяет заданию")
	} else if len(strings.Fields(line)) < 3 {
		panic("Cтрока не является математической операцией")
	}

	//считывание выражения
	words := strings.Split(line, " ")
	var a string = words[0]
	var expression string = words[1]
	var b string = strings.Trim(words[2], "\n")

	//if a == "" || expression == "" || b == "" {
	//	panic("Cтрока не является математической операцией")
	//}

	//Проверка: один вид цифр
	isSameTypeOfNumber(a, b)
	////Работа с римскими цифрами
	var result int
	var printResult string
	if isRoman(a) && isRoman(b) {
		var number1 int = romanToInteger(a)
		var number2 int = romanToInteger(b)
		isLessTen(number1, number2)
		result = calc(number1, number2, expression)
		if result < 0 {
			panic("Не бывает отрицательных римских цифр")
		}
		printResult = integerToRoman(result)
	} else {
		//вычисление арабскими цифрами
		var numberA = strToInt(a)
		var numberB = strToInt(b)
		isLessTen(numberA, numberB)
		result = calc(numberA, numberB, expression)
		printResult = strconv.Itoa(result)
	}

	fmt.Println("Результат вычисления:", printResult)
}
