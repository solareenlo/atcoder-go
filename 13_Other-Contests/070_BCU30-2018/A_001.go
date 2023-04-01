package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	AP := 1
	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)
		AP *= A
	}

	var M int
	fmt.Scan(&M)
	BP := 1
	for i := 0; i < M; i++ {
		var B int
		fmt.Scan(&B)
		BP *= B
	}

	if AP == BP {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
