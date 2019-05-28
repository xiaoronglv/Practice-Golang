package main

import (
	"encoding/json"
	"fmt"
	_ "os"
)

type student struct {
	name   string
	gender int
	age    int
}

type teacher struct {
	Name    string `json: name`
	Subject string `json: subject`
}

func main() {
	blobB, _ := json.Marshal(true)
	fmt.Println(string(blobB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	student := student{"Timothy Keller", 1, 18}
	teacher := teacher{"Ryan Lv", "CS"}

	studentB, _ := json.Marshal(student)
	fmt.Println(student)
	fmt.Println(string(studentB))

	teacherB, _ := json.Marshal(teacher)
	fmt.Println(string(teacherB))

}
