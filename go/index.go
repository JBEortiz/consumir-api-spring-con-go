package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Cripto struct {
	Id        int    `json:"id"`
	Suplay    string `json:"suplay"`
	Price     string `json:"price"`
	Name      string `json:"name"`
	Bockchain string `json:"bockchain"`
}

func main() {
	gerAll()
	getById()

}
func gerAll() {
	response, err := http.Get("http://localhost:8080/criptos")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []Cripto
	json.Unmarshal(responseData, &responseObject)
	fmt.Println("GET ALL")
	fmt.Println(responseObject)
	for i := 0; i < len(responseObject); i++ {
		fmt.Println(responseObject[i])
	}

}
func getById() {
	response, err := http.Get("http://localhost:8080/cripto/2")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Cripto
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("GET BY ID")
	responseObject.Suplay += " %"
	fmt.Println(responseObject)

}
