package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if N == 4 {
		fmt.Println(2, 3, 2, 1)
	} else if N == 5 {
		fmt.Println(3, 2, 3, 1, 1)
	} else if N < 7 {
		fmt.Println(-1)
	} else {
		fmt.Println(N-3, 3, 2)
		for i := 4; i < N-3; i++ {
			fmt.Printf(" %d", 1)
		}
		fmt.Println(" 2 1 1 1")
	}
}
