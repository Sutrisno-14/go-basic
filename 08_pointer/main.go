package main

import (
	"fmt"
	"strings"
)

func main() {
	//pointer adalah sebuah variabel spesial yang digunakan untuk menyimpan
			//alamat memory variabel lainnya
	// * = arterisk  || & = ampersand

	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memory address) :", &firstNumber)
	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memory address) :", secondNumber)
	fmt.Println(&secondNumber)

	//Kesimpulan 
		// * = untuk mendapatkan value asli dari variabel pointer
		// & = untuk mendapatkan alamat memory dari variabel

	fmt.Println("\n", strings.Repeat("#", 40), "\n")
	//mengubah nilai melalui pointer
	var person1 string = "ridho"
	var person2 *string = &person1

	fmt.Println("person1 (value) :", person1)
	fmt.Println("person1 (memory address) :", &person1)
	fmt.Println("person2 (value) :", *person2)
	fmt.Println("person2 (memory address) :", person2)

	fmt.Println("\n", strings.Repeat("#", 40), "\n")

	*person2 = "maulana"
	fmt.Println("person1 (value) :", person1)
	fmt.Println("person1 (memory address) :", &person1)
	fmt.Println("person2 (value) :", *person2)
	fmt.Println("person2 (memory address) :", person2)

	fmt.Println("\n", strings.Repeat("#", 40), "\n")
	//pointer sebagai paramater

	var a int = 10

	fmt.Println("before a :", a)
	changeValue(&a)
	fmt.Println("after a :", a)
}

func changeValue(number *int) {
	*number = 20
}