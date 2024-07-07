package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Aditya"
	middleName = "Sukma"
	lastName = "Pratama"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName, middleName, lastName)
}
