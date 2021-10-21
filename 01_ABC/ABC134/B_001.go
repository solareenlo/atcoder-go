package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	if n >= 2*d+1 {
		if n%(2*d+1) != 0 {
			fmt.Println(n/(2*d+1) + 1)
		} else {
			fmt.Println(n / (2*d + 1))
		}
	} else {
		fmt.Println(1)
	}
}
