package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	cnt1, cnt2 := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			t := 1
			for i-t >= 0 && s[i-t] == 'A' && i+t < n && s[i+t] == 'C' {
				t++
			}
			if t-1 != 0 {
				cnt1++
				cnt2 += t - 1
			}
		}
	}
	fmt.Println(min(2*cnt1, cnt2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
