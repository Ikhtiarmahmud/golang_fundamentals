package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


const myurl = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()

	// response, err := http.Get(url)

	// if err != nil {
	// 	panic(err)
	// }

	// defer response.Body.Close() // caller's responsibility is close the connection. It's very importent

	// data, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	panic(err)
	// }

	// content := string(data)

	// fmt.Printf("The response data type is: %T\n", response)

	// fmt.Println(content)
}

func PerformGetRequest() {
	const url = "http://localhost:8000/get"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	/* strings is a package. It has a builder, & A builder is used to effeciently build a string
	   Write Methods. It minimize memory copying */
	var responseString strings.Builder

	data, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(data)

	// Using strings pakcage
	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String())

	// regular
	content := string(data)
	fmt.Println(content)
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	requestData := strings.NewReader(`
		{
			"name" : "Ikhtiar",
			"designation" : "Software Engineer",
			"stack" : "Backend"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestData)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/post"

	data := url.Values{}
	data.Add("name", "ikhtiar")
	data.Add("stack", "backend")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}