package main

import "fmt"

func main() {
	name := "wiranto"

	if name == "adit" {
		fmt.Println("HALLO ADITYA")
	} else if name == "april" {
		fmt.Println("Hallo April")
	} else if name == "wiranto" {
		fmt.Println("Hallo Wiranto")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	
	if length := len(name); length > 5 {
		fmt.Println("Nama telalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}