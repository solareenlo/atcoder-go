package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n+1; i++ {
		sum += i
	}

	fmt.Println(sum)
}
