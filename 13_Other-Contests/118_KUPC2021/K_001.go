package main

import (
	"bufio"
	"fmt"
	"os"
)

var N int
var G [66][]int
var cnt [66]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var M int
	fmt.Fscan(in, &N, &M)
	for i := 0; i < M; i++ {
		var a, b int
		var x, y string
		fmt.Fscan(in, &a, &x, &b, &y)
		var fx int
		if x == "R" {
			fx = 0
		} else {
			if x == "G" {
				fx = 1
			} else {
				fx = 2
			}
		}
		var fy int
		if y == "R" {
			fy = 0
		} else {
			if y == "G" {
				fy = 1
			} else {
				fy = 2
			}
		}
		a--
		b--
		G[a*3+fx] = append(G[a*3+fx], b*3+fy)
	}
	fmt.Println(dfs(0))
}

func dfs(u int) int {
	if u == N*3 {
		return 1
	}
	none := 0
	ret := 0
	for i := 0; i < 3; i++ {
		if cnt[u+i] == 0 {
			if len(G[u+i]) == 0 {
				none++
			} else {
				for _, e := range G[u+i] {
					cnt[e]++
				}
				ret += dfs(u + 3)
				for _, e := range G[u+i] {
					cnt[e]--
				}
			}
		}
	}
	if none > 0 {
		ret += dfs(u+3) * none
	}
	return ret
}
