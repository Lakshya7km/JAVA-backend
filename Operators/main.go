package main

import "fmt"

//arithmetic operators
//logical operators
//comparison operators
//assignment operators
//bitwise operators

func main() {

	//arithmetic operators
	//+, -, *, /, %, ++, --

	// a:=34901
	// b:=12
	// fmt.Println("addition",a+b)
	// fmt.Println("subtraction",a-b)
	// fmt.Println("multiplication",a*b)
	// fmt.Println("division",a/b)
	// fmt.Println("modulus",a%b)

	// //i cannot preincrement and predecrement in go, only post increment and post decrement is allowed
	// a++
	// fmt.Println("post increment",a)
	// a--
	// fmt.Println("post decrement",a)

	//2. Assignment operators
	// =, +=, -=, *=, /=, %=
	// a:= 10
	// b:=34
	// a+=b
	// fmt.Println(a)

	// //3.comparison operators
	// // ==, !=, >, <, >=, <=
	// fmt.Println(a==b)
	// fmt.Println(a>=b)
	// fmt.Println(a<=b)
	// fmt.Println(a!=b)

	// 4. Logical Operator
	// a:=true
	// b:=false
	// // &&, ||, !
	// fmt.Println(a && b)
	// fmt.Println(a || b)
	// fmt.Println(!a)



	//5.Bitwise Operators
	// &, |, ^, <<, >>,&^
	//perform operations on bits and perform bit by bit operations
     

	a:=5
	b:=3
	//01011
	//00011
	fmt.Println("Bitwise AND",a&b)
	fmt.Println("Bitwise OR",a|b)
	fmt.Println("Bitwise XOR",a^b)
	fmt.Println("Left Shift",a<<1)
	fmt.Println("Right Shift",a>>1)
	fmt.Println("Bit Clear",a&^b)


	c:=8
	//1000
	fmt.Println("leftshift by 1",c<<1)//
	fmt.Println("leftshift by 2",c<<2)
	fmt.Println("leftshift by 3",c<<3)
   fmt.Println("rightshift by 1",c>>1)
   fmt.Println("rightshift by 2",c>>2)
   fmt.Println("rightshift by 3",c>>3)
}
