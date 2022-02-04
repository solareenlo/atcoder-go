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

	var n, D int
	fmt.Fscan(in, &n, &D)

	const N = 500005
	b := make([]int, N)
	a := make([]int, N)
	y := make([]int, N)
	b[0] = D
	y[n+1] = 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = min(b[i-1], abs(b[i-1]-a[i]))
	}

	for i := n; i > 0; i-- {
		if y[i+1] <= a[i]/2 {
			y[i] = y[i+1]
		} else {
			y[i] = y[i+1] + a[i]
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var q int
		fmt.Fscan(in, &q)
		if b[q-1] >= y[q+1] {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
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

	var n, D int
	fmt.Fscan(in, &n, &D)

	const N = 500005
	b := make([]int, N)
	a := make([]int, N)
	y := make([]int, N)
	b[0] = D
	y[n+1] = 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = min(b[i-1], abs(b[i-1]-a[i]))
	}

	for i := n; i > 0; i-- {
		if y[i+1] <= a[i]/2 {
			y[i] = y[i+1]
		} else {
			y[i] = y[i+1] + a[i]
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var q int
		fmt.Fscan(in, &q)
		if b[q-1] >= y[q+1] {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
