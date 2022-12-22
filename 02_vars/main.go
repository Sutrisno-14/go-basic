package main

import "fmt"

func main() {
	// variabel in golang

	//variabel using var

	var name string = "sutrisno"
	var address, major, email string

	address = "south sumatra"
	major = "engineering computer"
	email = "sutrisno***@gmail.com"

	println(name)

	//shorthand variabel
	fullName := "ridho mp"
	age, phone := 21, "085769871265"
	fmt.Println(fullName)
	fmt.Printf("age : %d\n", age)  // %d represent value type of number
	fmt.Printf("phone : %s\n", phone)

	

	//multiple variabel

	fmt.Printf("addres : %s\n", address) // %s represent value type of string
	fmt.Printf("major : %s\n", major)
	fmt.Printf("email : %s\n", email)
}