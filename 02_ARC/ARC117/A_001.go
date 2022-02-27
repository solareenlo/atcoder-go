package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	sum := 0
	for i := 1; i < a; i++ {
		fmt.Print(i, " ")
		sum += i
	}
	for i := 1; i < b; i++ {
		fmt.Print(-i, " ")
		sum -= i
	}

	if a > b {
		fmt.Println(a, -sum-a)
	} else if a < b {
		fmt.Println(-b, -sum+b)
	} else {
		fmt.Println(a, -b)
	}
}
