package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const maxn = 100010

var n int
var a, b, c [maxn]int
var sabc, sab, sac, sbc [maxn]int
var tmp [maxn]int
var od []int
var bit [maxn]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
	}
	od = make([]int, n)
	for i := 0; i < n; i++ {
		od[i] = i
	}
	sort.Slice(od, func(x, y int) bool {
		return a[od[x]] < a[od[y]]
	})
	cdq(0, n-1)
	for i := range bit {
		bit[i] = 0
	}
	for i := 0; i < n; i++ {
		sbc[od[i]] = sum(c[od[i]])
		add(c[od[i]], 1)
	}

	for i := range bit {
		bit[i] = 0
	}
	sort.Slice(od, func(x, y int) bool {
		return a[od[x]] < a[od[y]]
	})
	for i := 0; i < n; i++ {
		sab[od[i]] = sum(b[od[i]])
		add(b[od[i]], 1)
	}

	for i := range bit {
		bit[i] = 0
	}
	for i := 0; i < n; i++ {
		sac[od[i]] = sum(c[od[i]])
		add(c[od[i]], 1)
	}

	for i := 0; i < n; i++ {
		sab[i] -= sabc[i]
		sac[i] -= sabc[i]
		sbc[i] -= sabc[i]
	}

	one := n * (n - 1) * (n - 2) / 6
	for i := 0; i < n; i++ {
		one -= sab[i] * (sab[i] - 1) / 2
		one -= sac[i] * (sac[i] - 1) / 2
		one -= sbc[i] * (sbc[i] - 1) / 2
		one -= (sab[i] + sbc[i] + sac[i]) * sabc[i]
		one -= sabc[i] * (sabc[i] - 1) / 2
	}

	two := 0
	for i := 0; i < n; i++ {
		two += sab[i] + sac[i] + sbc[i]
	}

	three := n
	fmt.Println(one + two + three)

}

func add(x, v int) {
	for x < maxn {
		bit[x] += v
		x += x & -x
	}
}

func sum(x int) int {
	ret := 0
	for x > 0 {
		ret += bit[x]
		x -= x & -x
	}
	return ret
}

func cdq(l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	cdq(l, mid)
	cdq(mid+1, r)
	i := l
	j := mid + 1
	pos := l
	for i <= mid && j <= r {
		x := od[i]
		y := od[j]
		if b[x] < b[y] {
			add(c[x], 1)
			tmp[pos] = x
			pos++
			i++
		} else {
			sabc[y] += sum(c[y])
			tmp[pos] = y
			pos++
			j++
		}
	}
	for i <= mid {
		x := od[i]
		add(c[x], 1)
		tmp[pos] = x
		pos++
		i++
	}
	for j <= r {
		y := od[j]
		sabc[y] += sum(c[y])
		tmp[pos] = y
		pos++
		j++
	}
	for i = l; i <= mid; i++ {
		add(c[od[i]], -1)
	}
	for i = l; i <= r; i++ {
		od[i] = tmp[i]
	}
}
