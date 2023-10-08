package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToInt(s string) int {

	characterMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}
	length := len(s)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return characterMap[s[0]]
	}
	sum := characterMap[s[length-1]]
	for i := length - 2; i >= 0; i-- {
		if characterMap[s[i]] < characterMap[s[i+1]] {
			sum -= characterMap[s[i]]
		} else {
			sum += characterMap[s[i]]
		}
	}
	return sum
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func main() {

	var a int
	var b int
	var operator string
	var input string
	var con []string
	var result int

	fmt.Println("Введите выражение:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	con = strings.Fields(input)

	counter := 0
	a, err1 := strconv.Atoi(con[0])
	if err1 != nil {
		a = romanToInt(con[0])
		counter += 1
	}
	b, err2 := strconv.Atoi(con[2])
	if err2 != nil {
		b = romanToInt(con[2])
		counter += 1
	}
	operator = con[1]

	if counter != 1 {
		if a > 0 && a <= 10 && b > 0 && b <= 10 {
			if operator == "+" {
				result = a + b
			} else if operator == "-" {
				result = a - b
			} else if operator == "*" {
				result = a * b
			} else if operator == "/" {
				result = a / b
			} else {
				fmt.Println("Ошибка при вводе знака!")
			}
		} else {
			fmt.Println("Работаю только с арабскими и римскими числами от 1 до 10!")
		}
	} else {
		fmt.Println("Введенные числа должны быть одного типа!")
	}
	if counter == 2 && a > 0 && a <= 10 && b > 0 && b <= 10 {
		if result < 1 {
			fmt.Println("Римские числа не могут быть отрицательными или равны нулю!")
		} else {
			fmt.Println(intToRoman(result))
		}
	}
	if counter == 0 && a > 0 && a <= 10 && b > 0 && b <= 10 {
		fmt.Println(result)
	}

}
