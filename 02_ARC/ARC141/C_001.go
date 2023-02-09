package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 1000005

	var n int
	fmt.Fscan(in, &n)
	n = n * 2
	p := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	q := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &q[i])
	}
	s := make([]byte, N)
	for i := 1; i <= n; i += 2 {
		if p[i] > p[i+1] {
			if i != 1 && (p[i] < p[i-2] || p[i+1] < p[i-1]) {
				fmt.Println(-1)
				return
			}
			s[p[i]] = '('
			s[p[i+1]] = ')'
		}
	}
	for i := 1; i <= n; i += 2 {
		if q[i] < q[i+1] {
			if s[q[i]] != 0 || s[q[i+1]] != 0 || (i != 1 && (q[i] > q[i-2] || q[i+1] > q[i-1])) {
				fmt.Println(-1)
				return
			}
			s[q[i]] = '('
			s[q[i+1]] = ')'
		}
	}
	for i := 1; i <= n; i++ {
		if s[i] == 0 {
			fmt.Println(-1)
			return
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%c", s[i])
	}
}
