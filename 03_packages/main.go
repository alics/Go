package main

import (
	"fmt"
	"math"

	"github.com/alics/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.34))
	fmt.Println(math.Ceil(2.34))
	fmt.Println(math.Sqrt(36))
	fmt.Println(strutil.Reverse("ali faghani hamedani"))

}
