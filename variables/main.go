package main

import "fmt"

// func main() {
//entry point of the application
//types of declaring variables

// //1.standard way   var name type = value
// var name string = "John Doe"
// fmt.Println(name)

// //2. type inference   var name = value
// var age = 30
// fmt.Println(age)

// //3 multiple variable declaration var a,b,c int = 10,20,30
// var a,b,c =10,20,30
// fmt.Println(a,b,c)

// //4. short variable declaration   name := value
// //only allowed inside the scope of function
// city:= "New York"
// fmt.Println(city)

// }

func main() {
	// zero value of variable
	var name string    //""
	var age int        //0
	var salary float64 //0
	var isStudent bool //false
	fmt.Println(name,"_", age, salary, isStudent)


	//map ,slice default value = nil
}
