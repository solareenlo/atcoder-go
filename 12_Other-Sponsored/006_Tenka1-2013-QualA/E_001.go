package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

var dy [4]int = [4]int{0, 0, 2, 1}
var dx [4]int = [4]int{0, 2, 0, -1}
var ry [4][4]int = [4][4]int{
	{0, 1, 1, 2},
	{0, 0, 1, 0},
	{0, -1, -1, -2},
	{0, 0, -1, 0},
}
var rx [4][4]int = [4][4]int{
	{0, 0, 1, 0},
	{0, -1, -1, -2},
	{0, 0, -1, 0},
	{0, 1, 1, 2},
}

var H, W, N int
var S [12]string
var P, py, px, pd [36]int
var a [12][12]int
var b [20][20]bool
var used [40]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &H, &W)
	cnt1 := 0
	cnt2 := 0
	for y := 0; y < H; y++ {
		fmt.Fscan(in, &S[y])
		for x := 0; x < W; x++ {
			if S[y][x] == '#' {
				cnt1++
				a[y][x] = 1
			} else {
				a[y][x] = 0
			}
		}
	}
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &P[i])
		cnt2 += popcount(uint32(P[i] - 1))
	}

	if cnt1 != cnt2 || H%4 != 0 || W%4 != 0 {
		fmt.Println(-1)
		return
	}

	if !dfs(0, 0) {
		fmt.Println(-1)
	}
}

func dfs(x, y int) bool {
	if y == H {
		for i := 0; i < N; i++ {
			fmt.Println(px[i]+1, py[i]+1, pd[i])
		}
		return true
	}

	nx := x + 1
	ny := y
	if nx == W {
		ny++
		nx = 0
	}
	if b[y][x] {
		return dfs(nx, ny)
	}

	for d := 0; d < 4; d++ {
		x += dx[d]
		y += dy[d]
		key := 1
		flag := true
		for i := 0; i < 4; i++ {
			tx := x + rx[d][i]
			ty := y + ry[d][i]
			if tx < 0 || ty < 0 || W <= tx || H <= ty || b[ty][tx] {
				flag = false
				break
			}
			key += a[ty][tx] * (1 << i)
		}
		if flag {
			for i := 0; i < N; i++ {
				if !used[i] && P[i] == key {
					used[i] = true
					px[i] = x
					py[i] = y
					pd[i] = d
					for j := 0; j < 4; j++ {
						b[y+ry[d][j]][x+rx[d][j]] = true
					}
					if dfs(nx, ny) {
						return true
					}
					for j := 0; j < 4; j++ {
						b[y+ry[d][j]][x+rx[d][j]] = false
					}
					used[i] = false
					break
				}
			}
		}
		x -= dx[d]
		y -= dy[d]
	}
	return false
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}
