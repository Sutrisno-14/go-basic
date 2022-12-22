package main

import "fmt"

func main() {
	//map in golang

	person := map[string]string{
		"name": "ridho",
		"address": "cilegon",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Printf("my name is : %s and my addrees in %s\n", person["name"], person["address"])
	fmt.Println(len(person))
	person["address"] = "Lampung"
	fmt.Println(person["address"])

	//inisialisasi value of map
	var names map[int]string

	names = map[int]string{}
	names[0]="ridho"

	data := map[int]string {
		1: "ridho",
		2: "maulana",
		3: "prasetian",
	}

	fmt.Println(names[0])
	fmt.Println(data)

	//iterasi item map using for range
	fmt.Println("====using for range=====")

	days := map[int]string{
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
		7: "sunday",
	}

	for i, d := range days {
		fmt.Println(i, "=", d)
	}
	fmt.Println("==========")
	for j :=1; j <= len(days); j++ {
		fmt.Println(j, "=", days[j])
	}

	
	fmt.Println("=============")
	
	//making new map using make(map[typeKey]typeValue) 
	
	book := make(map[string]string)
	// var book = make(map[string]string)     

	book["title"] = "Golang"
	book["author"] = "ridho"
	book["wrong"] = "ups"
	
	fmt.Println(book)
	
	//delete book
	delete(book, "wrong")
	fmt.Println(book)
	
	//example delete from days
	fmt.Println(days)
	delete(days, 7)
	fmt.Println(days)
	
}