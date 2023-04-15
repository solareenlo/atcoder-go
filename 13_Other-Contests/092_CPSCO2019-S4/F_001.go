package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

var K, nex int
var D [1010][1010]int
var F [1010]int
var P []pair
var E [1010][][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &K)
	nex = K
	for j := 0; j < K; j++ {
		for i := 0; i < K; i++ {
			fmt.Fscan(in, &D[j][i])
		}
	}
	for i := 0; i < K; i++ {
		F[i] = D[0][i]
		if i != 0 {
			P = append(P, pair{F[i], i})
		}
	}
	sort.Slice(P, func(i, j int) bool {
		if P[i].x == P[j].x {
			return P[i].y < P[j].y
		}
		return P[i].x < P[j].x
	})
	for _, p := range P {
		dfs(0, p.y)
	}
	fmt.Println(nex)
	for i := 0; i < nex; i++ {
		for _, e := range E[i] {
			fmt.Println(i+1, e[0]+1, F[e[0]]-F[i])
		}
	}
}

func dfs(cur, tar int) {
	for _, e := range E[cur] {
		sp := (F[tar] + F[e[1]] - D[e[1]][tar]) / 2
		if sp >= F[e[0]] {
			dfs(e[0], tar)
			return
		}
		if sp <= F[cur] {
			continue
		}
		F[nex] = sp
		E[nex] = append(E[nex], []int{e[0], e[1]})
		E[nex] = append(E[nex], []int{tar, tar})
		e[0] = nex
		nex++
		return
	}
	E[cur] = append(E[cur], []int{tar, tar})
}
