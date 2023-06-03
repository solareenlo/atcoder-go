package main

import "fmt"

func main() {
	var N string
	fmt.Scan(&N)
	ans := 0
	Len := len(N)
	for i := 0; i < Len; i++ {
		if N[Len-i-1] == '1' {
			ans += pow(2, i)
		}
	}
	fmt.Println(ans)
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
