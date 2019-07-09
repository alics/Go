package main

import "fmt"

func main() {

	x := 10
	y := 15

	if x > y {
		fmt.Printf("%d is grater thay %d\n", x, y)
	} else {
		fmt.Printf("%d is less thay %d\n", x, y)
	}

	color := "red"

	if color == "red" {
		fmt.Println("the color is red")
	} else if color == "blue" {
		fmt.Println("the color is blue")
	} else {
		fmt.Println("the color is not blue and red")
	}

	switch color {
	case "red":
		fmt.Println("the color is red")
	case "blue":
		fmt.Println("the color is blue")
	default:
		fmt.Println("the color is not blue and red")
	}

}
