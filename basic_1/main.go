package main

import (
	"fmt"
	"time"
)

const Pi = 3.14

func main() {

	//constant access
	fmt.Println("Pi Value is : ", Pi)

	//for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum of first 10 numbers is : ", sum)

	//while loop - omit init and post statemetns to resemble While loop
	n := 1
	for n < 4 {
		n += n
	}
	fmt.Println("n value is: ", n)

	//range loop
	animals := []string{"dog", "cat"}
	for _, animal := range animals {
		fmt.Println("animal", animal)
	}

	//if else
	name := "test"
	if name == "test" {
		fmt.Println("true")
	} else {
		fmt.Println("fasle")
	}

	//switch - case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday and day is :", time.Now().Weekday())
	}
	//defer functions
	fmt.Println("start")
	for i := 0; i < 5; i++ {
		//remove defer and see sequential execution
		//defer executes in LIFO order
		defer fmt.Println("Index is : ", i)
	}
	fmt.Println("stop")

}
