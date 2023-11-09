package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const p = 998244353
	const N = 200005

	var v, e, d [N]int

	var n int
	fmt.Fscan(in, &n)
	v[1] = 1
	e[1] = 1
	d[1] = 1
	i := 2
	s := 1
	for i <= n {
		v[i] = (p - p/i) * v[p%i] % p
		e[i] = e[i-1] + v[i]
		var j int
		fmt.Fscan(in, &j)
		d[i] = d[j] + 1
		s += e[d[i]] % p
		i++
	}
	fmt.Println(s % p)
}
