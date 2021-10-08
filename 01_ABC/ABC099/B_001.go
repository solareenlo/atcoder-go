package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	sum := 0
	for i := 0; i < b-a; i++ {
		sum += i
	}
	fmt.Println(sum - a)
}
