package main

import (
	"bufio"
	"fmt"
	"os"
)

var x = [2002]int{}
var b = []int{2, 3, 5, 7, 11, 13, 17}
var c = []int{19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293}
var res int

func main() {
	in := bufio.NewReader(os.Stdin)

	var q, m int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &m)
		fmt.Fscan(in, &x[m])
	}
	dfs(0, 1)
	fmt.Println(res)
}

func dfs(l, k int) {
	if l < 7 {
		for t := 1; t <= 300; {
			dfs(l+1, k*t)
			t *= b[l]
		}
		return
	}
	s := 0
	for i := 2; i <= 300; i++ {
		if k%i == 0 {
			s += x[i]
		}
	}
	for i := 0; i < 55; i++ {
		t := 0
		for j := 1; j*c[i] <= 300; j++ {
			if k%j == 0 {
				t += x[j*c[i]]
			}
		}
		if t > 0 {
			s += t
		}
	}
	res = max(res, s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
