package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := 0
	for i := 1; i*i+i < n; i++ {
		if n%i == 0 {
			x += n/i - 1
		}
	}
	fmt.Println(x)
}
