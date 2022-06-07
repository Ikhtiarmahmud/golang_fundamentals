package main

import "fmt"

func main()  {
	mynum := 23

	ptr := &mynum

	fmt.Println("Pointer value is ", ptr)
	fmt.Println("Actual value is ", *ptr)
}