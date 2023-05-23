package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var n int
	fmt.Fscan(in, &n)
	var a [5][]int
	for i := range a {
		a[i] = make([]int, n)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
		sort.Ints(a[i])
	}
	var l [5]int
	ans := 0
	for i := 0; i < 5*n; i++ {
		id := -1
		mi := 1001001001
		for j := 0; j < 5; j++ {
			if l[j] < n && mi > a[j][l[j]] {
				mi = a[j][l[j]]
				id = j
			}
		}
		var p1, p2, p3, p4 int
		if id < 1 {
			p1 = l[1]
		} else {
			p1 = l[0]
		}
		if id < 2 {
			p2 = l[2]
		} else {
			p2 = l[1]
		}
		if id < 3 {
			p3 = l[3]
		} else {
			p3 = l[2]
		}
		if id < 4 {
			p4 = l[4]
		} else {
			p4 = l[3]
		}
		tmp1 := (((p1 * p2 % MOD) * p3 % MOD) * p4 % MOD) * 6 % MOD
		tmp2 := ((p1*p2%MOD)*p3%MOD + (p1*p2%MOD)*p4%MOD + (p1*p3%MOD)*p4%MOD + (p2*p3%MOD)*p4%MOD) % MOD
		tmp3 := (p1*p2%MOD + p1*p3%MOD + p1*p4%MOD + p2*p3%MOD + p2*p4%MOD + p3*p4%MOD) % MOD
		ans = (ans + (((tmp1-(tmp2*3%MOD)*n%MOD+MOD)%MOD+(tmp3*n%MOD)*n%MOD)%MOD)*mi%MOD) % MOD
		l[id]++
	}
	fmt.Println(ans)
}
