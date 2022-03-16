package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W, K int
	fmt.Fscan(in, &H, &W, &K)

	hw := make([][]string, H)
	for i := range hw {
		hw[i] = strings.Split(strings.Repeat("?", W), "")
	}
	for i := 0; i < K; i++ {
		var h, w int
		var c string
		fmt.Fscan(in, &h, &w, &c)
		h--
		w--
		hw[h][w] = c
	}

	dp := make([][]int, H+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	dp[0][0] = powMod(3, H*W-K)
	inv23 := divMod(2, 3)
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			d := dp[h][w]
			if hw[h][w] == "?" {
				d *= inv23
				d %= mod
			}
			if hw[h][w] != "D" {
				dp[h][w+1] += d
				dp[h][w+1] %= mod
			}
			if hw[h][w] != "R" {
				dp[h+1][w] += d
				dp[h+1][w] %= mod
			}
		}
	}
	fmt.Println(dp[H-1][W-1])
}

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
