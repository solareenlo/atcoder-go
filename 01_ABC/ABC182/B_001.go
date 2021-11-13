package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	cntMax := 0
	res := 0
	for i := 2; i <= 10000; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if a[j]%i == 0 {
				cnt++
			}
		}
		if cnt > cntMax {
			cntMax = cnt
			res = i
		}
	}

	fmt.Println(res)
}
