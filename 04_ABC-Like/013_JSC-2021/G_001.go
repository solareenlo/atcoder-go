package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	uf := New(N)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if a[i][j] == 1 {
				if uf.Same(i, j) {
					fmt.Println(0)
					return
				}
				uf.Merge(i, j)
			}
		}
	}

	g := uf.Groups()
	n := len(g)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < len(g[i]); k++ {
				for l := 0; l < len(g[j]); l++ {
					if a[g[i][k]][g[j][l]] == -1 {
						mat[i][j] -= 1
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		mat[i][i] -= sum(mat[i])
	}

	n--
	mat = mat[:n]
	ans := 1
	for i := 0; i < n; i++ {
		if mat[i][i] == 0 {
			for j := i + 1; j < n; j++ {
				if mat[j][i] != 0 {
					mat[i], mat[j] = mat[j], mat[i]
					ans *= -1
					break
				}
			}
			if mat[i][i] == 0 {
				fmt.Println(0)
				return
			}
		}
		ans *= mat[i][i]
		ans %= mod
		a := invMod(mat[i][i])
		for j := i + 1; j < n; j++ {
			c := a * mat[j][i] % mod
			for k := 0; k < n; k++ {
				mat[j][k] -= mat[i][k] * c % mod
				if mat[j][k] < 0 {
					mat[j][k] += mod
				}
			}
		}
	}
	fmt.Println(ans)
}

func sum(a []int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}

const mod = 1000000007

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

func invMod(a int) int {
	return powMod(a, mod-2)
}

type UnionFind struct {
	root []int
	size []int
	n    int
}

func New(size int) *UnionFind {
	uf := new(UnionFind)
	uf.n = size
	uf.root = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Merge(p, q int) bool {
	q = uf.Root(q)
	p = uf.Root(p)

	if q == p {
		return false
	}

	if uf.Size(p) < uf.Size(q) {
		p, q = q, p
	}
	uf.root[q] = p
	uf.size[p] += uf.size[q]
	return true
}

func (uf *UnionFind) Root(p int) int {
	if uf.root[p] == p {
		return p
	}
	uf.root[p] = uf.Root(uf.root[p])
	return uf.root[p]
}

func (uf *UnionFind) Same(p, q int) bool {
	return uf.Root(p) == uf.Root(q)
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Root(x)]
}

func (uf UnionFind) Groups() [][]int {
	rootBuf, groupSize := make([]int, uf.n), make([]int, uf.n)
	for i := 0; i < uf.n; i++ {
		rootBuf[i] = uf.Root(i)
		groupSize[rootBuf[i]]++
	}
	res := make([][]int, uf.n)
	for i := 0; i < uf.n; i++ {
		res[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < uf.n; i++ {
		res[rootBuf[i]] = append(res[rootBuf[i]], i)
	}
	result := make([][]int, 0, uf.n)
	for i := 0; i < uf.n; i++ {
		if len(res[i]) != 0 {
			r := make([]int, len(res[i]))
			copy(r, res[i])
			result = append(result, r)
		}
	}
	return result
}
