package main

import "fmt"

func main() {

	a := 6
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	fmt.Println(*b)
	fmt.Println(*&a)

	//change val with the pointer
	*b = 9
	fmt.Println(a)

}
