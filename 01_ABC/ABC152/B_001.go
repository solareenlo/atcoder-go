package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	A := byte(a) + '0'
	B := byte(b) + '0'

	strA, strB := "", ""
	for i := 0; i < b; i++ {
		strA += string(A)
	}
	for i := 0; i < a; i++ {
		strB += string(B)
	}

	fmt.Println(min(strA, strB))
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}
