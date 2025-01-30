// the purpose of this file is to generate a text file and write math problems to it

package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// function for creating the problem sheet
func CreateSheets() (string, string) {

	// variable declaration
	var problemFile string
	var answerFile string

	// debugging message
	fmt.Println("Creating a Problem Sheet...")

	// ask for user input for a file name
	fmt.Print("Please enter a file name: ")
	fmt.Scanln(&problemFile)
	fmt.Println("File Name: ", problemFile+".txt"+"\n")

	// generating the name of answer file
	answerFile = problemFile + "Answer"

	// creating the text file
	fileOne, err := os.Create(problemFile + ".txt")
	fileTwo, err := os.Create(answerFile + ".txt")

	// error handling
	if err != nil {
		panic(err)
	}

	// debugging message
	fmt.Print("Generated ", fileOne.Name(), " ...")
	fmt.Print("Generate ", fileTwo.Name(), " ...")

	return fileOne.Name(), fileTwo.Name()

}

// function for generating two random numbers
func GenerateRandomNumbers() (int, int) {

	// declaring the variables
	var numberOne int
	var numberTwo int

	// assigning random numbers to variables
	numberOne = rand.Intn(100)
	numberTwo = rand.Intn(100)

	// returning the numbers
	return numberOne, numberTwo

}

// function for generating a random operator
func GenerateRandomOperator() string {

	// list of operators
	randomOperator := []string{"+", "-", "/", "*"}

	// choosing a random index from the length
	randomIndex := rand.Intn(len(randomOperator))

	// choosing the random operator
	chooseOperator := randomOperator[randomIndex]

	// returning the random operator
	return chooseOperator

}

// function for performing the math calculation
func PerformMathOperation(numberOne int, numberTwo int, operator string) int {

	var totalNumber int

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

	// returning the final number
	return totalNumber

}

// function for writing to a already generated file
func WriteProblemSheet(fileName string) {

	// initializing the string builder
	var builder strings.Builder

	// get the file name of the math sheet
	fmt.Println("Writing to ", fileName)

	// opening the file
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)

	// error handling
	if err != nil {
		panic(err)
	}

	// generating ten word problems with answers
	for i := 0; i < 10; i++ {

		// retrieving randomly generated number
		operator := GenerateRandomOperator()
		numberOne, numberTwo := GenerateRandomNumbers()

		// retreiving the final answer
		// answer := PerformMathOperation(numberOne, numberTwo, operator)

		// building the string to write
		builder.WriteString(strconv.Itoa(numberOne) + " ")
		builder.WriteString(operator + " ")
		builder.WriteString(strconv.Itoa(numberTwo) + " ")
		builder.WriteString("= ")
		// builder.WriteString(strconv.Itoa(answer))
		builder.WriteString("\n")

		// writing to the file
		_, err := file.WriteString(builder.String())

		// error handling
		if err != nil {
			panic(err)
		}

		// resetting the builder
		builder.Reset()

	}

	// closing the file
	defer file.Close()

}

// function for writing to a already generated file
func WriteAnswerSheet(fileName string) {

	// accessing the problem file
	problemFile, err := os.Open(fileName)

	// error handling
	if err != nil {
		panic(err)
	}

	// closing the file
	defer problemFile.Close()

	// accessing the answer file
	answerFile, err := os.Open(fileName)

	// error handling
	if err != nil {
		panic(err)
	}

	// closing the file
	defer answerFile.Close()

	// copying contents from problem to answer file
	_, err = io.Copy(answerFile, problemFile)

	// error handling
	if err != nil {
		panic(err)
	}

	fmt.Print("Contents successfully copied!")

}
