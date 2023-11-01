package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	S := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&S[i])
	}
	for i := 0; i < N; i++ {
		fmt.Println(S[N-i-1])
	}
}
