package main

import "fmt"

func main() {
	// no inheritance in golang; also no super or parent

	ikhtiar := User{"Ikhtiar", "ikhtiar@fakemail.com", true, 24}
	fmt.Println(ikhtiar)

	/*
		use %+v for print full object. %v use for a single value
	*/

	fmt.Printf("My Details are: %+v\n", ikhtiar) 
	fmt.Printf("My Name is: %v\n", ikhtiar.Name) 
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}