package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)
	a, b := 0, 0
	ans := 0
	for i := 0; i < N; i++ {
		if S[i] == '0' {
			a = a + b
			b = 1
		} else {
			a, b = b, a
			a++
		}
		ans += a
	}
	fmt.Println(ans)
}
