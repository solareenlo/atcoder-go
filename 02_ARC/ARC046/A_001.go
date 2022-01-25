package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	div := n / 9
	rem := n % 9
	if rem != 0 {
		div++
	} else {
		rem = 9
	}
	for i := 0; i < div; i++ {
		fmt.Print(rem)
	}
	fmt.Println()
}
