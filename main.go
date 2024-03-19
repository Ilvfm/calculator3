package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MATH  = "Введите математическую операцию."
	SINT  = "Два числа и знак между ними."
	RANGE = "Только арабские и римские цифры от 0 до 10"
	SCALE = "Разные системы."
	ZERO  = "В римской нет 0 и отриц.чисел."
)

var romanMap = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convertIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var stringExp []string

func main() {

	fmt.Println("Приветствую вас,я будущий супер калькуляторф!\n" +
		"Правила моего использования:\n" +
		"1. Введите в меня два числа и знак между ними.\n" + "2. Пока что могу и арабские,и римские считать\n" +
		"3. Но учтите,только такие знаки воспринимаю: '+', '-', '*', '/'")
	for {
		fmt.Println("Что считаем?: ")
		text := reader()
		operator, stringsCheck, numbers, romanNumber := parse(strings.ToUpper(strings.TrimSpace(text)))
		operation(operator, stringsCheck, numbers, romanNumber)
	}
}

func reader() string {
	reader := bufio.NewReader(os.Stdin)
	var text string
	text, _ = reader.ReadString('\n')
	text = strings.ReplaceAll(text, " ", "")
	return text
}

func parse(text string) (string, int, []int, []string) {
	var operator string
	var stringsCheck int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	for i := range operators {
		for _, val := range text {
			if i == string(val) {
				operator += i
				stringExp = strings.Split(text, operator)
			}
		}
	}

	for _, elem := range stringExp {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsCheck++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}
	return operator, stringsCheck, numbers, romans
}

func funcArab(numbers []int, operator string) {
	if checkRange(numbers[0]) && checkRange(numbers[1]) {
		val := operators[operator]
		a, b = &numbers[0], &numbers[1]
		fmt.Printf("Ответ: \n%v\n", val())
	} else {
		panic(RANGE)
	}
}

func funcRoman(romans []string, operator string) {
	romanToInt := make([]int, 0)
	for _, elem := range romans {
		val := romanMap[elem]
		if checkRange(val) {
			romanToInt = append(romanToInt, val)
		} else {
			panic(RANGE)
		}
	}
	if val, ok := operators[operator]; ok {
		a, b = &romanToInt[0], &romanToInt[1]
		intToRoman(val())
	}
}

func intToRoman(tmpRes int) {
	var romanRes string
	if tmpRes == 0 || tmpRes < 0 {
		panic(ZERO)
	}
	for tmpRes > 0 {
		for _, elem := range convertIntToRoman {
			for i := elem; i <= tmpRes; {
				for index, value := range romanMap {
					if value == elem {
						romanRes += index
						tmpRes -= elem
					}
				}
			}
		}
	}
	fmt.Printf("Ответ: \n%v\n", romanRes)
}

func operation(operator string, stringsCheck int, numbers []int, romans []string) {
	checkOperator(operator)
	switch stringsCheck {
	case 1:
		panic(SCALE)
	case 0:
		funcArab(numbers, operator)
	case 2:
		funcRoman(romans, operator)
	}
}

func checkOperator(operator string) {
	switch {
	case len(operator) > 1:
		panic(SINT)
	case len(operator) < 1:
		panic(MATH)
	}
}

func checkRange(num int) bool {
	if num > 0 && num < 11 {
		return true
	} else {
		return false
	}
}
