package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var n, a, b, u, x, y, v int
var p [60000 + 5]int

func main() {
	defer out.Flush()

	fmt.Scan(&n)

	var t1, t2, t3, t4 int
	ask(1, 2, 3, &t1)
	ask(1, 2, 4, &t2)
	ask(1, 3, 4, &t3)
	ask(2, 3, 4, &t4)
	if t1 == t2 {
		a, b, u, x, y, v = 1, 2, t1, 3, 4, t3
	} else if t1 == t3 {
		a, b, u, x, y, v = 1, 3, t1, 2, 4, t2
	} else {
		a, b, u, x, y, v = 2, 3, t1, 1, 4, t2
	}
	if u > v {
		a, x = x, a
		b, y = y, b
		u, v = v, u
	}

	for i := 5; i <= n; i++ {
		var j int
		ask(a, x, i, &j)
		if j > u && j < v {
			p[i] = j
			continue
		}
		if j <= u {
			work1(i, j)
		} else {
			work2(i, j)
		}
	}
	var t int
	ask2(a, b, &t)
	p[t] = 1
	p[a+b-t] = 2
	ask2(x, y, &t)
	p[t] = n - 1
	p[x+y-t] = n
	fmt.Fprintf(out, "!")
	out.Flush()
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, " %d", p[i])
		out.Flush()
	}
	fmt.Fprintln(out)
	out.Flush()
}

func ask(a, b, c int, v *int) {
	fmt.Fprintf(out, "? 1 %d %d %d\n", a, b, c)
	out.Flush()
	fmt.Scan(v)
}

func ask2(a, b int, v *int) {
	fmt.Fprintf(out, "? 2 %d %d\n", a, b)
	out.Flush()
	fmt.Scan(v)
}

func work1(i, j int) {
	var t int
	ask(b, i, x, &t)
	if u == j {
		p[a] = j
		a = i
		u = t
	} else if u == t {
		p[b] = u
		b = i
		u = j
	} else {
		p[i] = j
	}
}

func work2(i, j int) {
	var t int
	ask(y, i, a, &t)
	if v == j {
		p[x] = j
		x = i
		v = t
	} else if v == t {
		p[y] = v
		y = i
		v = j
	} else {
		p[i] = j
	}
}
