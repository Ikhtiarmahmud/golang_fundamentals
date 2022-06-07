package main

import "fmt"

const Host string = "github.com" // Capital H mean it is public

func main() {
	// VARIABLES - TYPES AND CONSTANTS

    var username string = "Ikhtiar"
    fmt.Println("String Type Variable: ", username)

	var isLoggedIn bool = false
	fmt.Println("Boolean Type Variable:" , isLoggedIn)

	var smallVal uint8 = 255 // unsigned interger types
	var unsignedIntegerSixteen uint16 = 256 
	var typeInteger int = 10000
	
	 fmt.Println("Interger Types:", smallVal, unsignedIntegerSixteen, typeInteger)

	var smallfloat float32 = 234.00345  
	var typeFloat float64 = 234.00345345345
	fmt.Println("Float Types:", smallfloat, typeFloat)

	// IMPLICIT TYPE - TYPE DECALARED NOT MANDATORY

	var website = "ikhtiar-mitul.net"
	fmt.Println(website)

	// no var style - declare and initialize value
	// it won't work if we declare it outside function
	
	github := "github.com/ikhtiaruddin" 

	fmt.Println(github)


}