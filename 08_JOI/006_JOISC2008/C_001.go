package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var id [1000][1000]int
	var T [1 << 17]int

	var N, M, D, K int
	fmt.Fscan(in, &N, &M, &D, &K)
	P := make([]pair, 0)
	for i := 1; i <= N; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		id[x][y] = i
		if i == 1 {
			P = append(P, pair{x, y})
		} else {
			T[i] = 114514
		}
	}
	ans := 0
	for len(P) != 0 {
		x := P[0].x
		y := P[0].y
		P = P[1:]
		i := id[x][y]
		if K >= T[i] && K < T[i]+M {
			ans++
		}
		for dx := -D; dx <= D; dx++ {
			for dy := -D; dy <= D; dy++ {
				tx := x + dx
				ty := y + dy
				if dx*dx+dy*dy <= D*D && tx >= 0 && tx < 1000 && ty >= 0 && ty < 1000 && id[tx][ty] != 0 {
					j := id[tx][ty]
					if T[j] > T[i]+1 {
						T[j] = T[i] + 1
						P = append(P, pair{tx, ty})
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
