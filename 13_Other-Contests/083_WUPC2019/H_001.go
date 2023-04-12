package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007
const sz = 400

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	C := make([]club, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &C[i].K)
		for j := 0; j < C[i].K; j++ {
			var a int
			fmt.Fscan(in, &a)
			C[i].A = append(C[i].A, a)
		}
		C[i].check()
	}

	type pair struct {
		x, y int
	}
	mp := make(map[pair]int)

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if x > y {
			x, y = y, x
		}
		if _, ok := mp[pair{x, y}]; ok {
			fmt.Fprintln(out, mp[pair{x, y}])
			continue
		}
		if C[x].K > C[y].K {
			x, y = y, x
		}
		ans := 0
		if C[y].K <= sz || C[x].K > sz {
			for i := 0; i < C[y].K; i++ {
				ans = (ans + (C[y].A[i] * C[x].A[i%C[x].K] % mod)) % mod
			}
		} else {
			for i := 0; i < C[x].K; i++ {
				ans = (ans + (C[x].A[i] * C[y].L[C[x].K-1][i] % mod)) % mod
			}
		}
		if x > y {
			x, y = y, x
		}
		mp[pair{x, y}] = ans
		fmt.Fprintln(out, ans)
	}
}

type club struct {
	K int
	A []int
	L [][]int
}

func (c *club) check() {
	if c.K <= sz {
		return
	}
	c.L = make([][]int, sz)
	for i := 0; i < sz; i++ {
		c.L[i] = make([]int, i+1)
		for j := 0; j < c.K; j++ {
			c.L[i][j%(i+1)] = (c.L[i][j%(i+1)] + c.A[j]) % mod
		}
	}
}
