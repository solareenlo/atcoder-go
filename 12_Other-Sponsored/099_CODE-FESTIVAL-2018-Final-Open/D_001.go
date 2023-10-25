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

	var cnt [128][128][128]int
	var st1 [128]int
	var st2 [128][128]int

	var N int
	fmt.Fscan(in, &N)
	for i := 1; i <= N; i++ {
		var b string
		fmt.Fscan(in, &b)
		Len := len(b)
		v1 := make([]int, 0)
		v2 := make([]pair, 0)
		for j := 0; j < Len; j++ {
			v := int(b[j])
			for _, u := range v1 {
				if st2[u][v] == 0 {
					st2[u][v] = j + 1
					v2 = append(v2, pair{u, v})
				}
			}
			if st1[v] == 0 {
				v1 = append(v1, v)
			}
			st1[v] = j + 1
		}
		for _, a := range v2 {
			for _, b := range v1 {
				if st2[a.x][a.y] < st1[b] {
					cnt[a.x][a.y][b]++
				}
			}
			st2[a.x][a.y] = 0
		}
		for _, b := range v1 {
			st1[b] = 0
		}
	}
	m := 0
	var ret [3]string
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			for k := 0; k < 128; k++ {
				if cnt[i][j][k] > m {
					ret[0] = string(i)
					ret[1] = string(j)
					ret[2] = string(k)
					m = cnt[i][j][k]
				}
			}
		}
	}
	fmt.Print(ret[0], ret[1], ret[2], "\n")
}
