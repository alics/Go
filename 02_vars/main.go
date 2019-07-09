package main

import "fmt"

var country = "Iran"

func main() {

	var name = "Ali"
	var age int32 = 35
	var isCool = true
	var size float32 = 2.3

	//Shorthand
	city := "Tehran"

	email, lname := "ali.faghani@gmail.com", "faghani"

	fmt.Println(name, age, isCool, city, country)
	fmt.Println(email, lname)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", size)
}
