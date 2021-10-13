package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	if d == 25 {
		fmt.Println("Christmas")
	} else if d == 24 {
		fmt.Println("Christmas Eve")
	} else if d == 23 {
		fmt.Println("Christmas Eve Eve")
	} else if d == 22 {
		fmt.Println("Christmas Eve Eve Eve")
	}
}
