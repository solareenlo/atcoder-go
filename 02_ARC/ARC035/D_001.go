package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	fact = [3000000]float64{}
	bit  = [3000000]float64{}
	p    = [300000]int{}
	q    = [300000]int{}
)

func add(k int, x float64) {
	k++
	for k < 200000 {
		bit[k] += x
		k += k & -k
	}
}

func sum(k int) float64 {
	k++
	res := 0.0
	for k > 0 {
		res += bit[k]
		k -= k & -k
	}
	return res
}

func add_edge(i int) {
	x := p[i+1] - p[i]
	y := q[i+1] - q[i]
	add(i, fact[x+y]-fact[x]-fact[y])
}

func remove_edge(i int) {
	x := p[i+1] - p[i]
	y := q[i+1] - q[i]
	add(i, -fact[x+y]+fact[x]+fact[y])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fact[0] = math.Log2(1)
	fact[1] = math.Log2(1)
	for i := 2; i < 3000000; i++ {
		fact[i] = fact[i-1] + math.Log2(float64(i))
	}

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i], &q[i])
	}
	for i := 0; i < n-1; i++ {
		add_edge(i)
	}

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var k, a, b int
			fmt.Fscan(in, &k, &a, &b)
			k--
			if k != 0 {
				remove_edge(k - 1)
			}
			if k < n-1 {
				remove_edge(k)
			}
			p[k] = a
			q[k] = b
			if k != 0 {
				add_edge(k - 1)
			}
			if k < n-1 {
				add_edge(k)
			}
		} else {
			var l1, r1, l2, r2 int
			fmt.Fscan(in, &l1, &r1, &l2, &r2)
			l1 -= 2
			r1 -= 2
			l2 -= 2
			r2 -= 2
			if sum(r1)-sum(l1) > sum(r2)-sum(l2) {
				fmt.Fprintln(out, "FIRST")
			} else {
				fmt.Fprintln(out, "SECOND")
			}
		}
	}
}
