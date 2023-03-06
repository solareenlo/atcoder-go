package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	st := make(map[int]int)
	var n int
	fmt.Fscan(in, &n)
	x := make([]int, 1000)
	k := make([]int, 1000)
	d := make([]int, 1000)
	xs := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &k[i], &d[i])
		st[x[i]]++
		xs = append(xs, x[i])
		xs = append(xs, x[i]-d[i])
		xs = append(xs, x[i]+d[i])
	}
	sort.Ints(xs)
	xs = unique(xs)
	N := len(xs)

	for i := 0; i < n; i++ {
		if k[i] == 1 && d[i] != 0 {
			fmt.Println("impossible")
			return
		}
	}

	var lookup func(int) int
	lookup = func(x int) int {
		return lowerBound(xs, x)
	}

	e := make([][]pair, 3001)
	for i := 0; i < n; i++ {
		t := lookup(x[i])
		e[t] = append(e[t], pair{t + 1, -st[x[i]]})
		if d[i] == 0 {
			e[t] = append(e[t], pair{t + 1, -k[i]})
		} else {
			lb := lookup(x[i] - d[i])
			ub := lookup(x[i] + d[i])
			e[ub] = append(e[ub], pair{lb + 1, k[i] - 1})
			e[lb] = append(e[lb], pair{ub + 1, -k[i]})
		}
	}
	for i := 0; i < N; i++ {
		e[i] = append(e[i], pair{i + 1, 0})
	}

	var dist []int

	var spfa func() bool
	spfa = func() bool {
		dist = make([]int, N+1)
		q := make([]int, 0)
		inq := make([]bool, N+1)
		for i := 0; i <= N; i++ {
			q = append(q, i)
			dist[i] = 0
			inq[i] = true
		}
		times := 0
		for len(q) > 0 {
			if times > 100000 {
				return false
			}
			times++
			u := q[0]
			q = q[1:]
			inq[u] = false
			for _, i := range e[u] {
				if dist[u]+i.y < dist[i.x] {
					dist[i.x] = dist[u] + i.y
					if !inq[i.x] {
						q = append(q, i.x)
						inq[i.x] = true
					}
				}
			}
		}
		return true
	}

	if !spfa() {
		fmt.Println("impossible")
		return
	}

	res := dist[0] - dist[N] - n
	if res > 100000 {
		fmt.Println("too many")
		return
	}

	ans := make([]int, 0)
	for i := 1; i <= N; i++ {
		t := dist[i-1] - dist[i]
		if _, ok := st[xs[i-1]]; ok {
			t -= st[xs[i-1]]
		}
		for j := 0; j < t; j++ {
			ans = append(ans, xs[i-1])
		}
	}

	fmt.Println(res)
	if res != 0 {
		for i := 0; i < res; i++ {
			fmt.Printf("%d", ans[i])
			if i+1 != res {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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
	sort.Ints(result)
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
