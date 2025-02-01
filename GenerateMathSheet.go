// the purpose of this file is to generate a text file and write math problems to it

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
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
func WriteAnswerSheet(answerFile string, problemFile string) {

	// checking for presence of answer file
	if _, err := os.Stat(answerFile); err == nil {
		fmt.Println(answerFile + " exists!")
	}

	// opening up the problem file
	fileTwo, err := os.Open(problemFile)

	// error handling
	if err != nil {
		panic(err)
	}

	// closing the file
	defer fileTwo.Close()

	// opening up the answer file
	fileOne, err := os.OpenFile(answerFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	// error handling
	if err != nil {
		panic(err)
	}

	// closing the file
	defer fileOne.Close()

	for i := 0; i < 10; i++ {

		// retrieving the contents of the file
		num1, operator, num2, totalNumber := ReturnFileContents(problemFile)

		// creating the math string
		mathString := strconv.Itoa(num1) + " " + operator + " " + strconv.Itoa(num2) + " = " + strconv.Itoa(totalNumber)

		// writing to the answer file
		_, writeError := fileOne.WriteString(mathString + "\n")

		// handling the error
		if writeError != nil {
			panic(writeError)
		}

	}
}

// func GeneratePDF(textFileOne string) {

// 	// creating a PDF, adding a page and setting font
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()
// 	pdf.SetFont("Arial", "B", 16)

// 	// setting the cursor location
// 	pdf.MoveTo(10, 10)

// 	// setting the title and moving to the next line
// 	pdf.Cell(190, 10, "Math Problems and Solutions")
// 	pdf.Ln(10)

// 	// setting the font for the body text
// 	pdf.SetFont("Arial", "", 12)

// 	// opening the text file and reading line by line
// 	fileOne, err := os.Open(textFileOne)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fileOne.Close()

// }

func ReturnFileContents(textFileOne string) (int, string, int, int) {

	var num1 int
	var num2 int
	var operator string
	var totalNumber int

	// opening up the problem file
	fileOne, err := os.Open(textFileOne)

	// error handling
	if err != nil {
		panic(err)
	}

	// closing the file
	defer fileOne.Close()

	// reading the lines of the problem file
	scanner := bufio.NewScanner(fileOne)
	for scanner.Scan() {

		// getting rid of the spaces and newlines
		line := strings.TrimSpace(scanner.Text())

		// printing out the raw line
		fmt.Printf("Raw line: '%s'\n", strings.TrimSpace(scanner.Text()))

		// using regex to find two numbers
		re := regexp.MustCompile(`^(\d+)\s*([+\-*/])\s*(\d+)\s*=?$`)

		// extracting the matches
		match := re.FindStringSubmatch(line)

		// error handling
		if match == nil {
			fmt.Println("Skipping invalid line:", line) // Debugging print
			continue
		}

		// extracting the first number and error handling
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		// extracting the second number and error handling
		num2, err := strconv.Atoi(match[3])
		if err != nil {
			panic(err)
		}

		// extracting the operator
		operator := match[2]

		// debugging print statement
		fmt.Println("Extracted: ", num1, num2, operator)

		// finding the final number
		totalNumber = PerformMathOperation(num1, num2, operator)

	}

	return num1, operator, num2, totalNumber

}
