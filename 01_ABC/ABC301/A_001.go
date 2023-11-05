package main

import "fmt"

func main() {
	var n int
	var c string
	fmt.Scan(&n, &c)
	a, b := 0, 0
	for i := 0; i < n; i++ {
		if c[i] == 'A' {
			a++
		}
		if c[i] == 'T' {
			b++
		}
	}
	if b > a || (b == a && c[n-1] == 'A') {
		fmt.Println("T")
	} else {
		fmt.Println("A")
	}
}
