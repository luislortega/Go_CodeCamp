package main

import "fmt"

func main() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input) // scan text
	fmt.Print(input)   //show text
}
