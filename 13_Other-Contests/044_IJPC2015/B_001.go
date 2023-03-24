package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const MN = 10100
const MA = 30

func main() {
	in := bufio.NewReader(os.Stdin)

	type P struct {
		x, y int
	}

	var g, g2 [MA][MA][]int
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		a := s[0] - 'a'
		b := s[len(s)-1] - 'a'
		g[a][b] = append(g[a][b], i)
		g2[a][b] = append(g2[a][b], i)
	}
	idx := make([]int, 26)
	for i := 0; i < 26; i++ {
		idx[i] = i
	}
	res := make([]int, n)
	for k := 0; k < 10; k++ {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(idx), func(i, j int) { idx[i], idx[j] = idx[j], idx[i] })
		for i := 0; i < MA; i++ {
			for j := 0; j < MA; j++ {
				if len(g[i][j]) == 0 {
					continue
				}
				for a := 0; a < MA; a++ {
					for b := 0; b < MA; b++ {
						g[a][b] = g2[a][b]
					}
				}
				for l := range res {
					res[l] = -2
				}

				u := g[i][j][len(g[i][j])-1]
				g[i][j] = g[i][j][:len(g[i][j])-1]
				res[u] = -1

				que := make([]P, 0)
				que = append(que, P{j, u})
				for len(que) > 0 {
					tmp := que[0]
					p := tmp.x
					ed := tmp.y
					que = que[1:]
					c := 0
					for _, i := range idx {
						c += len(g[p][i])
					}
					if c <= 1 {
						continue
					}
					c = 2
					for _, i := range idx {
						if i != p {
							continue
						}
						for c != 0 && len(g[p][i]) != 0 {
							u := g[p][i][len(g[p][i])-1]
							res[u] = ed
							g[p][i] = g[p][i][:len(g[p][i])-1]
							que = append(que, P{i, u})
							c--
						}
					}
					for _, i := range idx {
						for c != 0 && len(g[p][i]) != 0 {
							u := g[p][i][len(g[p][i])-1]
							res[u] = ed
							g[p][i] = g[p][i][:len(g[p][i])-1]
							que = append(que, P{i, u})
							c--
						}
					}
				}
				f := true
				for i := 0; i < n; i++ {
					if res[i] == -2 {
						f = false
						break
					}
				}
				if !f {
					continue
				}
				fmt.Println("YES")
				for i := 0; i < n; i++ {
					fmt.Println(res[i] + 1)
				}
				return
			}
		}
	}
	fmt.Println("NO")
}
