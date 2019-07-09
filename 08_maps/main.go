package main

import "fmt"

func main() {

	emails := make(map[string]string)

	emails["a"] = "a@gmail.com"
	emails["b"] = "b@gmail.com"
	emails["c"] = "c@gmail.com"
	emails["d"] = "d@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["b"])

	delete(emails, "c")

	fmt.Println(emails)

	emails1 := map[string]string{"a1": "a1@gmail.com", "b1": "b1@gmail.com"}

	fmt.Println(emails1)
}
