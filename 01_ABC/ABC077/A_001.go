package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	b = reverse(b)

	if a == b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverse(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
