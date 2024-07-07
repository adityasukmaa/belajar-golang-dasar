package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Aditya")
	fmt.Println(result)

	fmt.Println(getHello("April"))
	fmt.Println(getHello("Wiranto"))
}