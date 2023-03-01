package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 333333

	var n, m int
	fmt.Fscan(in, &n, &m)
	sude := make([][]bool, 2)
	for i := range sude {
		sude[i] = make([]bool, N)
	}
	tmp := make([]bool, N)
	for i := 0; i < n*2; i++ {
		tmp[i] = true
		sude[0][i] = false
		sude[1][i] = false
	}
	v := make([][][]int, 2)
	for i := range v {
		v[i] = make([][]int, N)
	}
	coun := make([]int, N)
	all := 0
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		if c != 0 {
			tmp[a] = false
			tmp[b] = false
			v[0][b] = append(v[0][b], a)
			v[1][a] = append(v[1][a], b)
			coun[a] = b
			coun[b] = a
		} else {
			v[0][a] = append(v[0][a], b)
			v[1][b] = append(v[1][b], a)
		}
		all += c
	}

	var dfs func(int, int)
	dfs = func(s, p int) {
		if !sude[s][p] {
			sude[s][p] = true
			for _, i := range v[s][p] {
				dfs(s, i)
			}
		}
	}
	for i := 0; i < n; i++ {
		if tmp[i*2] {
			dfs(0, i*2)
		}
		if tmp[i*2+1] {
			dfs(1, i*2+1)
		}
	}
	fmt.Println(all)
	for i := 0; i < n; i++ {
		if !sude[0][i*2] {
			sude[0][i*2] = true
			dfs(1, coun[i*2])
			fmt.Println(i*2 + 1)
		}
		if !sude[1][i*2+1] {
			sude[1][i*2] = true
			dfs(0, coun[i*2+1])
			fmt.Println(i*2 + 2)
		}
	}
}
