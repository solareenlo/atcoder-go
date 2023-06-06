package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007
	const inv24 = 41666667

	var N int
	fmt.Fscan(in, &N)
	var cnt, used [100][100]int
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		cnt[x][y]++
	}
	ans := N * (N - 1) % mod * (N - 2) % mod * (N - 3) % mod * inv24 % mod
	for a := -49; a < 50; a++ {
		for b := 0; b < 50; b++ {
			tmp := a
			if a < 0 {
				tmp = -a
			}
			if a == 0 && b == 0 || a == -1 && b == 0 || gcd(tmp, b) > 1 {
				continue
			}
			id := a*100 + b + 114514
			for j := 0; j < 100; j++ {
				for i := 0; i < 100; i++ {
					if used[i][j] == id {
						continue
					}
					c := 0
					x := i
					y := j
					for 0 <= x && x < 100 && 0 <= y && y < 100 {
						c += cnt[x][y]
						used[x][y] = id
						x += a
						y += b
					}
					if c >= 3 {
						now := c * (c - 1) * (c - 2) % mod * ((N-c)*4 + c - 3) % mod * inv24 % mod
						ans = (ans - now + mod) % mod
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
