package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)

	if x+a-b >= 0 {
		if a-b >= 0 {
			fmt.Println("delicious")
		} else {
			fmt.Println("safe")
		}
	} else {
		fmt.Println("dangerous")
	}
}
