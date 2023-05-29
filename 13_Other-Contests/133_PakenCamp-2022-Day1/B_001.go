package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	if N == 1 {
		fmt.Println(0)
	} else if M == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println((N-2)/(M-1) + 1)
	}
}
