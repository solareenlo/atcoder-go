package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for n > 0 {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				n /= i
				i--
			}
		}
		fmt.Println(n)
		fmt.Scan(&n)
	}
}
