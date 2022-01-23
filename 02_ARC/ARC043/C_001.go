package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	tree = [100005]int{}
)

func add(x int) {
	for ; x > 0; x -= x & -x {
		tree[x]++
	}
}

func query(x int) int {
	res := 0
	for ; x <= 100000; x += x & -x {
		res += tree[x]
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	x := make([]int, n+1)
	a := make([]int, 100005)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
		a[x[i]] = i
	}

	b := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		var t int
		fmt.Fscan(in, &t)
		b[i] = a[t]
	}

	s := make([]int, 100005)
	for i := 1; i < n+1; i++ {
		s[i] = query(b[i])
		add(b[i])
	}

	sum := 0
	for i := 1; i < n+1; i++ {
		sum += s[i]
	}

	if sum&1 != 0 {
		fmt.Fprintln(out, -1)
		return
	}

	sum >>= 1
	p := 0
	for i := 1; i < n+1; i++ {
		if sum < s[i] {
			p = i
			break
		}
		sum -= s[i]
	}

	sort.Slice(b[:p], func(i, j int) bool {
		return b[i] < b[j]
	})

	for i := 1; i < sum+1; i++ {
		b[p], b[p-1] = b[p-1], b[p]
		p--
	}

	for i := 1; i < n+1; i++ {
		fmt.Fprint(out, x[b[i]], " ")
	}
}
