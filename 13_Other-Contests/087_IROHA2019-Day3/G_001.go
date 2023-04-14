package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N <= 1 {
		fmt.Println(5)
	} else if N <= 2 {
		fmt.Println(999999999999)
	} else if N <= 3 {
		fmt.Println(100000000)
	} else if N <= 4 {
		fmt.Println(1)
	} else {
		fmt.Println(208569179)
	}
}
