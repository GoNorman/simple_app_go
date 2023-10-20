package main

import (
	"fmt"

	"rsc.io/quote"
)

type mile int
type ID int

type Contact struct {
	phone string
}

type User struct {
	id           ID
	name         string
	contact_info Contact
}

func main() {
	message := quote.Hello()
	fmt.Println(message)
}

func changeValue(_value *int) {
	*_value = (*_value) * (*_value)
}

func add(numbers ...int) {
	var summ = 0

	for _, value := range numbers {
		summ += value
	}

	fmt.Println(summ)
}

func f(a int) int {
	if a == 0 {
		panic("BLYYYYYY")
	} else {
		return a
	}
}

func return_sum(numbers ...int) (int, string) {
	var sum = 0
	for _, value := range numbers {
		sum += value
	}
	var msg string = "ebat"
	return sum, msg
}
