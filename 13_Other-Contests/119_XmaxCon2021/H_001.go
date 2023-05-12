package main

import (
	"bufio"
	"fmt"
	"os"
)

const Mod = 998244353
const MAXN = 510

var A, B G
var pink PinkRabbit
var gauss Gauss

var IN = bufio.NewReader(os.Stdin)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(IN, &A.N, &B.N)
	A.solve()
	B.solve()

	p := make([]int, len(A.charp))
	copy(p, A.charp)
	q := make([]int, len(B.charp))
	copy(q, B.charp)
	for i := len(p) - 2; i >= 0; i -= 2 {
		p[i] = reduce(-p[i])
	}
	p = p[1:]
	q = q[1:]
	r := res(p, q)
	for i := 1; i <= A.N; i++ {
		for j := 1; j <= B.N; j++ {
			if j == B.N {
				fmt.Fprintln(out, (A.adj[i]*B.adj[j]%Mod)*r%Mod)
			} else {
				fmt.Fprintf(out, "%d ", (A.adj[i]*B.adj[j]%Mod)*r%Mod)
			}
		}
	}
}

type G struct {
	N     int
	mat   [MAXN][MAXN]int
	adj   []int
	charp []int
}

func (g *G) det() int {
	ret := 1
	for i := 1; i <= g.N; i++ {
		pivot := i
		for g.mat[pivot][i] == 0 && pivot <= g.N {
			pivot++
		}
		if pivot > g.N {
			return 0
		}
		if pivot != i {
			ret = Mod - ret
			for j := i; j <= g.N; j++ {
				g.mat[i][j], g.mat[pivot][j] = g.mat[pivot][j], g.mat[i][j]
			}
		}
		ret = ret * g.mat[i][i] % Mod
		nv := PowMod(g.mat[i][i], Mod-2)
		for j := i; j <= g.N; j++ {
			g.mat[i][j] = g.mat[i][j] * nv % Mod
		}
		for j := i + 1; j <= g.N; j++ {
			for k := g.N; k >= i; k-- {
				g.mat[j][k] = fam(g.mat[j][k], Mod-g.mat[j][i], g.mat[i][k])
			}
		}
	}
	return ret
}

func (g *G) solve() {
	for i := 1; i <= g.N; i++ {
		for j := 1; j <= g.N; j++ {
			var x int
			fmt.Fscan(IN, &x)
			g.mat[j][j] = add(g.mat[j][j], x)
			g.mat[i][j] = sub(g.mat[i][j], x)
		}
	}
	pink.N = g.N
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			pink.A[i][j] = g.mat[i][j]
		}
	}
	pink.Characteristic()
	resize(&g.charp, g.N+1)
	for i := 0; i < g.N+1; i++ {
		g.charp[i] = pink.Ans[i]
	}
	for i := 1; i <= g.N; i++ {
		for j := 1; j != i; j++ {
			g.mat[i][j], g.mat[j][i] = g.mat[j][i], g.mat[i][j]
		}
	}
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			gauss.mat[i][j] = g.mat[i][j]
		}
	}
	g.adj = make([]int, MAXN)
	gauss.solve(g.N, g.adj)
	piv := 1
	for g.adj[piv] == 0 {
		piv++
	}
	for i := 1; i <= g.N; i++ {
		g.mat[piv][i] = 0
		g.mat[i][piv] = 0
	}
	g.mat[piv][piv] = 1
	r := g.det() * PowMod(g.adj[piv], Mod-2) % Mod
	for i := 1; i <= g.N; i++ {
		g.adj[i] = g.adj[i] * r % Mod
	}
}

type PinkRabbit struct {
	N    int
	A, H [MAXN][MAXN]int
	Ans  [MAXN]int
}

func (pink *PinkRabbit) Hessenberg() {
	for i := 1; i <= pink.N; i++ {
		for j := 1; j <= pink.N; j++ {
			pink.H[i][j] = pink.A[i][j]
		}
	}
	for i := 1; i <= pink.N-2; i++ {
		if pink.H[i+1][i] == 0 {
			for j := i + 2; j <= pink.N; j++ {
				if pink.H[j][i] != 0 {
					for k := i; k <= pink.N; k++ {
						pink.H[i+1][k], pink.H[j][k] = pink.H[j][k], pink.H[i+1][k]
					}
					for k := 1; k <= pink.N; k++ {
						pink.H[k][i+1], pink.H[k][j] = pink.H[k][j], pink.H[k][i+1]
					}
					break
				}
			}
		}
		if pink.H[i+1][i] == 0 {
			continue
		}
		val := PowMod(pink.H[i+1][i], Mod-2)
		for j := i + 2; j <= pink.N; j++ {
			coef := val * pink.H[j][i] % Mod
			for k := i; k <= pink.N; k++ {
				pink.H[j][k] = (pink.H[j][k] + pink.H[i+1][k]*(Mod-coef)) % Mod
			}
			for k := 1; k <= pink.N; k++ {
				pink.H[k][i+1] = (pink.H[k][i+1] + pink.H[k][j]*coef) % Mod
			}
		}
	}
}

