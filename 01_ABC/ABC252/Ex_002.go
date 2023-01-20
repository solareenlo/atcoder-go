package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, c, k int
	fmt.Fscan(in, &n, &c, &k)

	const N = 75
	v := make([][]int, N)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x] = append(v[x], y)
	}

	p := make([]int, N)
	sum := 1
	for i := 1; i <= c; i++ {
		p[i] = i
		sum *= len(v[i])
	}

	k = sum - k + 1
	tmp := p[1 : c+1]
	sort.Slice(tmp, func(i, j int) bool {
		return len(v[i]) < len(v[j])
	})

	s := make([][]int, N)
	for i := 1; i <= c; i++ {
		s[i&1] = append(s[i&1], p[i])
	}

	res := make([][]int, N)
	var dfs func(int, int, int)
	dfs = func(pos, sum, i int) {
		if pos == len(s[i]) {
			res[i] = append(res[i], sum)
			return
		}
		for _, x := range v[s[i][pos]] {
			dfs(pos+1, sum^x, i)
		}
		return
	}
	dfs(0, 0, 0)
	dfs(0, 0, 1)

	const M = 20000010
	var sz [M]int
	var ch [M][2]int
	tot := 1
	var insert func(int)
	insert = func(val int) {
		x := 1
		sz[x]++
		for i := 60; i >= 0; i-- {
			if val&(1<<i) != 0 {
				if ch[x][1] == 0 {
					tot++
					ch[x][1] = tot
				}
				x = ch[x][1]
			} else {
				if ch[x][0] == 0 {
					tot++
					ch[x][0] = tot
				}
				x = ch[x][0]
			}
			sz[x]++
		}
		return
	}
	for _, x := range res[0] {
		insert(x)
	}

	var pos [M]int
	for i := 0; i < len(res[1]); i++ {
		pos[i] = 1
	}

	ans := 0
	for i := 60; i >= 0; i-- {
		sum := 0
		c := 0
		for j := 0; j < len(res[1]); j++ {
			sum += sz[ch[pos[j]][(res[1][j]>>i)&1]]
		}
		if k > sum {
			c = 1
			ans |= 1 << i
			k -= sum
		}
		for j := 0; j < len(res[1]); j++ {
			pos[j] = ch[pos[j]][((res[1][j]>>i)&1)^c]
		}
	}

	fmt.Println(ans)
}
