package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N == 1 {
		fmt.Println(0)
	} else if N%2 == 1 {
		fmt.Println(-1)
	} else {
		for i := 0; i < N/2; i++ {
			fmt.Printf("%d %d ", (N-2*i)%N, 2*i+1)
		}
		fmt.Println()
	}
}
