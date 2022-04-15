package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	a := make([]int, k)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		b[i] = -1
	}

	v := make([]int, 0)
	was := make([]bool, 100100)
	for i := k - 1; i >= 0; i-- {
		if !was[a[i]] {
			was[a[i]] = true
			b[a[i]] = len(v)
			v = append(v, a[i])
		}
	}

	var i int
	for i = 1; i <= m; i++ {
		if !was[i] {
			break
		}
	}

	for m = len(v); m > 0 && v[m-1] < i; m-- {
		i = v[m-1]
	}

	f := make([]int, 100100)
	for i := k - 1; i >= 0; i-- {
		j := b[a[i]]
		var tmp int
		if j > 0 {
			tmp = min(f[j]+1, f[j-1])
		} else {
			tmp = (f[j] + 1)
		}
		f[j] = min(n, tmp)
	}
	if m == 0 || f[m-1] >= n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
