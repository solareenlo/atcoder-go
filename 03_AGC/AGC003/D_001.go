package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const maxn = 100010
	cnt := 0
	f := [maxn + 10]bool{}
	p := [maxn + 10]int{}
	for i := 2; i*i <= maxn; i++ {
		if !f[i] {
			cnt++
			p[cnt] = i
		}
		for j := 1; j <= cnt && i*p[j] <= maxn; j++ {
			f[i*p[j]] = true
			if i%p[j] == 0 {
				break
			}
		}
	}

	var n int
	fmt.Fscan(in, &n)
	ans := 0
	C := 0
	a := make([]int, n+1)
	b := make([]int, (n+1)<<1)
	v := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		t := a[i]
		tmp := t
		a[i] = 1
		v[i] = 1
		for j := 1; j <= cnt && p[j]*p[j] <= tmp; j++ {
			c := 0
			for t%p[j] == 0 {
				c++
				t /= p[j]
			}
			c %= 3
			switch c {
			case 1:
				v[i] *= p[j] * p[j]
				a[i] *= p[j]
			case 2:
				v[i] *= p[j]
				a[i] *= p[j] * p[j]
			}
		}
		v[i] *= t * t
		a[i] *= t
		if v[i] == 1 {
			ans = 1
			continue
		}
		C++
		b[C] = a[i]
		C++
		b[C] = v[i]
	}
	tmp := b[1 : C+1]
	sort.Ints(tmp)
	C = len(unique(b[1 : C+1]))

	val := [maxn << 1]int{}
	link := [maxn << 1]int{}
	for i := 1; i <= n; i++ {
		if v[i] != 1 {
			a[i] = LowerBound(b[1:C+1], a[i]) + 1
			v[i] = LowerBound(b[1:C+1], v[i]) + 1
			val[a[i]]++
			link[a[i]] = v[i]
		}
	}

	vis := [maxn << 1]bool{}
	for i := 1; i <= C; i++ {
		if !vis[i] {
			vis[i] = true
			vis[link[i]] = true
			if val[i] > val[link[i]] {
				ans += val[i]
			} else {
				ans += val[link[i]]
			}
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
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func LowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
