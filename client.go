package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"qunatity"`
}

func get() {
	Base_URL := "http://localhost:8081"
	response, err := http.Get(Base_URL + "/books")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var books []book

	json.Unmarshal(responseData, &books)
	fmt.Println(books)
}

func post() {

	Base_URL := "http://localhost:8081"
	postBody := book{ID: "4", Title: "My book 4", Author: "Jack 4", Quantity: 5}

	bodyBytes, err := json.Marshal(&postBody)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	requestBody := bytes.NewReader(bodyBytes)

	resp, err := http.Post(Base_URL+"/books", "application/json", requestBody)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	fmt.Println(s)
}

func main() {
	// post()
	get()
}
