package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	cnt := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		var a int
		fmt.Scan(&a)
		cnt[i] = (cnt[i-1] + a - 1) % k
	}

	m := map[int]int{}
	m[0] = 1
	res := 0
	for i := 1; i < n+1; i++ {
		if i >= k {
			m[cnt[i-k]]--
		}
		res += m[cnt[i]]
		m[cnt[i]]++
	}

	fmt.Println(res)
}
