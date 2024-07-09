package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is ", customer.Name)
}

func main() {
	var eko Customer
	fmt.Println(eko)

	eko.Name = "Aditya Sukma Pratama"
	eko.Address = "Margadadi"
	eko.Age = 20
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	april := Customer {
		Name : "Aprilliady",
		Address : "Paoman",
		Age : 20,
	}
	fmt.Println(april)

	wiranto := Customer {"Wiranto", "Karangampel", 20}
	fmt.Println(wiranto)

	wiranto.sayHello("Aditya")
	april.sayHello("Aditya")
	eko.sayHello("Aditya")
}
