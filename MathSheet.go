package main

import (
	"fmt"
	"os"
)

func GenerateSheet() {

	fmt.Println("Creating a file...")

	// creating the math sheet
	file, err := os.Create("Math Sheet.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File Name: %s", file.Name())

}

func WriteSheet() {

	fmt.Println("Writing the file... ")

}
