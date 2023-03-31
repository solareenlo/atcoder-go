package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const P = 900001
	const C = 950001

	var n, m int
	fmt.Scan(&n, &m)
	ma := make(map[int]int)
	for i := 1; i <= 412222; i++ {
		ma[(i*i)%P] = i
	}
	for i := 1; i <= n; i++ {
		a := (i * i) % P
		b := i % P
		fmt.Fprint(out, a*C+b)
		out.Flush()
		if i+1 <= n {
			fmt.Fprint(out, " ")
			out.Flush()
		}
	}
	fmt.Fprintln(out)
	out.Flush()
	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Scan(&x)
		sos := x / C
		sum := x % C
		prod := ((sum*sum - sos) * ((P + 1) / 2)) % P
		diff := ma[((sum*sum-4*prod)%P+3*P)%P]
		a := ((sum + diff) * ((P + 1) / 2)) % P
		b := ((sum - diff) * ((P + 1) / 2)) % P
		if a > b {
			a, b = b, a
		}
		fmt.Fprintln(out, a, b)
		out.Flush()
	}
}
