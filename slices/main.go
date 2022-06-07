package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Array")

	// Note: this is slice, not array.
	var bookList = []string{"Clean-Code", "Programming-Fundamental"}

	// add new element in slices
	bookList = append(bookList, "Golang", "PHP")

	// remove first element
	bookList = append(bookList[1:])

	// take element between 1 to 2
	bookList = append(bookList[1:2])
	fmt.Println(bookList)

	scores := make([]int, 4); //slice size is 4
	
	scores[0] = 234
	scores[1] = 343
	scores[2] = 236
	scores[3] = 334

	scores = append(scores, 456, 434) // reallocate the memory and increase the slice size

	sort.Ints(scores) // sort the scores
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores)) // check sort or not


	// How to remove a value from slice based on index
	var courses = []string{"php", "java", "kotlin", "go", "node"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // remove kotlin
	fmt.Println(courses)
}