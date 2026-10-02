package main

import "fmt"

//constant using const keyword
//basic constant
//typed constant
//multiple constant

const Pi = 3.14    //basic constant
const age int = 32 //typed constant
const (
	active  = false
	salary  = 1000
	TimeOut = 300
)//multiple constant

func main() {

	// const age = 32
	fmt.Println("PI--", Pi)
	fmt.Println("age--", age)
	fmt.Println("active--", active)
	fmt.Println("salary--", salary)
	fmt.Println("TimeOut--", TimeOut)

}
