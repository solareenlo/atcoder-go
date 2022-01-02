package main

import "fmt"

var (
	dx = [5]int{0, 0, 0, 1, -1}
	dy = [5]int{0, 1, -1, 0, 0}
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sx, sy, tx, ty := 0, 0, 0, 0
	a := [505][505]int{}
	for i := 1; i <= n; i++ {
		var b string
		fmt.Scan(&b)
		b = " " + b
		for j := 1; j <= m; j++ {
			if b[j] == 's' {
				sx = i
				sy = j
				a[i][j] = 10
			} else if b[j] == 'g' {
				tx = i
				ty = j
				a[i][j] = 10
			} else if b[j] == '#' {
				a[i][j] = -1
			} else {
				a[i][j] = int(b[j] ^ 48)
			}
		}
	}

	ans := [505][505]float64{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			ans[i][j] = -1.0
		}
	}

	const eps = 1e-10
	type pair struct{ x, y int }
	q := make([]pair, 0)
	q = append(q, pair{tx, ty})
	ans[tx][ty] = 10
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for i := 1; i <= 4; i++ {
			v := pair{u.x + dx[i], u.y + dy[i]}
			if v.x < 1 || v.x > n || v.y < 1 || v.y > m || a[v.x][v.y] == -1 {
				continue
			}
			if min(ans[u.x][u.y]*0.99, float64(a[v.x][v.y]))-ans[v.x][v.y] > eps {
				ans[v.x][v.y] = min(ans[u.x][u.y]*0.99, float64(a[v.x][v.y]))
				q = append(q, v)
			}
		}
	}
	if ans[sx][sy] == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans[sx][sy])
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
