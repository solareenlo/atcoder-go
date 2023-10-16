package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	if N%2 == 1 {
		fmt.Println(-1)
		return
	}

	ans := 0
	for i := 0; i < N/2; i++ {
		if S[i] != S[N/2+i] {
			ans++
		}
	}
	fmt.Println(ans)
}
