package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)
	if n >= s {
		fmt.Println("Congratulations!")
	} else {
		fmt.Println("Enjoy another semester...")
	}
}
