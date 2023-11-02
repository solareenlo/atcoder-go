package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	sa := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sa += a[i]
	}
	b := make([]int, m+1)
	sb := 0
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
		sb += b[i]
	}
	sort.Ints(a[1:])
	sort.Ints(b[1:])
	ans := (n&1)*sb + (m&1)*sa
	n = n / 2 * 2
	m = m / 2 * 2
	nn := n / 2
	mm := m / 2
	lb, rb, la, ra := 0, 0, 0, 0
	for n != 0 || m != 0 {
		if a[n] > b[m] {
			if lb <= rb && la < nn || ra == nn {
				la++
				ans += 2 * lb * a[n]
			} else {
				ra++
				ans += 2 * rb * a[n]
			}
			n--
		} else {
			if la <= ra && lb < mm || rb == mm {
				lb++
				ans += 2 * la * b[m]
			} else {
				rb++
				ans += 2 * ra * b[m]
			}
			m--
		}
	}
	fmt.Println(ans)
}
