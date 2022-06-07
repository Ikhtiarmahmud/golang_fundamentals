package main

import "fmt"

func main() {
	fmt.Println("Array")

	var bookList [4]string

	bookList[0] =  "Clean Code"
	bookList[1] = "Programming Fundamental"
	bookList[2] = "Data Structure"
	bookList[3] = "Algorithm"

	fmt.Println("Book List is ", bookList)
	fmt.Println("Book List Array Lenght is ", len(bookList))

	// another way to declare an array

	var cities = [2]string{"Dhaka", "Chittagong"}
	fmt.Println("Cities are ", cities)
	fmt.Println("Array Length ", len(cities))
}