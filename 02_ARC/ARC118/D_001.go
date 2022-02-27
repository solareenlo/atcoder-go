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

	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)

	vis := make([]bool, n)
	va := make([]int, 0)
	vb := make([]int, 0)
	for i := 1; !vis[i]; i = i * a % n {
		vis[i] = true
		va = append(va, i)
	}
	vis[1] = false

	for i := 1; !vis[i]; i = i * b % n {
		vis[i] = true
		vb = append(vb, i)
	}
	if len(va)*len(vb) < n-1 {
		fmt.Fprintln(out, "No")
		return
	}

	fmt.Fprintln(out, "Yes")
	for _, x := range vb {
		fmt.Fprint(out, x, " ")
	}
	for i := len(vb) - 1; i >= 0; i-- {
		if (len(vb)-i)&1 != 0 {
			for j := len(va) - 1; j > 0; j-- {
				fmt.Fprint(out, vb[i]*va[j]%n, " ")
			}
		} else {
			for j := 1; j < len(va); j++ {
				fmt.Fprint(out, vb[i]*va[j]%n, " ")
			}
		}
	}
	fmt.Fprintln(out, 1)
}
