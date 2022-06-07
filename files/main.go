package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "I don't wanna go with golang!"

	file, err := os.Create("./mytext.txt")

	throwErr(err)

	length, err := io.WriteString(file, content)

	throwErr(err)

	fmt.Println("Length: ", length)

	/*
		==-- Defer Keyword --==
		When we use defer, the line will execute end of the code. Defer uses stack data structure.
	*/

	defer file.Close()

	readFile("./mytext.txt")
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	throwErr(err)

	fmt.Println("File text is: ", string(data))
}

func throwErr(err error) {
	if err != nil {
		panic(err)
	}
}