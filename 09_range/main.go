package main

import "fmt"

func main() {

	ids := []int{23, 345, 34, 67, 689, 243, 23, 12}

	for i, id := range ids {
		fmt.Printf("%d - Id: %d \n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("Id: %d \n", id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	emails := map[string]string{"a1": "a1@gmail.com", "b1": "b1@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}

}
