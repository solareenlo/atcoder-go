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

	var n, m int
	fmt.Fscan(in, &n, &m)

	p := [10]int{}
	p[1] = 1
	p[2] = m
	a := make([]int, m+2)
	t := 0
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
		if a[i]&1 != 0 {
			t++
			p[t] = i
		}
		if t > 2 {
			fmt.Fprintln(out, "Impossible")
			return
		}
	}

	a[1], a[p[1]] = a[p[1]], a[1]
	a[m], a[p[2]] = a[p[2]], a[m]

	for i := 1; i <= m; i++ {
		fmt.Fprint(out, a[i], " ")
	}

	fmt.Fprintln(out)
	if m == 1 {
		m++
		a[m], a[1] = a[1], a[m]
	}

	a[m]--
	a[1]++

	if a[m] == 0 {
		m--
	}

	fmt.Fprintln(out, m)
	for i := 1; i <= m; i++ {
		if a[i] != 0 {
			fmt.Fprint(out, a[i], " ")
		}
	}
}
