package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const MOD = 998244353
const N = 105
const M = 20005
const INF = 1_000_000_000
const inv2 = 499122177

var (
	n   int
	m   int
	tX  int
	tY  int
	sX  int
	sY  int
	lst int
	e   = make([][]int, M)
	dst = make([]int, M)
	z   = make([]int, M)
	q   = make([]int, M)
	tg  = make([]int, M)
	st  = make([]int, M)
	a   = make([]string, N)
)

func W(x, y int) int {
	x += y
	if x >= MOD {
		x -= MOD
	}
	return x
}

func W1(x1, x2 *int, w1, w2 int) {
	if w1 < *x1 {
		*x1 = w1
		*x2 = w2
	} else if w1 == *x1 {
		*x2 = W(*x2, w2)
	}
}

func f(x, y int) int { return (x-1)*m + y }

func chk(x, y, x1, y1 int) bool {
	if x1 == -1 && x == tX+1 && y+y1 >= sY && y+y1 <= tY {
		return true
	}
	if x1 == 1 && x == tX && y >= sY && y <= tY {
		return true
	}
	if y1 == -1 && y == sY && x >= sX && x <= tX {
		return true
	}
	if y1 == 1 && y == sY-1 && x+x1 >= sX && x+x1 <= tX {
		return true
	}
	return false
}

func uni(a []int) []int {
	lst = 0
	res := make([]int, 0)
	sort.Ints(a)
	for _, i := range a {
		if i != lst {
			res = append(res, i)
			lst = i
		}
	}
	return res
}

func addE(u, v int, fl bool) {
	if fl {
		e[u] = append(e[u], n*m+v)
		e[n*m+u] = append(e[n*m+u], v)
	} else {
		e[u] = append(e[u], v)
		e[n*m+u] = append(e[n*m+u], n*m+v)
	}
}

func bfs(S int) {
	for i := 1; i <= n*m*2; i++ {
		dst[i] = INF
		z[i] = 0
	}
	q[0] = 2
	q[1] = 1
	q[1]++
	q[q[1]] = S
	dst[S] = 0
	z[S] = 1
	for q[0] <= q[1] {
		u := q[q[0]]
		q[0]++
		if tg[u] > tg[S] {
			continue
		}
		for _, v := range e[u] {
			if dst[v] == INF {
				dst[v] = dst[u] + 1
				z[v] = z[u]
				q[1]++
				q[q[1]] = v
			} else if dst[v] == dst[u]+1 {
				z[v] = W(z[v], z[u])
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	a[0] = strings.Repeat(" ", m+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] = " " + a[i] + " "
	}
	a[n+1] = strings.Repeat(" ", m+2)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i][j] == 'S' {
				sX = i
				sY = j
			}
			if a[i][j] == 'G' {
				tX = i
				tY = j
			}
		}
	}
	if sX > tX {
		sX = n - sX + 1
		tX = n - tX + 1
		tmp := a[1 : n+1]
		tmp = reverseOrderString(tmp)
		for i := 0; i < n; i++ {
			a[i+1] = tmp[i]
		}
	}
	if sY > tY {
		sY = m - sY + 1
		tY = m - tY + 1
		for i := 1; i <= n; i++ {
			a[i] = " " + reverseString(a[i][1:m+1]) + " "
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i][j] == '.' {
				if a[i-1][j] == '.' {
					addE(f(i, j), f(i-1, j), chk(i, j, -1, 0))
				}
				if a[i][j-1] == '.' {
					addE(f(i, j), f(i, j-1), chk(i, j, 0, -1))
				}
				if a[i+1][j] == '.' {
					addE(f(i, j), f(i+1, j), chk(i, j, 1, 0))
				}
				if a[i][j+1] == '.' {
					addE(f(i, j), f(i, j+1), chk(i, j, 0, 1))
				}
				if a[i-1][j-1] == '.' {
					addE(f(i, j), f(i-1, j-1), chk(i, j, -1, -1))
				}
				if a[i-1][j+1] == '.' {
					addE(f(i, j), f(i-1, j+1), chk(i, j, -1, 1))
				}
				if a[i+1][j-1] == '.' {
					addE(f(i, j), f(i+1, j-1), chk(i, j, 1, -1))
				}
				if a[i+1][j+1] == '.' {
					addE(f(i, j), f(i+1, j+1), chk(i, j, 1, 1))
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if (i == 1 || j == 1 || i == n || j == m) && a[i][j] == '.' {
				st[0]++
				st[st[0]] = f(i, j)
				if (i == n || j == 1) && chk(i, j, 1, -1) {
					tg[st[0]] = 1
				} else {
					tg[st[0]] = 0
				}
			}
		}
	}
	for i := 1; i <= st[0]; i++ {
		for j := 1; j <= st[0]; j++ {
			if i != j {
				if tg[i]^tg[j] != 0 {
					addE(st[i], st[j], true)
				} else {
					addE(st[i], st[j], false)
				}
			}
		}
	}
	for i := 1; i <= st[0]; i++ {
		st[i] = 0
		tg[i] = 0
	}

	st[0] = 0
	for i := 1; i <= n*m*2; i++ {
		e[i] = uni(e[i])
	}
	for i := sX; i <= tX; i++ {
		for j := sY; j <= tY; j++ {
			if (i == sX || j == sY || i == tX || j == tY) && a[i][j] == '.' {
				u := f(i, j)
				st[0]++
				tg[n*m+u] = st[0]
				tg[u] = st[0]
				st[tg[u]] = u
			}
		}
	}

	ans1, ans2 := INF, 0
	for i := 1; i <= st[0]; i++ {
		bfs(st[i])
		W1(&ans1, &ans2, dst[n*m+st[i]], z[n*m+st[i]])
	}
	if ans1 == INF {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	fmt.Println(ans1, ans2*inv2%MOD)
}

func reverseOrderString(a []string) []string {
	n := len(a)
	res := make([]string, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
