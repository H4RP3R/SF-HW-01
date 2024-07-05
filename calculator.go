package main

import (
	"errors"
	"fmt"
	"log"
)

func calc(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	}

	err := errors.New("unknown operator")
	return 0, err
}

func getUserInput(variable interface{}, msg string) error {
	fmt.Printf("%s: ", msg)
	_, err := fmt.Scanln(variable)
	return err
}

func main() {
	log.SetFlags(0)

	var (
		firstNum  float64
		secondNum float64
		operator  string
	)

	err := getUserInput(&firstNum, "Enter first number")
	if err != nil {
		log.Fatal(err)
	}

	err = getUserInput(&operator, "Enter operator")
	if err != nil {
		log.Fatal(err)
	}

	err = getUserInput(&secondNum, "Enter second number")
	if err != nil {
		log.Fatal(err)
	}

	res, err := calc(firstNum, secondNum, operator)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v %s %v = %v\n", firstNum, operator, secondNum, res)
}
