package main

import (
	"bufio"
	"fmt"
	"os"
)

var f [200200]int

func fnd(x int) int {
	if f[x] == x {
		return x
	}
	f[x] = fnd(f[x])
	return f[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type P struct {
		x, y int
	}
	mp := make(map[P]int)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		f[i] = i
	}
	for m > 0 {
		m--
		var x, y int
		fmt.Fscan(in, &x, &y)
		f[fnd(x)] = fnd(y)
	}
	var k int
	fmt.Fscan(in, &k)
	for k > 0 {
		k--
		var x, y int
		fmt.Fscan(in, &x, &y)
		mp[P{fnd(x), fnd(y)}] = 1
	}
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		_, ok0 := mp[P{fnd(x), fnd(y)}]
		_, ok1 := mp[P{fnd(y), fnd(x)}]
		if !ok0 && !ok1 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
