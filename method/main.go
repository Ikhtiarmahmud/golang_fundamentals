package main

import "fmt"

func main() {
	// no inheritance in golang; also no super or parent

	ikhtiar := User{"Ikhtiar", "ikhtiar@fakemail.com", false, 24}
	fmt.Println(ikhtiar)

	/*
		use %+v for print full object. %v use for a single value
	*/

	fmt.Printf("My Details are: %+v\n", ikhtiar) 
	fmt.Printf("My Name is: %v\n", ikhtiar.Name) 
	ikhtiar.getStatus()
	ikhtiar.newMail()
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

// create a method
// u is an object of User Struct
func (u User) getStatus() {
	fmt.Println("User status is: ", u.Status)
}

// u is a copy of User. It won't be change of origin email value
func (u User) newMail() {
	u.Email = "newTest@fakemail.com"
	fmt.Println("User new mail is: ", u.Email)
}