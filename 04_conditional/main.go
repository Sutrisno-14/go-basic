package main

import "fmt"

func main() {
	//conditiional statements
	// if, else if and else
	//switch - case

	// using if

	var point int = 20

	if point > 90 {
		fmt.Println("A+")
	}else if point > 80 {
		fmt.Println("A")
	}else if point > 70 {
		fmt.Println("B")
	}else if point > 60 {
		fmt.Println("C")
	}else if point > 50 {
		fmt.Println("D")
	}else {
		fmt.Println("E")
	}

	fmt.Println("=============")

	var currentYear = 2022

	if age := currentYear-1998; age < 17 {
		fmt.Println("Your are not able to make SIM")
	}else {
		fmt.Println("You are able to make SIM")
	}

	fmt.Println("=============")

	//switch case 

	var scor int = 9

	switch scor {
	case 10:
		fmt.Println("Perfect")
	case 9,8:
		fmt.Println("Good Job")
	case 7:
		fmt.Println("Enough")
	case 6,5:
		fmt.Println("Learn a lot")
	default:
		fmt.Println("Keep Going")
	}

	fmt.Println("=============")

	//switch with relational operators

	var score = 0

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8 && score > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study hard")
			fmt.Println("You need to learn more")
		}
	}

	fmt.Println("=============")
	
	//nested conditions

	var scores = 10

	if scores > 7 {
		switch scores {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	}else {
		if scores == 5 {
			fmt.Println("not bad!")
		}else if scores == 3 {
			fmt.Println("keep going")
		}else {
			fmt.Println("you arble to do it")
			if scores == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}