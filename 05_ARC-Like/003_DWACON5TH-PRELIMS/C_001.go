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

	var n, q int
	var s string
	fmt.Fscan(in, &n, &s, &q)

	D := make([]int, 0)
	C := make([]int, 0)
	t := [3][1000001]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			t[j][i+1] = t[j][i]
		}
		switch s[i] {
		case 'D':
			t[0][i+1]++
			D = append(D, i+1)
			break
		case 'M':
			t[1][i+1]++
			break
		case 'C':
			t[2][i+1]++
			C = append(C, i)
			break
		}
	}

	suma := 0
	sumc := 0
	for _, a := range D {
		suma += t[1][a] * t[2][a]
	}
	for _, c := range C {
		sumc += t[1][c] * t[0][c]
	}

	for i := 0; i < q; i++ {
		var k int
		fmt.Fscan(in, &k)
		ans := suma + sumc
		for _, a := range D {
			if a+k-1 < n {
				ans -= t[1][a] * t[2][a+k-1]
			} else {
				ans -= t[1][a] * t[2][n]
			}
		}
		for _, c := range C {
			if c-k+1 > 0 {
				ans -= t[1][c] * t[0][c-k+1]
			} else {
				ans -= t[1][c] * t[0][0]
			}
		}
		fmt.Fprintln(out, ans)
	}
}
