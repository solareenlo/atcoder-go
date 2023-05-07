package main

import "fmt"

var n, b, ans int
var prime [4]int = [4]int{2, 3, 5, 7}

func main() {
	fmt.Scan(&n, &b)
	dfs(1, 0)
	if b <= n && f(b) == 0 {
		fmt.Println(ans + 1)
	} else {
		fmt.Println(ans)
	}
}

func dfs(a, t int) {
	if a+b <= n {
		if f(a+b) == a {
			ans++
		}
		for i := t; i < 4; i++ {
			dfs(a*prime[i], i)
		}
	}
}

func f(a int) int {
	if a != 0 {
		return a % 10 * f(a/10)
	}
	return 1
}
