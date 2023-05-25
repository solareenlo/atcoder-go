package main

import "fmt"

func main() {
	var L, R int
	fmt.Scan(&L, &R)
	R++
	N := R - L
	ans := 0
	for i := 1; i < N; i++ {
		for X := (L + i - 1) / i * i; X+i < R; X += i {
			if X&i == 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
