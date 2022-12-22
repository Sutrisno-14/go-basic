package main

import "fmt"


func main() {
	//First way of looping -> keyword for
	var data int = 3

	for i :=0; i < data; i++ {
		fmt.Println("This is number: ", i )
	}

	fmt.Println("================")

	//Second way of looping -> keyword while
	var i = 0
	
	for i < data {
		fmt.Println("This is number start from :", i)
		i++
	}

	fmt.Println("================")

	//Third way of looping -> keyword break
	var j = 0;
	
	for {
		fmt.Println("Number :", j)
		j++

		if j == 3 {
			break
		}
	}

	fmt.Println("================")

	//Loopings (break and continue keyword)
	for a :=1; a <= 10; a++ {
		if a%2 == 1 {
			continue  // if the condition is true, so the progress will be keep to continue but the number will be not display
		}

		if a > 8 {
			break   // if the condition is true, so the looping will be break/ stop and number > 8 is not executed
		}

		fmt.Println("My Number :", a) // result = 2, 4, 6, 8
	}

	fmt.Println("================")

	// Loopings (nested looping)
	for b :=0; b < 5; b++ {
		for c :=b; c < 5; c++ {
			fmt.Print(c, " ") // string " "-> making space // c -> making column
		}

		fmt.Println()        // making row
	}

	fmt.Println("================")
	
	//Looping label

	outerloop:
		for d := 0; d <3; d++ {
			fmt.Println("This is looping : - ", d+1)  // +1 for looping - start from number 1
			for e := 0; e < 3; e++ {
				if d == 2 {
					break outerloop
				}
				fmt.Print(e, " ")
			}
			fmt.Print("\n")

		}

}