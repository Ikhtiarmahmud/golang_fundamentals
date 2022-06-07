package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// EncodeJson()
	DecodeJson()
}

type course struct {
	Name string `json:"coursename"`
	Price int `json:"price"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func EncodeJson() {
	courses := []course{
		{"PHP", 120, "12345", []string{"php", "web"}},
		{"Java", 220, "12345", []string{"java", "web"}},
		{"NodeJs", 320, "12345", []string{"nodejs", "web"}},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "PHP",
			"price": 120,
			"tags": ["php","web"]
		}
	`)

	var courses course

	checkValidJson := json.Valid(jsonDataFromWeb)

	if checkValidJson {
		fmt.Println("Json is valid")
		
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("Json is not valid")
	}

	// print json through key value pair
	var courseContent map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &courseContent)
	fmt.Printf("%#v\n", courseContent)
}