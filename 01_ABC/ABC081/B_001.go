package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := int(1e18)
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		res = min(res, divCnt(&a, 2))
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func divCnt(num *int, div int) int {
	cnt := 0
	for *num%div == 0 {
		*num /= div
		cnt++
	}
	return cnt
}
