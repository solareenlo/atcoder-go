package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Println(i + 1)
	}
	fmt.Println(1)
}
