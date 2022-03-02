package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	f := make([]int, 150)
	f[0] = 1
	f[1] = 1
	t := 1
	for f[t-1]+f[t] <= n {
		t++
		f[t] = f[t-1] + f[t-2]
	}

	m := n
	s := t
	for i := t; i >= 1; i-- {
		if m >= f[i] {
			m -= f[i]
			s++
		}
	}

	fmt.Fprintln(out, s)
	for i := t; i >= 1; i-- {
		if n >= f[i] {
			fmt.Fprintln(out, (i&1)+1)
			n -= f[i]
		}
		fmt.Fprintln(out, ((i-1)&1)+3)
	}
}
