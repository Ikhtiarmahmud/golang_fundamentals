package main

import "fmt"

func main() {
	greeting()

	result := adder(10, 40)
	fmt.Println("Result: ", result)

	total, message := autoAdder(30, 40 , 50)
	fmt.Println(message, total)
}

// take multiple input, return multiple value
func autoAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	} 

	return total, "Total Sum"
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func greeting() {
	fmt.Println("Welcome to hacker earth!")
}