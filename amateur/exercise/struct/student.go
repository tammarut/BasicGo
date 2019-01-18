package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Id        string
	Firstname string
	Lastname  string
	No        int
	Phone     string
}

func main() {
	studentLists := []student{}

	student1 := student{
		Id:        "605233000",
		Firstname: "Tammarut",
		Lastname:  "Nualmeechue",
		No:        5,
		Phone:     "087-1542740",
	}
	student2 := student{
		Id:        "5099992322",
		Firstname: "Bruce",
		Lastname:  "Lee",
		No:        6,
		Phone:     "094-554618",
	}
	student3 := student{
		Id:        "1011101010",
		Firstname: "The",
		Lastname:  "Flash",
		No:        99,
		Phone:     "011-1111111",
	}
	studentLists = append(studentLists, student1)
	studentLists = append(studentLists, student2)
	studentLists = append(studentLists, student3)

	studentListsJSON, err := json.Marshal(studentLists)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(studentListsJSON))
}
func init() {
	fmt.Println("Hi, student")
	fmt.Println("--------------------------")
}
