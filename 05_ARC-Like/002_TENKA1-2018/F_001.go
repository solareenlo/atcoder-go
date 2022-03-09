package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	const mod = 998244353
	fac := make([]int, n+1)
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = i * fac[i-1] % mod
	}
	a := make([]int, 2*(n+1))
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[n+i] = a[i]
	}
	pos := 2
	for a[pos] == a[pos-1] {
		pos++
	}
	if pos == 2*n+1 && a[1] != 1 {
		fmt.Println(0)
		return
	}
	if pos == 2*n+1 {
		fmt.Println(fac[n])
		return
	}

	st := make([]int, n+1)
	ed := make([]int, n+1)
	flag := false
	for i := pos; i < pos+n; i++ {
		if a[i] != a[i-1] {
			if st[a[i]] != 0 {
				flag = true
			}
			st[a[i]] = i
		}
		if a[i] != a[i+1] {
			ed[a[i]] = i
		}
	}
	if flag {
		fmt.Println(0)
		return
	}

	t := ed[1] - st[1]
	for i := 1; i <= n; i++ {
		if ed[i]-st[i] > t {
			flag = true
		}
	}
	if flag {
		fmt.Println(0)
		return
	}

	ans := 1
	cnt := 1
	sum := t + 1
	for i := 2; i <= n; i++ {
		if st[i] == 0 {
			ans = (sum - t*cnt - i + 1) * ans % mod
			continue
		}
		k1 := 0
		if a[st[i]-1] < i {
			k1 = 1
		}
		k2 := 0
		if a[ed[i]+1] < i {
			k2 = 1
		}
		if k1 != 0 && k2 != 0 {
			ans = (st[i] - (ed[i] - t) + 1) * ans % mod
			cnt--
		}
		if k1 == 0 && k2 == 0 {
			cnt++
		}
		if k1 == 0 && k2 == 0 && ed[i]-st[i] != t {
			ans = 0
		}
		sum += ed[i] - st[i] + 1
	}
	fmt.Println(ans)
}
