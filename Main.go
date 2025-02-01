package main

import (
	"fmt"
)

func main() {

	// debugging message
	fmt.Println("Running the Program")

	// creating two different sheets
	fileNameOne, fileNameTwo := CreateSheets()
	fmt.Println("Created ", fileNameOne)
	fmt.Println("Created ", fileNameTwo)

	// writing out the problem sheet
	WriteProblemSheet(fileNameOne)
	fmt.Println("Written to the Problem Sheet")
	WriteAnswerSheet(fileNameTwo, fileNameOne)

}
