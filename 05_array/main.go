package main

import "fmt"

func main() {
	//array -> tipe data berisi kumpulan data dengan tipe yang sama,
			//perlu menentukan jumlah array yang ditampung
			// Daya tampung Array tidak bisa bertambah setelah Array dibuat
	var number = [5]int{1,2,3,4,5}
	fmt.Println(number)
	fmt.Println("========")
	var names [3]string

	names[0] = "Ridho"
	names[1] = "Maulana"
	names[2] = "Prasetian"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println("=============")

	//harizontal

	var fruits [4]string
	fruits = [4]string{"apple", "manggo", "melon", "banana"}
	fmt.Println("Kids of fruits : ", fruits)
	
	fmt.Println("=====for using array=======")

	//for using array
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("this is fruit : %s\n", fruits[i])
	}

	fmt.Println("=============")

	fruits = [4]string{
		"watermelon",
		"papaya",
		"Avocados",
		"Dragonrfruit",
	}
	fmt.Println("name of fruits: ", fruits)

	fmt.Println("=====for using for _ range=======")

	for j, fruit := range fruits {
		fmt.Printf("fruit fresh : %s number: %d\n", fruit, j)
	}

	fmt.Println("=====for using variael underscore in for _ range=======")

	for _, f := range fruits {
		fmt.Printf("This freash fruit : %s\n", f)
	}
	
	//array multidimensi
	fmt.Println("====array multidimensi======")

	var numbers = [2][3]int{[3]int{1,2,3}, [3]int{4,5,6}}
	fmt.Println(numbers)
	
	numbers = [2][3]int{{3,2,1}, {6,5,4}}
	fmt.Println(numbers)

	fmt.Println("==========")

	// alocation array using keyword make
	var frut = make([]string, 3) 
		frut[0] = "manggo"
		frut[1] = "banana"
		frut[2] = "pepaya"
	for _, fr := range frut {
		fmt.Println("fruit : ", fr)
	}
}