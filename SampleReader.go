package main

import (
	"fmt"
	"math/rand"
	"os"
)

func CreateFile() {

	fmt.Println("Writing file")
	file, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	length, err := file.WriteString("welcome to golang \n" +
		"demonstrates reading and writing operations to a file in golang")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File Name: %s", file.Name())
	fmt.Printf("\nfile length: %d\n", length)

}

// CreateTestDocument creates a new file named "TestDocument.txt" in current directory
// the method writes and string to it and prints out the length of the strings in the file
func CreateTestDocument() {

	fmt.Println("Creating a Test Document with Sample Text")
	file, err := os.Create("TestDocument.txt")

	if err != nil {
		panic(err)
	}

	length, err := file.WriteString("Hi there! This is my first test file")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated File Name: %s\n ", file.Name())
	fmt.Printf("Length of the File: %d", length)

}

func CreateMathSheet() {

	// taking in a filename and then printing it out
	var fileName string
	fmt.Println("Creating a Math Test Sheet with Sample Text")
	fmt.Print("Please enter a file name: ")
	fmt.Scanln(&fileName)
	fmt.Println("File Name: ", fileName+".txt"+"\n")

	// creating the text file
	file, err := os.Create(fileName)

	// error handling
	if err != nil {
		panic(err)
	}

	fmt.Print("Generated " + file.Name() + "\n")

	var numberOne int
	var numberTwo int
	var totalNumber int

	for i := 0; i < 10; i++ {

		// the necessary variables
		numberOne = rand.Intn(100)
		numberTwo = rand.Intn(100)
		totalNumber = numberOne + numberTwo

		// writes the file
		_, err := file.WriteString(fmt.Sprintf("%d + %d = %d\n", numberOne, numberTwo, totalNumber))

		if err != nil {
			panic(err)
		}

	}

}
