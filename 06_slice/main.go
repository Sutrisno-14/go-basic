package main

import "fmt"

func main() {
	//data slice is part of array
	//slice same with array, but slice can be changed
	//slice and array always connected

	//slice has 3, pointer, length and capacity

	var (
		fruitA = []string{"apple","melon","banana", "dragonfruit","watermelon"}    //slice
		fruitB = [...]string{"pepaya", "pear", "strobery","melon","banana"}	//Array
		fruitC = [2]string{"banana", "papaya"} //array
		names = [...]string{"maulana", "prasetian", "reinaldo", "dimas", "ridho", "sutrisno"}  //array
		days = [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "monday"}  //array
	)
	sliceA := fruitA[:3]
	slice := names[4:6]

	fmt.Println(sliceA)
	fmt.Println("THIS IS SLICE",fruitA)
	fmt.Println("THIS IS ARRAY", fruitB)
	fmt.Println("THIS IS ARRAY",fruitC)
	fmt.Println(fruitB[0:2])
	fmt.Println("===========")
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(len(names))   // len is to get length 
	fmt.Println(cap(names))	  // cap is to get capacity
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println("==============")
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "newSaturday"
	daySlice1[1] = "newMonday"

	fmt.Println(days)

	fmt.Println("==============")

	daySlice2 := append(daySlice1, "newVocation")  // append is making new slice and will add in the last position
	fmt.Println(daySlice2)
	daySlice2[0] = "upps"
	fmt.Println(daySlice2)
	daysNew := days[:]
	newDays := append(daysNew,"newsDay")
	fmt.Println(days)
	fmt.Println(newDays)

	// make slice
	fmt.Println("======make slice======")

	newSlice := make([]string,3 ,3)
	newSlice[0]="ridho"
	newSlice[1]= "maulana"
	newSlice[2]= "prasetian"


	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice from days

	fmt.Println("====copy slice====")
	formSlice := days[:]
	toSlice := make([]string,len(formSlice), cap(formSlice))
	copy(toSlice, formSlice)

	fmt.Println(formSlice)
	fmt.Println(toSlice)
}
