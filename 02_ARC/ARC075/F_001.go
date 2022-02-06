package main

import "fmt"

var (
	ans int
	pw  = make([]int, 19)
)

func dfs(x, y, u, v int) {
	if x >= y {
		if u == 0 {
			if x == y {
				ans += v * 10
			} else {
				ans += v
			}
		}
		return
	}
	i := (u/pw[x]%10 + 10) % 10
	j := (pw[y] - pw[x]) / 9
	tmp := 0
	if x == 0 {
		tmp = 1
	}
	dfs(x+1, y-1, u-i*j, v*(10-i-tmp))
	if i != 0 {
		dfs(x+1, y-1, u+(10-i)*j, v*(i-tmp))
	}
}

func main() {
	pw[0] = 1
	for i := 1; i < 19; i++ {
		pw[i] = pw[i-1] * 10
	}

	var d int
	fmt.Scan(&d)
	if d%9 != 0 {
		fmt.Println(0)
		return
	}

	d /= 9
	for i := 0; i < 19; i++ {
		dfs(0, i, d, 1)
	}
	fmt.Println(ans)
}
