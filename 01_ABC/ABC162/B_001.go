package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 1; i < n+1; i++ {
		if i%3 != 0 && i%5 != 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
