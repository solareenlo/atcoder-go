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

	ans := 4
	utpc := "UTPC"
	for i := 0; i < n-3; i++ {
		t := 4
		stk := make(map[byte]struct{})
		for j := 0; j < 4; j++ {
			stk[s[i+j]] = struct{}{}
		}
		if s[i] == 'U' {
			t--
		}
		if s[i+1] == 'T' {
			t--
		}
		if s[i+2] == 'P' {
			t--
		}
		if s[i+3] == 'C' {
			t--
		}
		for j := 0; j < 3; j++ {
			for k := j + 1; k < 4; k++ {
				if s[i+j] == utpc[k] && s[i+k] == utpc[j] {
					t--
				}
			}
		}
		if len(stk) == 4 && t == 4 {
			t = 3
		}
		for j := 0; j < 2; j++ {
			for k := j + 1; k < 4; k++ {
				for l := j + 1; l < 4; l++ {
					if k == l {
						continue
					}
					if s[i+j] == utpc[k] && s[i+k] == utpc[l] && s[i+l] == utpc[j] {
						t--
					}
				}
			}
		}
		ans = min(t, ans)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
