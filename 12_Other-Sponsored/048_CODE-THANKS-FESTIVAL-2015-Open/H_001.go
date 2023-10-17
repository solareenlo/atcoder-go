package main

import "fmt"

func main() {
	var H, W, K int
	fmt.Scan(&H, &W, &K)
	s := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scan(&s[i])
	}

	sum := make([][]int, H+1)
	cnt := make([][][]int, 10)

	for i := 0; i <= H; i++ {
		sum[i] = make([]int, W+1)
	}
	for i := 0; i < 10; i++ {
		cnt[i] = make([][]int, H+1)
		for j := 0; j <= H; j++ {
			cnt[i][j] = make([]int, W+1)
		}
	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + int(s[i-1][j-1]-'0')
			for k := 1; k < 10; k++ {
				cnt[k][i][j] = cnt[k][i-1][j] + cnt[k][i][j-1] - cnt[k][i-1][j-1]
			}
			cnt[int(s[i-1][j-1]-'0')][i][j]++
		}
	}

	ans := 0
	for i := 0; i < H; i++ {
		for j := i + 3; j <= H; j++ {
			r := 0
			for k := 0; k < W; k++ {
				if k+3 > r {
					r = k + 3
				}
				for r <= W && sum[j][r]-sum[i][r]-sum[j][k]+sum[i][k] <= K {
					r++
				}
				R := r
				for R <= W {
					now := sum[j][R] - sum[i][R] - sum[j][k] + sum[i][k]
					if now > K+9 {
						break
					}
					id := now - K
					ans += cnt[id][j-1][R-1] - cnt[id][i+1][R-1] - cnt[id][j-1][k+1] + cnt[id][i+1][k+1]
					R++
				}
			}
		}
	}

	fmt.Println(ans)
}
