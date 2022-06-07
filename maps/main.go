package main

import "fmt"

func main() {

	// maps in golang

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"

	fmt.Println("JS form is: ", languages["JS"])

	delete(languages, "PY") // delete element by key

	fmt.Println(languages)

	// loops in map
	for key, value := range languages {
		fmt.Printf("Key is %v, Value is %v\n", key, value)
	}

}