var P [MAXN][MAXN]int

func (pink *PinkRabbit) Characteristic() {
	pink.Hessenberg()
	for i := 1; i <= pink.N; i++ {
		for j := 1; j <= pink.N; j++ {
			pink.H[i][j] = Mod - pink.H[i][j]
		}
	}
	P[0][0] = 1
	for i := 1; i <= pink.N; i++ {
		P[i][0] = 0
		for j := 1; j <= i; j++ {
			P[i][j] = P[i-1][j-1]
		}
		val := 1
		for j := i - 1; j >= 0; j-- {
			coef := val * pink.H[j+1][i] % Mod
			for k := 0; k <= j; k++ {
				P[i][k] = (P[i][k] + P[j][k]*coef) % Mod
			}
			if j != 0 {
				val = val * (Mod - pink.H[j+1][j]) % Mod
			}
		}
	}
	for i := 0; i <= pink.N; i++ {
		pink.Ans[i] = P[pink.N][i]
	}
}

type Gauss struct {
	bas, mat [MAXN][MAXN]int
}

func (g *Gauss) solve(N int, adj []int) {
	for i := range g.bas {
		for j := range g.bas[i] {
			g.bas[i][j] = 0
		}
	}
	for i := 1; i <= N; i++ {
		g.bas[i][i] = 1
	}
	for i := 1; i <= N; i++ {
		piv := 1
		for piv <= N && g.mat[i][piv] == 0 {
			piv++
		}
		if piv > N {
			for j := 0; j < MAXN; j++ {
				adj[j] = g.bas[i][j]
			}
			break
		}
		c := PowMod(g.mat[i][piv], Mod-2)
		for j := piv; j <= N; j++ {
			g.mat[i][j] = g.mat[i][j] * c % Mod
		}
		for j := 1; j <= i; j++ {
			g.bas[i][j] = g.bas[i][j] * c % Mod
		}
		for j := i + 1; j <= N; j++ {
			c = Mod - g.mat[j][piv]
			for k := piv; k <= N; k++ {
				g.mat[j][k] = fam(g.mat[j][k], g.mat[i][k], c)
			}
			for k := 1; k <= i; k++ {
				g.bas[j][k] = fam(g.bas[j][k], g.bas[i][k], c)
			}
		}
	}
}

func res(p, q []int) int {
	ret := 1
	N := len(p) - 1
	M := len(q) - 1
	if N < M {
		N, M = M, N
		p, q = q, p
		if ((N * M) % 2) != 0 {
			ret = reduce(-ret)
		}
	}
	for M != 0 {
		for i := N - M; i >= 0; i-- {
			c := reduce(-p[M+i])
			for j := 0; j <= M; j++ {
				p[j+i] = fam(p[j+i], q[j], c)
			}
		}
		for N != 0 && p[N] == 0 {
			N--
		}
		resize(&p, N+1)
		if p[N] == 0 {
			return 0
		}
		ret = ret * PowMod(p[N], M) % Mod
		c := PowMod(p[N], Mod-2)
		for i := 0; i < N+1; i++ {
			p[i] = p[i] * c % Mod
		}
		N, M = M, N
		p, q = q, p
		if ((N * M) % 2) != 0 {
			ret = reduce(-ret)
		}
	}
	return ret
}

func reduce(x int) int {
	if x < 0 {
		return x + Mod
	}
	return x
}

func add(x int, y int) int {
	x += y - Mod
	if x < 0 {
		x += Mod
	}
	return x
}

func sub(x, y int) int {
	x -= y
	if x < 0 {
		x += Mod
	}
	return x
}

func fam(x, y, z int) int {
	x = (x + y*z) % Mod
	return x
}

func PowMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % Mod
		}
		a = a * a % Mod
		n /= 2
	}
	return res
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}
