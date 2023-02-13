package main

import "fmt"

func main() {
	a, b, c := true, false, true
	if a {
		fmt.Print("At")
	} else {
		fmt.Print("Yo")
	}
	if !a && b {
		fmt.Print("Bo")
	} else if !b || c {
		fmt.Print("Co")
	}

	if a && b && c {
		fmt.Print("foo!")
	} else if true && false {
		fmt.Print("yeah!")
	} else if !a || c {
		fmt.Print("der")
	}

	fmt.Println()
}
