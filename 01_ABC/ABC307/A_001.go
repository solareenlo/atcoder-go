package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		ans := 0
		for j := 0; j < 7; j++ {
			var A int
			fmt.Scan(&A)
			ans += A
		}
		fmt.Printf("%d ", ans)
	}
	fmt.Println()
}
