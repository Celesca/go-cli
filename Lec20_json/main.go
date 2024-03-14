package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	var input string
	fmt.Println("Enter 'w' to write to JSON, 'r' to read from JSON : ")
	fmt.Scan(&input)

	if input == "w" {
		writeJSON()
	} else {
		readJSON()
	}
}

func writeJSON() {
	var newPerson Person
	fmt.Println("Enter a new firstname : ")
	fmt.Scan(&newPerson.Firstname)
	fmt.Println("Enter a new lastname : ")
	fmt.Scan(&newPerson.Lastname)
	fmt.Println("Enter a new age : ")
	fmt.Scan(&newPerson.Age)

	var people []Person
	people = append(people, newPerson)

	newBS, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("data.json", newBS, 0644)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range people {
		fmt.Println(v.Firstname, v.Lastname, v.Age)
	}
}

func readJSON() {

	bs, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	var people []Person
	err = json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range people {
		fmt.Println(v.Firstname, v.Lastname, v.Age)
	}

}
