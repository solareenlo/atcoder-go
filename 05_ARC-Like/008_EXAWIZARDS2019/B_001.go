package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	a, b := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			a++
		} else {
			b++
		}
	}

	if a > b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
