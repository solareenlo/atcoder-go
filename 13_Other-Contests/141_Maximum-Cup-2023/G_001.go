package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 4001001001001001001

type LiCiaoTree struct {
	N   int
	mid []int
	A   []pair
}

func (l *LiCiaoTree) init(X []int) {
	n := len(X)
	l.N = 1
	for l.N < n {
		l.N *= 2
	}
	P := make([]int, l.N)
	for i := range P {
		P[i] = X[len(X)-1]
	}
	for i := 0; i < n; i++ {
		P[i] = X[i]
	}
	l.mid = make([]int, l.N*2)
	for d := 0; (1 << d) < l.N; d++ {
		q := l.N >> d
		for i := 0; i < 1<<d; i++ {
			l.mid[(1<<d)+i] = P[q*i+q/2]
		}
	}
	for i := 0; i < l.N; i++ {
		l.mid[l.N+i] = P[i]
	}
	l.A = make([]pair, l.N*2)
	for i := range l.A {
		l.A[i] = pair{0, -INF}
	}
}

func (l LiCiaoTree) maxval(pos int) int {
	ans := -INF
	x := l.mid[l.N+pos]
	pos += l.N
	for pos >= 1 {
		ans = max(ans, l.A[pos].x*x+l.A[pos].y)
		pos /= 2
	}
	return ans
}

func (l *LiCiaoTree) update(a, b int) {
	i := 1
	f := pair{a, b}
	for {
		if l.A[i].x*l.mid[i]+l.A[i].y < f.x*l.mid[i]+f.y {
			l.A[i], f = f, l.A[i]
		}
		if i >= l.N {
			break
		}
		if f.x-l.A[i].x > 0 {
			i = i*2 + 1
		} else {
			i = i * 2
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	var X, Y int
	fmt.Fscan(in, &X, &Y)
	points := make([]pair, N)
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x -= X
		y -= Y
		points[i] = pair{x, y}
	}
	lines := make([]pair, M)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		b -= Y
		b += a * X
		lines[i] = pair{a, b}
	}
	defaultVal := -1
	maxLine := M
	for i := 0; i < M; i++ {
		if lines[i].y == 0 {
			maxLine = i
			defaultVal = powMod(2, i)
			break
		}
	}
	resize(&lines, maxLine)
	M = maxLine
	var splitPoints func([]pair, int, int) int
	splitPoints = func(lns []pair, pl, pr int) int {
		if pl == pr {
			return pl
		}
		if len(lns) == 0 {
			return pl
		}
		sortPair(points[pl:pr])
		pos := make([]int, pr-pl)
		for i := 0; i < pr-pl; i++ {
			pos[i] = points[pl+i].x
		}
		var upperConvex, lowerConvex LiCiaoTree
		upperConvex.init(pos)
		lowerConvex.init(pos)
		for _, tmp := range lns {
			a, b := tmp.x, tmp.y
			if b > 0 {
				upperConvex.update(-a, -b)
			} else {
				lowerConvex.update(a, b)
			}
		}
		pass := make([]int, pr-pl)
		for i := pl; i < pr; i++ {
			y := points[i].y
			if -upperConvex.maxval(i-pl) <= y {
				pass[i-pl] = 1
				continue
			}
			if lowerConvex.maxval(i-pl) >= y {
				pass[i-pl] = 1
				continue
			}
		}
		for i := pl; i < pr; i++ {
			if pass[i-pl] != 0 {
				continue
			}
			pass[i-pl], pass[pr-pl-1] = pass[pr-pl-1], pass[i-pl]
			points[i], points[pr-1] = points[pr-1], points[i]
			i--
			pr--
		}
		return pr
	}
	if splitPoints(lines, 0, N) != N {
		fmt.Println(defaultVal)
		return
	}
	var dfs func(int, int, int, int) []int
	dfs = func(l, r, pl, pr int) []int {
		if pl == pr {
			return []int{}
		}
		if r == l+1 {
			return []int{l}
		}
		m := (l + r) / 2
		pm := splitPoints(lines[l:m], pl, pr)
		ans := dfs(m, r, pm, pr)
		uselines := make([]pair, 0)
		for _, a := range ans {
			uselines = append(uselines, lines[a])
		}
		pl = splitPoints(uselines, pl, pm)
		ans2 := dfs(l, m, pl, pm)
		for _, i := range ans2 {
			ans = append(ans, i)
		}
		return ans
	}
	ans := dfs(0, maxLine, 0, N)
	fans := 0
	for _, i := range ans {
		fans = (fans + powMod(2, i)) % MOD
	}
	fmt.Println(fans)
}

func resize(a *[]pair, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, pair{0, 0})
		}
	}
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const MOD = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}
