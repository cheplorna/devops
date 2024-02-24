package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var person1 Person
	var person2 Person

	person1.name = "lorna"
	person1.age = 31
	person1.job = "Engineer"
	person1.salary = 100000

	person2.name = "lorney"
	person2.age = 30
	person2.job = "Engineer"
	person2.salary = 120000

	person3 := Person{
		name:   "Tiff",
		age:    28,
		job:    "sales",
		salary: 80000,
	}

	fmt.Println(person1)

	fmt.Println(person2)

	fmt.Println(person3)

	//validate url

	u, err := url.ParseRequestURI("https://google.com")
	if err != nil {
		//os.Exit(1)
		panic(err)
	}
	fmt.Println(u)

	resp, err := http.Get("https://google.com")
	if err != nil {
		//os.Exit(1)
		panic(err)
	}

	fmt.Println(resp, err)

}
