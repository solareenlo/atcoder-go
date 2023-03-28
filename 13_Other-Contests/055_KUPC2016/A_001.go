package main

import (
	"fmt"
)

func main() {
	var N, A, B int64
	fmt.Scan(&N, &A, &B)
	ans := 0
	for i := int64(0); i < N; i++ {
		var t int64
		fmt.Scan(&t)
		if !(A <= t && t < B) {
			ans++
		}
	}
	fmt.Println(ans)
}
