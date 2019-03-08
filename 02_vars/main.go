package main

import "fmt" //Evit print like c with variable definition

func main() {
	//variables like javascript
	var test = 12
	var testing = "Luis Leon"
	const pi = 3.1416

	//Shorthand ¿?
	name, email := "Luis Leon", "leon.test@gmail.com"
	fmt.Println(test)
	fmt.Println(testing)
	fmt.Println(pi)
	fmt.Println(name)
	fmt.Print(email)
	//Or
	fmt.Println(test, testing, pi, name, email)
}
