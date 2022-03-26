package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 3 {
		fmt.Println(2, 5, 63)
		return
	}

	a := n - 4
	if n%3 != 0 {
		a = n - 2
	}

	if a > 15000 {
		a = 15000
		if n%2 != 0 {
			a = 14999
		}
	}

	for i := 1; i <= a; i++ {
		fmt.Print(i*2, " ")
	}

	for i := 0; i < n-a; i++ {
		fmt.Print(i*6+3, " ")
	}
}
