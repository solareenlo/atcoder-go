package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type num struct {
	id, p, e int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	const N = 200005
	var t int
	m := 0
	a := make([]num, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &t)
		for t > 0 {
			m++
			a[m].id = i
			fmt.Fscan(in, &a[m].p, &a[m].e)
			t--
		}
	}
	tmp := a[1 : m+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].p < tmp[j].p || tmp[i].p == tmp[j].p && tmp[i].e > tmp[j].e
	})

	flg := make([]int, N)
	for i := 1; i <= m; i++ {
		if a[i].p != a[i-1].p && (a[i].p != a[i+1].p || a[i].e > a[i+1].e) {
			flg[a[i].id] = 1
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += flg[i]
	}
	if ans < n {
		ans++
	}
	fmt.Println(ans)
}
