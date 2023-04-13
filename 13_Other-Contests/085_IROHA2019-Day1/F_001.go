package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	if k >= 30 {
		fmt.Println(-1)
		return
	}
	m := 2
	ans := make([]int, k-1)
	for i := range ans {
		for n%m != 0 && m*m <= n {
			m++
		}
		if n%m != 0 {
			fmt.Println(-1)
			return
		} else {
			ans[i] = m
			n /= m
		}
	}
	for _, i := range ans {
		fmt.Printf("%d ", i)
	}
	fmt.Println(n)
}
