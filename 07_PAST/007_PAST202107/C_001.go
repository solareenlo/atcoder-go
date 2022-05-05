package main

import "fmt"

func main() {
	var S string
	var L, R int
	fmt.Scan(&S, &L, &R)

	N := len(S)
	if N == 1 || (S[0] != '0' && N <= 10) {
		n := 0
		for i := 0; i < N; i++ {
			n = n*10 + int(S[i]-'0')
		}
		if L <= n && n <= R {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
