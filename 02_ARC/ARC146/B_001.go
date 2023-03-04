package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	for p := 30; p >= 0; p-- {
		cost := 0
		ans += 1 << p
		tmp := a[1:]
		sort.Ints(tmp)
		for i := n - k + 1; i <= n; i++ {
			if a[i] < ans {
				cost += ans - a[i]
			}
		}
		if cost > m {
			for i := n - k + 1; i <= n; i++ {
				if a[i] >= ans {
					a[i] -= 1 << p
				}
			}
			ans -= 1 << p
		}
	}
	fmt.Println(ans)
}
