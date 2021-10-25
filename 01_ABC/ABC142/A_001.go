package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Printf("%.6f\n", 1.0)
	} else if n%2 != 0 {
		fmt.Printf("%.6f\n", float64(n/2+1)/float64(n))
	} else {
		fmt.Printf("%.6f\n", 0.5)
	}
}
