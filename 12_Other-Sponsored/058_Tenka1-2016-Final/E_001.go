package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, s, x [10009]int

	var N, L int
	fmt.Fscan(in, &N, &L)
	for i := 0; i < N; i++ {
		c := 0
		for j := 0; j < L; j++ {
			fmt.Fscan(in, &a[j])
			for c >= 2 && (a[s[c-1]]-a[s[c-2]]+(s[c-1]-s[c-2])*(s[c-1]-s[c-2]))*(j-s[c-2]) >= (a[j]-a[s[c-2]]+(j-s[c-2])*(j-s[c-2]))*(s[c-1]-s[c-2]) {
				c--
			}
			s[c] = j
			c++
		}
		ptr := 0
		for j := 0; j < L; j++ {
			for ptr != c-1 && a[s[ptr]]+(s[ptr]-j)*(s[ptr]-j) >= a[s[ptr+1]]+(s[ptr+1]-j)*(s[ptr+1]-j) {
				ptr++
			}
			x[j] += a[s[ptr]] + (s[ptr]-j)*(s[ptr]-j)
		}
	}
	ans := x[0]
	for i := 0; i < L; i++ {
		ans = min(ans, x[i])
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
