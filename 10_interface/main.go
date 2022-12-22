package main

import "fmt"

//interface -> Kumpulan dari defini satu atau lebih method / tipe data
			//abstrak yang tidak memilikiimplementasi langsung
			// biasanya digunakan sebagai kontrak

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}
func main() {
	var name Person
	name.Name = "Ridho"
	SayHello(name)
}
