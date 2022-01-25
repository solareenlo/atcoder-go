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

	var n, k int
	fmt.Fscan(in, &n, &k)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = i + 1
	}

	x := 1
	for i := 0; i < n; i++ {
		y := (x*(n-i) - 1) / k
		fmt.Fprintln(out, v[y])
		v = erase(v, y)
		x = (x*(n-i)-1)%k + 1
	}
}

func erase(a []int, pos int) []int {
	return append(a[:pos], a[pos+1:]...)
}
