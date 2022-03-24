package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	ans int
	cnt = make([]int, 200005)
)

func Add(x int) {
	if x > 0 {
		cnt[x]++
		if (cnt[x]) == 1 {
			ans++
		}
	}
}

func Del(x int) {
	if x > 0 {
		cnt[x]--
		if (cnt[x]) == 0 {
			ans--
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, 200005)
	c := make([]int, 200005)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		c[a[i]]++
		Add(a[i] - (c[a[i]]) + 1)
	}

	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)

		Del(a[x] - (c[a[x]]) + 1)
		c[a[x]]--
		c[y]++
		Add(y - (c[y]) + 1)
		a[x] = y
		fmt.Println(n - ans)
	}
}
