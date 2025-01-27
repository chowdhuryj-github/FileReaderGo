package main

import (
	"math/rand"
)

// function to return a random operator
func RandomOperator() string {

	randomOperator := []string{"+", "-", "/", "*"}
	randomIndex := rand.Intn(len(randomOperator))
	chooseOperator := randomOperator[randomIndex]
	return chooseOperator

}

// function to return a random number
func RandomNumber() (int, int) {

	var numberOne int
	var numberTwo int

	numberOne = rand.Intn(100)
	numberTwo = rand.Intn(100)

	return numberOne, numberTwo

}

// function to perform the math
func MathPerformer() int {

	var totalNumber int

	operator := RandomOperator()
	numberOne, numberTwo := RandomNumber()

	// swithcing between chosen operations
	switch operator {
	case "+":
		totalNumber = numberOne + numberTwo
	case "-":
		totalNumber = numberOne - numberTwo
	case "*":
		totalNumber = numberOne * numberTwo
	case "/":
		totalNumber = numberOne / numberTwo
	}

	// returning the final operations
	return totalNumber
}
