package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200200

	var b1, b2 [N]int
	var sb [N][]int
	var vis [N]bool

	var n int
	fmt.Fscan(in, &n)
	x := make([]int, n+1)
	y := make([]int, n+1)
	v := make([]int, n+1)
	p1 := make([]int, n+1)
	p2 := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i], &v[i])
		p1[i] = x[i]
		p2[i] = y[i]
	}
	sort.Ints(p1[1:])
	sort.Ints(p2[1:])
	p1 = unique(p1[1:])
	cnt1 := len(p1)
	p2 = unique(p2[1:])
	cnt2 := len(p2)
	for i := 1; i <= n; i++ {
		x[i] = lowerBound(p1, x[i]) + 1
		y[i] = lowerBound(p2, y[i]) + 1
		b1[x[i]] += v[i]
		b2[y[i]] += v[i]
		sb[x[i]] = append(sb[x[i]], i)
	}
	id := make([]int, cnt2+1)
	for i := 1; i <= cnt2; i++ {
		id[i] = i
	}
	tmp := id[1:]
	sort.Slice(tmp, func(i, j int) bool {
		return b2[tmp[i]] > b2[tmp[j]]
	})
	ans := 0
	for i := 1; i <= cnt1; i++ {
		tmp := 0
		for _, j := range sb[i] {
			b2[y[j]] -= v[j]
			vis[y[j]] = true
		}
		now := 1
		for now <= cnt2 {
			tmp = max(tmp, b2[id[now]])
			if !vis[id[now]] {
				break
			}
			now++
		}
		ans = max(ans, tmp+b1[i])
		for _, j := range sb[i] {
			b2[y[j]] += v[j]
			vis[y[j]] = false
		}
	}
	fmt.Println(ans)
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
