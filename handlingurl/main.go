package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.daraz.com.bd/catalog/?q=test&_keyori=ss&from=input&spm=a2a0e.home.search.go.28d312f7u4wogq"

func main() {
	// Handling Urls in Golang

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println(qparams["q"])

	for key, val := range qparams {
		fmt.Printf("Index is %v and value is %v\n", key, val)
	}

	// make an url
	partsOfUrl := &url.URL{
		Scheme: "Https",
		Host: "devskill.com",
		Path: "Course",	
	}

	devurl := partsOfUrl.String()

	fmt.Println(devurl)
}