package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const big = int(1e18)

	var H, W int
	fmt.Fscan(in, &H, &W)
	var rink [777][777]int
	num := 0
	sx, sy := 0, 0
	for i := 0; i < H; i++ {
		var str string
		fmt.Fscan(in, &str)
		for j := 0; j < W; j++ {
			rink[i][j] = -1
			if str[j] == '#' {
				rink[i][j] = num
				num++
			}
			if str[j] == 'G' {
				sy = i
				sx = j
			}
		}
	}
	rink[sy][sx] = num
	num++
	type pair struct {
		x, y int
	}
	Go := make([][]pair, num)
	for i := 0; i < H; i++ {
		mae := -1
		bas := -1
		for j := 0; j < W; j++ {
			if rink[i][j] == -1 {
				continue
			}
			if mae != -1 {
				Go[rink[i][j]] = append(Go[rink[i][j]], pair{mae, j - bas})
				Go[mae] = append(Go[mae], pair{rink[i][j], j - bas})
			}
			mae = rink[i][j]
			bas = j
		}
	}
	for j := 0; j < W; j++ {
		mae := -1
		bas := -1
		for i := 0; i < H; i++ {
			if rink[i][j] == -1 {
				continue
			}
			if mae != -1 {
				Go[rink[i][j]] = append(Go[rink[i][j]], pair{mae, i - bas})
				Go[mae] = append(Go[mae], pair{rink[i][j], i - bas})
			}
			mae = rink[i][j]
			bas = i
		}
	}
	var time [1555][7778]int
	for i := 0; i <= H+W; i++ {
		for j := 0; j < num-1; j++ {
			time[i][j] = big
		}
	}
	for i := H + W; i > 0; i-- {
		for j := 0; j < num; j++ {
			for _, it := range Go[j] {
				time[i-1][it.x] = min(time[i-1][it.x], time[i][j]+i*it.y)
			}
		}
	}
	var ans [777][777]int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			ans[i][j] = big
			if rink[i][j] >= 0 {
				ans[i][j] = time[1][rink[i][j]]
			}
		}
	}
	for i := 0; i < H; i++ {
		gen := big
		for j := 0; j < W; j++ {
			gen++
			ans[i][j] = min(ans[i][j], gen)
			if rink[i][j] >= 0 {
				gen = time[1][rink[i][j]]
			}
		}
	}
	for i := 0; i < H; i++ {
		gen := big
		for j := W - 1; j >= 0; j-- {
			gen++
			ans[i][j] = min(ans[i][j], gen)
			if rink[i][j] >= 0 {
				gen = time[1][rink[i][j]]
			}
		}
	}
	for j := 0; j < W; j++ {
		gen := big
		for i := 0; i < H; i++ {
			gen++
			ans[i][j] = min(ans[i][j], gen)
			if rink[i][j] >= 0 {
				gen = time[1][rink[i][j]]
			}
		}
	}
	for j := 0; j < H; j++ {
		gen := big
		for i := H - 1; i >= 0; i-- {
			gen++
			ans[i][j] = min(ans[i][j], gen)
			if rink[i][j] >= 0 {
				gen = time[1][rink[i][j]]
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if ans[i][j] == big {
				ans[i][j] = -1
			}
			fmt.Fprintf(out, "%d ", ans[i][j])
		}
		fmt.Fprintln(out)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
