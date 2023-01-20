package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnt := make([]int, 2<<17)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		cnt[a]++
	}

	r := n
	l := 0
	ans := 0
	for i := 1; i < 200000; i++ {
		r -= cnt[i]
		ans += l * r * cnt[i]
		l += cnt[i]
	}
	fmt.Println(ans)
}
