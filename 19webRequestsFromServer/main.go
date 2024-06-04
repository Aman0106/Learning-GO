package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	const myUrl = "http://localhost:8000"

	PerformPostFormRequest(myUrl + "/postform")
	// PerformPostJsonRequest(myUrl + "/post")
	// PerformGetRequest(myUrl + "/get")
}

func PerformGetRequest(myUrl string) {

	responseBody, err := http.Get(myUrl)

	handleError(err)

	dataBytes, err := io.ReadAll(responseBody.Body)
	handleError(err)

	var responseString strings.Builder

	byteCount, err := responseString.Write(dataBytes)

	fmt.Println("byte count:", byteCount)
	fmt.Println("Respoonse body:", responseString.String())

	// fmt.Println("Response body:", string(dataBytes))
	// fmt.Println("COntent length:", responseBody.ContentLength)

	defer responseBody.Body.Close()
}

func PerformPostJsonRequest(myUrl string) {

	requestBody := strings.NewReader(`
		{
			"Name": "Aman",
			"Age": 21,
			"isMale": true
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	handleError(err)
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	var responseString strings.Builder

	_, err = responseString.Write(databytes)
	handleError(err)

	fmt.Println("Post Response: ", responseString.String())

}

func PerformPostFormRequest(myUrl string) {

	data := url.Values{}

	data.Add("BName", "Amandeep")
	data.Add("Age", "23")
	data.Add("Cnumber", "25.6")

	response, err := http.PostForm(myUrl, data)
	handleError(err)
	defer response.Body.Close()

	dataByte, err := io.ReadAll(response.Body)
	handleError(err)

	var responseString strings.Builder
	_, err = responseString.Write(dataByte)
	handleError(err)

	fmt.Println("Response body:", responseString.String())

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
