package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	login := true

	if login {
		fmt.Println("You are logged in!")
	} else {
		fmt.Println("Please log in first!")
	}

	if num := 11; num < 10 {
		fmt.Println("Value is less than 10")
	} else if num == 10 {
		fmt.Println("Value is equals to 10")
	} else {
		fmt.Println("Value is greater than 10")
	}

	/*
	    ====================== 
		      Switch Case
		======================
	*/

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1

	switch diceNum {
		case 1:
			fmt.Println("Dice Number is: ", diceNum)
		case 2:
			fmt.Println("Dice Number is: ", diceNum)
		case 3:
			fmt.Println("Dice Number is: ", diceNum)
		case 4:
			fmt.Println("Dice Number is: ", diceNum)
		case 5:
			fmt.Println("Dice Number is: ", diceNum)
		case 6:
			fmt.Println("Dice Number is: ", diceNum)
		default:
			fmt.Println("Out of board :D")
	}
}