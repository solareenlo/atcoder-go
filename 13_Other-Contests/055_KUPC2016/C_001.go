package main

import "fmt"

func main() {
	var q, a, b int
	fmt.Scan(&q)
	for ; q > 0; q-- {
		fmt.Scan(&a, &b)
		if a&1 == 1 {
			fmt.Println((a-1)*127 + b)
		} else {
			fmt.Println((a-1)*127 + (127 ^ b))
		}
	}
}
