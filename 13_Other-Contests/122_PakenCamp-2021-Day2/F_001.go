package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N == 2 {
		fmt.Println("1 2 1 1 2 1")
	} else if N == 3 {
		fmt.Println("1 3 1 2 2 3")
	} else {
		fmt.Println(N-3, N-2, N-2, N-1, N-1, N)
	}
}
