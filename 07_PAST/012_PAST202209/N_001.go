package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w, q int
	fmt.Fscan(in, &h, &w, &q)
	mn := min(h, w)
	mx := max(h, w)
	A := make([][]string, mn)
	for i := 0; i < h; i++ {
		var a string
		fmt.Fscan(in, &a)
		for j := 0; j < w; j++ {
			if h > w {
				A[j] = append(A[j], string(a[j]))
			} else {
				A[i] = append(A[i], string(a[j]))
			}
		}
	}
	for i := 0; i < mn; i++ {
		reverseOrderString(A[i])
	}
	reverseOrderStringSlice(A)
	s := make([]string, 0)
	for q > 0 {
		q--
		var t, p int
		var c string
		fmt.Fscan(in, &t, &p, &c)
		p--
		if h > w {
			t = 3 - t
		}
		if t == 1 {
			p = mn - 1 - p
			s = append(s, A[p][len(A[p])-mx])
			A[p] = append(A[p], c)
		} else {
			p = mx - 1 - p
			s = append(s, A[0][len(A[0])-mx+p])
			for i := 0; i < mn-1; i++ {
				A[i][len(A[i])-mx+p] = A[i+1][len(A[i+1])-mx+p]
			}
			A[mn-1][len(A[mn-1])-mx+p] = c
		}
	}
	fmt.Fprintln(out, strings.Join(s, ""))
}

func reverseOrderString(a []string) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func reverseOrderStringSlice(a [][]string) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
