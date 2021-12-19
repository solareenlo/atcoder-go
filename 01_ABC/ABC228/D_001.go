package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1 << 20

var (
	a  = [N + 10]int{}
	nx = [N + 10]int{}
)

func find(w int) int {
	if a[w] == -1 {
		return w
	} else {
		nx[w] = find(nx[w])
		return nx[w]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := range a {
		a[i] = -1
	}
	for i := range nx {
		nx[i] = (i + 1) % N
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var t, x int
		fmt.Fscan(in, &t, &x)
		if t == 1 {
			a[find(x%N)] = x
		} else {
			fmt.Fprintln(out, a[x%N])
		}
	}
}
