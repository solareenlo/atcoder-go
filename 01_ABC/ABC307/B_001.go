package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	A := make([]string, n)
	B := make([]string, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		A[i] = s
		B[i] = reverseString(s)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				if A[i]+A[j] == B[j]+B[i] {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
