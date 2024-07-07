package main

import "fmt"

func main() {

	type NoKTP string

	var ktpAdit NoKTP = "3212150806040002"

	var contoh string = "3212150907050002"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpAdit)
	fmt.Println(contohKTP)

}