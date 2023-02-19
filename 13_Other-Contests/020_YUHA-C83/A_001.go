package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type pair struct {
	x, y int
}

var M, N int

var f [][]string
var cow []int
var sc [][][][]int
var weight [][]int
var ok [][]pair
var cccx, cccy []int
var perm []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &M, &N)
	f = make([][]string, 60)
	for i := 0; i < N; i++ {
		var t string
		fmt.Fscan(in, &t)
		f[i] = strings.Split(t, "")
	}

	cow = make([]int, 30)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			cow[f[i][j][0]-'a']++
		}
	}

	sc = make([][][][]int, 51)
	for i := range sc {
		sc[i] = make([][][]int, 51)
		for j := range sc[i] {
			sc[i][j] = make([][]int, 51)
			for k := range sc[i][j] {
				sc[i][j][k] = make([]int, 51)
			}
		}
	}
	weight = make([][]int, 30)
	for i := range weight {
		weight[i] = make([]int, 30)
	}

	moru := make([][]pair, 30)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			moru[f[i][j][0]-'a'] = append(moru[f[i][j][0]-'a'], pair{i, j})
		}
	}

	for y0 := 0; y0 < N; y0++ {
		for x0 := 0; x0 < M; x0++ {
			for r := 0; r < 2; r++ {
				y1 := y0
				if r == 0 {
					y1++
				}
				x1 := x0
				if r == 1 {
					x1++
				}
				if x1 < 0 || y1 >= N || x1 >= M {
					continue
				}
				if f[y1][x1] != f[y0][x0] {
					c0 := f[y0][x0][0]
					c1 := f[y1][x1][0]
					weight[c0-'a'][c1-'a']++
					weight[c1-'a'][c0-'a']++
					for _, iiict := range moru[c0-'a'] {
						cy := iiict.x
						cx := iiict.y
						for _, iiidt := range moru[c1-'a'] {
							dy := iiidt.x
							dx := iiidt.y
							if comp(abs(cy-y0)+abs(cx-x0), abs(dy-y0)+abs(dx-x0), c0, c1) && !comp(abs(cy-y1)+abs(cx-x1), abs(dy-y1)+abs(dx-x1), c0, c1) {
								sc[cy][cx][dy][dx]++
								sc[dy][dx][cy][cx]++
							}
						}
					}
				}
			}
		}
	}

	ok = make([][]pair, 30)
	for i := 0; i < 30; i++ {
		ok[i] = make([]pair, 0)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			stt := make(map[byte]bool)
			stu := make(map[byte]bool)
			for k := 0; k < N; k++ {
				for l := 0; l < M; l++ {
					if weight[f[i][j][0]-'a'][f[k][l][0]-'a'] != 0 {
						stt[f[k][l][0]] = true
						if weight[f[i][j][0]-'a'][f[k][l][0]-'a'] == sc[i][j][k][l] {
							stu[f[k][l][0]] = true
						}
					}
				}
			}
			if len(stt) == len(stu) {
				ok[f[i][j][0]-'a'] = append(ok[f[i][j][0]-'a'], pair{i, j})
			}
		}
	}
	cccx = make([]int, 30)
	for i := range cccx {
		cccx[i] = -1
	}
	cccy = make([]int, 30)

	pp := make([]pair, 30)
	for i := 0; i < 26; i++ {
		pp[i] = pair{len(ok[i]), i}
	}
	sort.Slice(pp, func(i, j int) bool {
		if pp[i].x == pp[j].x {
			return pp[i].y > pp[j].y
		}
		return pp[i].x > pp[j].x
	})

	perm = make([]int, 26)
	for i := 0; i < 26; i++ {
		perm[i] = pp[i].y
	}
	perm = reverseOrderInt(perm)

	dfs(0)

	for i := 0; i < 26; i++ {
		if cccx[i] >= 0 {
			fmt.Fprintf(out, "%c %d %d\n", int('a')+i, cccx[i], cccy[i])
		}
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func dfs(cc int) bool {
	if cc == 26 {
		return true
	}
	c := perm[cc]
	if cow[c] == 0 {
		return dfs(cc + 1)
	}

	for _, okit := range ok[c] {
		i := okit.x
		j := okit.y
		bad := false
		for uu := 0; uu < cc; uu++ {
			u := perm[uu]
			if cow[u] != 0 {
				if sc[i][j][cccy[u]][cccx[u]] != weight[f[i][j][0]-'a'][f[cccy[u]][cccx[u]][0]-'a'] {
					bad = true
					break
				}
			}
		}
		if !bad {
			cccy[c] = i
			cccx[c] = j
			if dfs(cc + 1) {
				return true
			}
		}
	}
	return false
}

func comp(d1, d2 int, c1, c2 byte) bool {
	if d1 == d2 {
		return c1 < c2
	}
	return d1 < d2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
