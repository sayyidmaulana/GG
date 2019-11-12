package main

import "fmt"

func main() {
//	type 1
	var firstName string = "John"
//	type 2
	var middleName string
	middleName = "Mikel"
//	type 3 
	lastName := "Wick"

fmt.Printf("hello %s %s %s!\n", firstName, middleName, lastName)
}