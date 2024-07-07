package main

import "fmt"

func main() {
	const (
		firstName string = "Aditya"
		lastName         = "Pratama"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Tidak bisa diubah"
	// lastName = "Tidak bisa diubah"
}
