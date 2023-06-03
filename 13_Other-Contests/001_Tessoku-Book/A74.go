package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	R := make([]int, n)
	C := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var p int
			fmt.Fscan(in, &p)
			if p > 0 {
				C[j] = p
				R[i] = p
			}
		}
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if C[i] > C[j] {
				ans++
			}
			if R[i] > R[j] {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
