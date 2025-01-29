// the purpose of this file is to generate a text file and write math problems to it

package main

import (
	"fmt"
	"math/rand"
	"os"
)

// function for creating the problem sheet
func CreateProblemSheet() string {

	// variable declaration
	var fileName string

	// debugging message
	fmt.Println("Creating a Problem Sheet...")

	// ask for user input for a file name
	fmt.Print("Please enter a file name: ")
	fmt.Scanln(&fileName)
	fmt.Println("File Name: ", fileName+".txt"+"\n")

	// creating the text file
	file, err := os.Create(fileName + ".txt")

	// error handling
	if err != nil {
		panic(err)
	}

	// debugging message
	fmt.Print("Generated ", file.Name(), " ...")

	return file.Name()

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

// function for writing to a already generated file
func WriteProblemSheet() {

	// get the file name of the math sheet

}
