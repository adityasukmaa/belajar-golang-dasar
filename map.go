package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string {
	// 	person["name"] = "Aditya"
	// 	person["address"] = "Indramayu"
	// }

	person := map[string]string {
		"name" : "Aditya",
		"address" : "Indramayu",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string) 
	book["title"] = "Book Go-lang"
	book["author"] = "Aditya Sukma"
	book["upss"] = "Salah"

	fmt.Println(book)

	delete(book, "upss")

	fmt.Println(book)
}