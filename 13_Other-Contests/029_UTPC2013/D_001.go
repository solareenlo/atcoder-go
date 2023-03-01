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
	a := make([]int, 2222)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, 2222)
	for i := n; i <= 4*n; i++ {
		for j := range b {
			b[j] = int(1e18)
		}
		s := make(map[int]bool)
		flag := true
		for j := i; flag && j > 0; j-- {
			if i-j+1 <= n {
				if b[j] < a[n-i+j] {
					flag = false
				}
				b[j] = a[n-i+j]
			} else {
				p := b[j]
				for s[p] {
					p--
				}
				if p < 0 {
					flag = false
				}
				b[j] = p
			}
			s[b[j]] = true
			b[j>>1] = min(b[j>>1], b[j])
		}
		if flag {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
