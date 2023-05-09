package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	var a, b float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	var result float64
	var err error

	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result, err = divide(a, b)
	default:
		fmt.Println("Invalid operator")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
