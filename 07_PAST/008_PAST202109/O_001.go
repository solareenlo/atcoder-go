package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	x := 1 << n
	q := make([]int, x)
	for i := 0; i < x; i++ {
		var p int
		fmt.Fscan(in, &p)
		p--
		q[p] = i
	}

	a := make([][2]int, x)
	for i := 0; i < x; i++ {
		a[i] = [2]int{-1, -1}
	}
	for i := 0; i < m; i++ {
		var w, l int
		fmt.Fscan(in, &w, &l)
		w--
		l--
		w = q[w]
		l = q[l]
		w1 := w + x
		l1 := l + x
		for {
			w1 /= 2
			l1 /= 2
			if w1 == l1 {
				if a[w1][1] >= 0 || (a[w1][0] >= 0 && a[w1][0] != w) {
					fmt.Println(0)
					return
				}
				a[w1] = [2]int{w, l}
				break
			} else {
				if a[w1][0] < 0 {
					a[w1][0] = w
				} else if a[w1][0] != w {
					fmt.Println(0)
					return
				}
				if a[l1][0] < 0 {
					a[l1][0] = l
				} else if a[l1][0] != l {
					fmt.Println(0)
					return
				}
			}
		}
	}

	k := x
	for i := 1; i < x; i++ {
		if a[i][0] < 0 {
			continue
		}
		k--
	}
	fmt.Println(powMod(2, k))
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
