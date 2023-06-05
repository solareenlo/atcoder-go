package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 2; i < n; i++ {
		if zs(i) {
			sum++
		}
	}
	fmt.Println(sum)
}

func zs(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
