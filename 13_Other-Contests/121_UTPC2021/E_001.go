package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type A struct {
	x, y, c int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n, k int
	fmt.Fscan(in, &n, &k)
	v := make([]A, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i].x, &v[i].y, &v[i].c)
	}
	if n < 16 {
		for i := 0; i < 16; i++ {
			v = append(v, A{0, 0, -INF})
		}
		n += 16
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].c > v[j].c
	})
	sum := 0
	xmax := 0
	ymax := 0
	xmin := INF
	ymin := INF
	cnt := 0
	for i := 0; i < k-4; i++ {
		sum += v[i].c
		xmax = max(xmax, v[i].x)
		ymax = max(ymax, v[i].y)
		xmin = min(xmin, v[i].x)
		ymin = min(ymin, v[i].y)
		cnt++
	}
	rem := make([]A, 0)
	for i := cnt; i < k; i++ {
		rem = append(rem, v[i])
	}
	v = reverseOrderA(v)
	resize(&v, max(0, n-k))
	if len(v) <= 9 {
		for _, p := range v {
			rem = append(rem, p)
		}
	} else {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				sort.Slice(v, func(a, b int) bool {
					return i*v[a].x+j*v[a].y+v[a].c < i*v[b].x+j*v[b].y+v[b].c
				})
				rem = append(rem, v[len(v)-1])
				v = v[:len(v)-1]
			}
		}
	}
	id := make([]int, len(rem))
	for i := 0; i < len(id); i++ {
		if i < k-cnt {
			id[i] = 1
		} else {
			id[i] = 0
		}
	}
	id = reverseOrderInt(id)

	ans := 0
	S := sum
	xxmax := xmax
	yymax := ymax
	xxmin := xmin
	yymin := ymin
	for i := 0; i < len(id); i++ {
		if id[i] != 0 {
			S += rem[i].c
			xxmax = max(xxmax, rem[i].x)
			xxmin = min(xxmin, rem[i].x)
			yymax = max(yymax, rem[i].y)
			yymin = min(yymin, rem[i].y)
		}
	}
	ans = max(ans, xxmax-xxmin+yymax-yymin+S)
	for nextPermutation(sort.IntSlice(id)) {
		S := sum
		xxmax := xmax
		yymax := ymax
		xxmin := xmin
		yymin := ymin
		for i := 0; i < len(id); i++ {
			if id[i] != 0 {
				S += rem[i].c
				xxmax = max(xxmax, rem[i].x)
				xxmin = min(xxmin, rem[i].x)
				yymax = max(yymax, rem[i].y)
				yymin = min(yymin, rem[i].y)
			}
		}
		ans = max(ans, xxmax-xxmin+yymax-yymin+S)
	}
	fmt.Println(ans)
}

func reverseOrderA(a []A) []A {
	n := len(a)
	res := make([]A, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func resize(a *[]A, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, A{0, 0, 0})
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
