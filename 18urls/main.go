package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3921/learn?coursename=GO&id=hfhdsjnsdjkf"

func main() {

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	// fmt.Println("resutlt: ", result.RawQuery)

	qParams := result.Query()

	for key, value := range qParams {
		fmt.Printf("%v : %v\n", key, value)
	}

	//Constructig a URL

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "loc.dev",
		Path:   "learn",
	}

	fmt.Println("Constructed Url :", partsOfUrl)
}
