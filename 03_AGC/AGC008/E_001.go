package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1000005
	const P = 1000000007

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	deg := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		deg[a[i]]++
	}

	vs := make([]int, N)
	pre := make([]int, N)
	var x int
	for i := 1; i <= n; i++ {
		if deg[i] > 2 {
			fmt.Println(0)
			return
		}
		if deg[i] < 2 || vs[i] != 0 {
			continue
		}
		for x = i; (vs[x] == 0) || (x != i); x = a[x] {
			if vs[x] != 0 {
				fmt.Println(0)
				return
			}
			vs[x] = 1
			pre[a[x]] = x
		}
	}

	ans := 1
	for i := 1; i <= n; i++ {
		if deg[i] == 0 {
			l1 := 0
			l2 := 0
			for x = i; vs[x] == 0; x = a[x] {
				vs[x] = 1
				l1++
			}
			x = pre[x]
			l2++
			for deg[x]^2 != 0 {
				x = pre[x]
				l2++
			}
			if l1 < l2 {
				ans <<= 1
				if ans >= P {
					ans -= P
				}
			}
			if l1 > l2 {
				fmt.Println(0)
				return
			}
		}
	}

	cnt := make([]int, N)
	for i := 1; i <= n; i++ {
		if vs[i] == 0 {
			l := 0
			x = i
			x = a[x]
			vs[x] = 1
			l++
			for x^i != 0 {
				x = a[x]
				vs[x] = 1
				l++
			}
			cnt[l]++
		}
	}

	f := make([]int, N)
	for i := 1; i <= n; i++ {
		tmp := 0
		if i != 1 && (i&1) != 0 {
			tmp = 1
		}
		x = 1 + tmp
		f[0] = 1
		f[1] = x
		for j := 2; j <= cnt[i]; j++ {
			f[j] = ((j-1)*f[j-2]%P*i + f[j-1]*x) % P
		}
		ans = ans * f[cnt[i]] % P
	}
	fmt.Println(ans)
}
