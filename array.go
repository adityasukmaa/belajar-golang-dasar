package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Aditya"
	name[1] = "Sukma"
	name[2] = "Pratama"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [...]int {
		90,
		80,
		70,
		100,
		120,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}