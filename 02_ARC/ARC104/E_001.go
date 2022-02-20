package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 1_000_000_007
const N = 10

var (
	inv = [N]int{}
	lis = [N]int{}
)

func powMod(n, k int) int {
	ret := 1
	for k > 0 {
		if k&1 != 0 {
			ret = ret * n % mod
		}
		n = n * n % mod
		k >>= 1
	}
	return ret
}

func nCk(n, k int) int {
	ret := 1
	for i := 1; i <= k; i++ {
		ret = ret * (n - i + 1) % mod * inv[i] % mod
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i]
	}
	sort.Ints(b)

	for i := 1; i <= n; i++ {
		inv[i] = powMod(i, mod-2)
	}

	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = i
	}

	ans := 0
	for {
		dp := [N]int{}
		dp[0] = 1
		for i := 0; i < n; i++ {
			L := 0
			if i > 0 {
				L = b[i-1]
			}
			R := b[i]
			for l := n - 1; l >= 0; l-- {
				cnt := 0
				for r := l + 1; r <= n; r++ {
					if a[pos[r-1]] <= L {
						break
					}
					if r-l >= 2 && pos[r-1] < pos[r-2] {
						cnt++
					}
					dp[r] = (dp[r] + dp[l]*nCk(R-L+cnt, r-l)) % mod
				}
			}
		}
		leng := 0
		for i := 0; i < n; i++ {
			lis[i] = 1
			for j := 0; j < i; j++ {
				if pos[j] < pos[i] {
					lis[i] = max(lis[i], lis[j]+1)
				}
			}
			leng = max(leng, lis[i])
		}
		ans = (ans + leng*dp[n]) % mod
		if !nextPermutation(sort.IntSlice(pos)) {
			break
		}
	}

	for i := 0; i < n; i++ {
		ans = ans * powMod(a[i], mod-2) % mod
	}
	fmt.Println(ans)
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
