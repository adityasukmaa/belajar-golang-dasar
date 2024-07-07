package main

import "fmt"

func main() {
	name := "april"

	switch name {
	case "adit":
		fmt.Println("Hallo Adit")
	case "wiranto":
		fmt.Println("Hallo Wiranto")
	case "april":
		fmt.Println("Hallo April")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Telalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
