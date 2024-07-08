package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Ini Panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("UPSS ERROR")
	}
}

func main() {
	runApp(true)
	fmt.Println("Aditya Sukma Pratama")
}
