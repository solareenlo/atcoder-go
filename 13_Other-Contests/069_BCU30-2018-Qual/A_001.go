package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var b [30]int
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		b[a]++
	}
	for i := 0; i < 29; i++ {
		fmt.Printf("%d ", b[i])
	}
	fmt.Printf("%d\n", b[29])
}
