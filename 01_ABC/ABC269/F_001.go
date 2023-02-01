package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353

var m int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &m, &q)
	for i := 1; i <= q; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		fmt.Fprintln(out, (query(a-1, c-1)+query(b, d)+(mod<<1)-query(a-1, d)-query(b, c-1))%mod)
	}
}

func query(a, b int) int {
	c := (b + 1) >> 1
	d := b >> 1
	e := (a + 1) >> 1
	f := a >> 1
	l1 := c * c % mod
	l2 := d * (d + 1) % mod
	return (l1*e + l2*f + m*c%mod*e%mod*(e-1)%mod + m*d%mod*f%mod*f%mod) % mod
}
