package main

import "fmt"

type pair struct{ x, y int }

var (
	mp  = make(map[pair]int)
	ans int
	n   int
	P   int = 1331
	ch  string
)

func dfs(dep, r, b, d int) {
	if (dep > n) && (d == 1) {
		mp[pair{r, b}]++
		return
	}
	if (dep <= n) && (d == -1) {
		ans += mp[pair{r, b}]
		return
	}
	dfs(dep+d, r*P+int(ch[dep]-'a')+1, b, d)
	dfs(dep+d, r, b*P+int(ch[dep]-'a'+1), d)
}

func main() {
	fmt.Scan(&n, &ch)
	ch = " " + ch

	dfs(1, 0, 0, 1)
	dfs(n<<1, 0, 0, -1)

	fmt.Println(ans)
}
