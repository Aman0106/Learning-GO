package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type course struct {
	Name       string `json:"courseName"`
	Price      int
	Instructor string
	Password   string   `json:"-"`
	Tags       []string `json:"tags,omitempty"`
	Copies     int
}

func main() {
	DecodeJson()
}

func EncodetoJson() []byte {

	courses := []course{
		{"GO", 199, "Hitesh Chowdhary", "qwerty1", []string{"Go, Backend, web"}, 29},
		{"DSA", 149, "Striver", "qwerty2", []string{"C++, JAVA, Coding"}, 20},
		{"Android Native", 299, "Philip Lackner", "qwerty3", nil, 20},
	}

	jsonData, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	// fmt.Println("Json Data:", string(jsonData))

	return jsonData
}

func DecodeJson() {
	data := EncodetoJson()

	var stringResponse strings.Builder

	_, err := stringResponse.Write(data)
	if err != nil {
		panic(err)
	}

	// fmt.Println("byteCOunt:", byteCount)
	// fmt.Println("data:", stringResponse.String())

	var courses []course

	checkValid := json.Valid(data)

	if checkValid {
		json.Unmarshal(data, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("Notvalid")
	}
}
