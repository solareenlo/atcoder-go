package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	cnt := map[int]int{}
	cnt[0] = 1
	a, sum, res := 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum = (sum + a) % m
		res += cnt[sum]
		cnt[sum]++
	}
	fmt.Println(res)
}
