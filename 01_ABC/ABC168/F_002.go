package main

import (
	"fmt"
	"sort"
)

const INF = 1 << 60
const maxn = 1010

var (
	n, m   int
	k      = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	border = [maxn][maxn][4]int{}
	vis    = [maxn][maxn]int{}
	qx     = [maxn * maxn]int{}
	qy     = [maxn * maxn]int{}
	ans    int
	X      = make([]int, 0)
	Y      = make([]int, 0)
	XN, YN int
	f      = [maxn]node{}
	g      = [maxn]node{}
)

type node struct{ a, b, c int }

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func bfs(xx, yy int) {
	var nx, ny, x, y int
	h, t := 0, 0
	qx[t] = xx
	qy[t] = yy
	t++
	vis[xx][yy] = 1
	for h < t {
		x = qx[h]
		y = qy[h]
		h++
		ans += (X[x+1] - X[x]) * (Y[y+1] - Y[y])
		for i := 0; i < 4; i++ {
			if border[x][y][i] != 0 {
				continue
			}
			nx = x + k[i][0]
			ny = y + k[i][1]
			if nx == 0 || nx == XN-1 || ny == 0 || ny == YN-1 {
				ans = -1
				return
			}
			if vis[nx][ny] != 0 {
				continue
			}
			vis[nx][ny] = 1
			qx[t] = nx
			qy[t] = ny
			t++
		}
	}
}

func SliceUnique(target []int) (unique []int) {
	m := map[int]bool{}
	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

func main() {
	var x, y, r int
	fmt.Scan(&n, &m)
	X = make([]int, n+2)
	XN = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&f[i].a, &f[i].b, &f[i].c)
		X[XN] = f[i].c
		XN++
	}
	X[XN] = -INF
	XN++
	X[XN] = INF
	XN++
	sort.Ints(X)
	X = SliceUnique(X)
	XN = len(X)

	Y = make([]int, m+2)
	for i := 0; i < m; i++ {
		fmt.Scan(&g[i].a, &g[i].b, &g[i].c)
		Y[YN] = g[i].a
		YN++
	}
	Y[YN] = -INF
	YN++
	Y[YN] = INF
	YN++
	sort.Ints(Y)
	Y = SliceUnique(Y)
	YN = len(Y)

	for i := 0; i < n; i++ {
		x = lowerBound(X, f[i].c)
		r = lowerBound(Y, f[i].b)
		if Y[r] != f[i].b {
			r--
		}
		for j := lowerBound(Y, f[i].a); j < r; j++ {
			border[x-1][j][1] = 1
			border[x][j][0] = 1
		}
	}
	for i := 0; i < m; i++ {
		y = lowerBound(Y, g[i].a)
		r = lowerBound(X, g[i].c)
		if X[r] != g[i].c {
			r--
		}
		for j := lowerBound(X, g[i].b); j < r; j++ {
			border[j][y-1][3] = 1
			border[j][y][2] = 1
		}
	}

	bfs(lowerBound(X, 0)-1, lowerBound(Y, 0)-1)

	if ans == -1 {
		fmt.Println("INF")
	} else {
		fmt.Println(ans)
	}
}
