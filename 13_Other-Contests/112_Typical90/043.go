package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	type pair struct {
		x, y int
	}

	var h, w int
	fmt.Fscan(in, &h, &w)

	var rs, rt, cs, ct int
	fmt.Fscan(in, &rs, &cs, &rt, &ct)
	rs--
	cs--
	rt--
	ct--

	S := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &S[i])
	}

	dis := make([][]int, h)
	for i := range dis {
		dis[i] = make([]int, w)
		for j := range dis[i] {
			dis[i][j] = INF
		}
	}

	dis[rs][cs] = 0
	Q := make([]pair, 0)
	Q = append(Q, pair{rs, cs})
	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	for len(Q) > 0 {
		r := Q[0].x
		c := Q[0].y
		Q = Q[1:]
		var rr, cc int
		for i := 0; i < 4; i++ {
			rr = r
			cc = c
			for {
				rr += dy[i]
				cc += dx[i]
				if rr < 0 || rr >= h || cc < 0 || cc >= w {
					break
				}
				if S[rr][cc] == '#' {
					break
				}
				if dis[rr][cc] <= dis[r][c] {
					break
				}
				dis[rr][cc] = dis[r][c] + 1
				Q = append(Q, pair{rr, cc})
			}
		}
	}

	fmt.Println(dis[rt][ct] - 1)
}
