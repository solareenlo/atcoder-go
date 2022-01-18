package main

import (
	"bufio"
	"fmt"
	"os"
)

const B = 3001

var (
	mod int = 1_000_000_007
	sum     = [B + 4][B + 4]int{}
	k   int
	com = [100001]int{}
)

func calc(x, y int) int {
	if y == 0 {
		return 1
	}
	ret := calc(x, y/2)
	ret = ret * ret % mod
	if y&1 != 0 {
		ret = ret * x % mod
	}
	return ret
}

func rect(xs, xe, ys, ye int) int {
	return sum[xe][ye] - sum[xs-1][ye] - sum[xe][ys-1] + sum[xs-1][ys-1]
}

func can(p int) bool {
	for i := 1; i+p <= B; i++ {
		for j := 1; j+p <= B; j++ {
			if rect(i, i+p, j, j+p) >= k {
				return true
			}
		}
	}
	return false
}

func cnt(p int) int {
	ret := 0
	for i := 1; i+p <= B; i++ {
		for j := 1; j+p <= B; j++ {
			ret = (ret + com[rect(i, i+p, j, j+p)]) % mod
			if i > 1 {
				ret = (ret + mod - com[rect(i, i+p-1, j, j+p)]) % mod
			}
			if j > 1 {
				ret = (ret + mod - com[rect(i, i+p, j, j+p-1)]) % mod
			}
		}
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n, &k)

	x := make([]int, n+1)
	y := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		sum[x[i]+1][y[i]+1]++
	}
	for i := 1; i <= B; i++ {
		for j := 1; j <= B; j++ {
			sum[i][j] += sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}

	inv := make([]int, n+1)
	for i := 1; i <= n; i++ {
		inv[i] = calc(i, mod-2)
	}

	com[k] = 1
	for i := k + 1; i <= n; i++ {
		com[i] = com[i-1] * i % mod * inv[i-k] % mod
	}

	start := 0
	end := B - 1
	for start < end {
		mid := (start + end) / 2
		if can(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	fmt.Println(start)
	fmt.Println(cnt(start))
}
