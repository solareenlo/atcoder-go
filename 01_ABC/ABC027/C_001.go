package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	i := int64(1)
	n++
	for ; n > 1; i++ {
		n = (n + i%2) / 2
	}
	if i%2 == 1 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
