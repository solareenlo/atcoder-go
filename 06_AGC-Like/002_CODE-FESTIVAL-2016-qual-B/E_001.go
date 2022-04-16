package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 400010

var (
	son = [N][26]int{}
	ind = [N]int{}
)

func mod(x int) int {
	ch := -1
	for i := 0; i < 26; i++ {
		if son[x][i] != 0 {
			son[x][i] = mod(son[x][i])
			if ch == -1 {
				ch = i
			} else {
				ch = -2
			}
		}
	}
	if ch == -2 || ind[x] != 0 || x == 1 {
		return x
	} else {
		return son[x][ch]
	}
}

var (
	fa  = [N]int{}
	ch  = [N]int{}
	nxt = [N]int{}
	fst = [N]int{}
)

func dfs(x int) {
	for i := 0; i < 26; i++ {
		if son[x][i] == 0 {
			continue
		}
		fa[son[x][i]] = x
		ch[son[x][i]] = i
		nxt[son[x][i]] = fst[x]
		fst[x] = son[x][i]
		dfs(son[x][i])
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	size := make([]int, N)
	cnt := 1
	pos := make([]int, 100010)
	for i := 1; i <= n; i++ {
		var str string
		fmt.Fscan(in, &str)
		l := len(str)
		str = " " + str
		now := 1
		size[1]++
		for j := 1; j <= l; j++ {
			if son[now][str[j]-'a'] == 0 {
				cnt++
				son[now][str[j]-'a'] = cnt
			}
			now = son[now][str[j]-'a']
			size[now]++
		}
		ind[now] = i
		pos[i] = now
	}

	rt := mod(1)
	dfs(rt)
	var q int
	fmt.Fscan(in, &q)
	p := make([]int, 26)
	for i := 1; i <= q; i++ {
		var k int
		var str string
		fmt.Fscan(in, &k, &str)
		for i := 0; i < 26; i++ {
			p[str[i]-'a'] = i
		}
		ans := 0
		for tem := pos[k]; tem != rt; tem = fa[tem] {
			if ind[tem] != 0 {
				ans++
			}
			c := p[ch[tem]]
			for s := fst[fa[tem]]; s > 0; s = nxt[s] {
				if p[ch[s]] < c {
					ans += size[s]
				}
			}
		}
		fmt.Fprintln(out, ans)
	}
}
