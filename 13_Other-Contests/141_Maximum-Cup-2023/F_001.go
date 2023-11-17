package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, M, K, S int
	fmt.Scan(&N, &M, &K, &S)
	Len := N + M
	for i := 0; i <= Len%K; i++ {
		x := i
		y := S - i
		if y < 0 || y > K-Len%K {
			continue
		}
		cnt := i + (Len/K)*S
		if cnt == M {
			t := make([]string, K)
			for j := range t {
				t[j] = "0"
			}
			for j := 0; j < x; j++ {
				t[j] = "1"
			}
			for j := 0; j < y; j++ {
				t[Len%K+j] = "1"
			}
			fmt.Println("Yes")
			for j := 0; j < Len; j += K {
				fmt.Print(strings.Join(t[:min(K, Len-j)], ""))
			}
			fmt.Println()
			return
		}
	}
	fmt.Println("No")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
