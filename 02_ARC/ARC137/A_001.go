package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	for i := r - l; i >= 1; i-- {
		for j := r - i; j >= l; j-- {
			if gcd(j, i+j) == 1 {
				fmt.Println(i)
				return
			}
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
