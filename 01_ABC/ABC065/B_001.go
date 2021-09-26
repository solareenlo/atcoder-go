package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	cnt, tmp := 1, a[1]
	for i := 0; i < n; i++ {
		if tmp == 2 {
			fmt.Println(cnt)
			return
		}
		if tmp == 1 {
			break
		}
		tmp = a[tmp]
		cnt++
	}
	fmt.Println(-1)
}
