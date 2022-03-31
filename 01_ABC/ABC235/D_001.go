package main

import "fmt"

var book = [1000002]int{}

func dfs(cnt, a, n int) {
	if book[n] > cnt || book[n] == 0 {
		book[n] = cnt
		if n%a == 0 {
			dfs(cnt+1, a, n/a)
		}
		m := n
		k := 1
		m /= 10
		for ; m > 0; m /= 10 {
			k *= 10
		}
		z := n / k
		n = (n%k)*10 + z
		k = n / k
		if k != 0 {
			dfs(cnt+1, a, n)
		}
	}
}

func main() {
	var a, n int
	fmt.Scan(&a, &n)

	dfs(0, a, n)

	if book[1] == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(book[1])
	}
}
