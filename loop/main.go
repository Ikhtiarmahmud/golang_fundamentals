package main

import "fmt"

func main() {
	days := []string{"Saturday", "Sunday", "Monday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, value := range days {
		fmt.Printf("Index is: %v and Value is: %v\n", index, value)
	}

	i := 1

	for i <= 10 {

		if i == 2 {
			goto fun
		}

		if i == 5 {
			i++
			continue
		}

		fmt.Println(i)
		i++
	}

	fun:
		fmt.Println("Jump on Goto")
}

