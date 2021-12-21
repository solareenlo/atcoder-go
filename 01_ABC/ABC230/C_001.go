package main

import "fmt"

func main() {
	var n, a, b, p, q, r, s int
	fmt.Scan(&n, &a, &b, &p, &q, &r, &s)

	for i := p; i < q+1; i++ {
		for j := r; j < s+1; j++ {
			if i-j == a-b || i+j == a+b {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
