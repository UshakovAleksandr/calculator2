package main

import (
	"bufio"
	"calculator/calc"
	"calculator/funcs"
	"fmt"
	"os"
)

func main() {
	const add string = "+"
	const sub string = "-"
	const mult string = "*"
	const div string = "/"

	var firstNumber float64
	var secondNumber float64
	var mathSign string

	reader := bufio.NewReader(os.Stdin)
	for {
		number, err := funcs.ConvertToFloat(funcs.InputNumber(*reader))
		if err != nil {
			fmt.Println("Введено неверное значение, повторите ввод")
			continue
		}
		firstNumber = number
		break
	}

	for {
		mathSign = funcs.InputMathSign(*reader)
		if funcs.CheckMathSign(mathSign, add, sub, mult, div) {
			break
		}
		fmt.Println("Введено неверное значение, повторите ввод")
		continue
	}

	for {
		number, err := funcs.ConvertToFloat(funcs.InputNumber(*reader))
		if err != nil {
			fmt.Println("Введено неверное значение, повторите ввод")
			continue
		}
		secondNumber = number
		break
	}

	c := calc.CreateCalc(firstNumber, secondNumber, mathSign)
	result := c.Calculate(c.NumberOne, c.NumberTwo, c.MathSign)

	switch mathSign {
	case add:
		fmt.Printf("Результат сложения: %v\n", result)
	case sub:
		fmt.Printf("Результат вычитания: %v\n", result)
	case mult:
		fmt.Printf("Результат умножения: %v\n", result)
	case div:
		if secondNumber == 0 {
			fmt.Println("Деление на 0 запрещено")
		} else {
			fmt.Printf("Результат деления: %v\n", result)
		}
	}
}
