package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	k := 0
	for n > 0 {
		n >>= 1
		k++
	}
	fmt.Println(k - 1)
}
