package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	if K > N/2 {
		fmt.Println(-1)
		return
	}

	odd := 0
	if N%2 != 0 {
		odd = 1
		N--
	}

	for i := 0; i < K; i++ {
		if i != 0 {
			fmt.Println()
		}
		fmt.Printf("%d %d\n", i+1, i+N/2+1)
		for j := 0; j < N/2-1; j++ {
			fmt.Printf("%d %d\n", i+1, (i+j+1)%N+1)
			fmt.Printf("%d %d\n", i+N/2+1, (i+N/2+j+1)%N+1)
		}
		if odd != 0 {
			fmt.Printf("%d %d\n", i+1, N+1)
		}
	}
}
